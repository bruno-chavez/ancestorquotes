//Package structs takes care of processing the whole json top level array.
package structs

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/bruno-chavez/ancestorquotes/quotes"
	"math/rand"
	"os"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	quoteSlice := make([]Quote, 0)

	err := json.Unmarshal(quotes.Q(), &quoteSlice)
	if err != nil {
		panic(err)
	}
	Quotes = Picker(quoteSlice)
}

//Quotes holds all quotes.
var Quotes Picker

//Quote is used to contain quotes.
type Quote struct {
	Quote string `json:"quote"`
}

//Print outputs the quote to stdout.
func (q Quote) Print() {
	fmt.Printf("%v", q.Quote+"\n")
}

//A Filter returns true if the quote meets a condition.
type Filter func(Quote) bool

//Contains creates a filter that tests if the quote contains a substring.
func Contains(s string) Filter {
	return func(q Quote) bool {
		return strings.Contains(q.Quote, s)
	}
}

//Picker is a quote picked by the search command
type Picker []Quote

//Filter picks a quote selected by the filters parameter.
func (p Picker) Filter(filters ...Filter) Picker {
	var result Picker

outer:
	for _, q := range p {
		for _, f := range filters {
			if !f(q) {
				continue outer
			}
		}
		result = append(result, q)
	}
	return result
}

//Random picks a random quote.
func (p Picker) Random() Quote {
	if len(p) < 1 {
		return Quote{}
	}
	return p[rand.Intn(len(p))]
}

//AllQuotes prints all quotes to standard output.
func AllQuotes() {

	quoteSlice := make([]Quote, 0)
	err := json.Unmarshal(quotes.Q(), &quoteSlice)
	if err != nil {
		panic(err)

	} else {
		for _, quote := range quoteSlice {
			fmt.Println(quote.Quote + "\n")
		}
	}
}

//Chat randomly selects a quote ending with a "?" and another one ending with a "." and prints them.
func Chat() {
	quoteSlice := make([]Quote, 0)
	err := json.Unmarshal(quotes.Q(), &quoteSlice)

	if err != nil {
		panic(err)
	} else {
		//split quotes into questions and statements
		var questions []Quote
		var statements []Quote
		for _, quote := range quoteSlice {
			if strings.HasSuffix(quote.Quote, "?") {
				questions = append(questions, quote)
			} else if strings.HasSuffix(quote.Quote, ".") {
				statements = append(statements, quote)
			}
		}
		//select one for each at random
		fmt.Println(questions[rand.Intn(len(questions))].Quote)
		fmt.Println(statements[rand.Intn(len(statements))].Quote)
	}
}

//Talk to the Ancestor and the Ancestor replies back in a random manner
func TalkBack() {
	userName := ""
	userReply := ""
	reader := bufio.NewReader(os.Stdin)
	quoteSlice := make([]Quote, 0)
	err := json.Unmarshal(quotes.Q(), &quoteSlice)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Enter your name:")
		fmt.Scanln(&userName)
		fmt.Println("Hi " + userName)
		for {
			fmt.Println("What do you wanna say?")
			userReply, _ = reader.ReadString('\n')
			if userReply == "stop\n" {
				fmt.Println("GoodBye " + userName)
				fmt.Println("Bear in mind my last quote")
				fmt.Println("Ancestor says: " + quoteSlice[rand.Intn(len(quoteSlice))].Quote)
				break
			}
			fmt.Println("Ancestor says: " + quoteSlice[rand.Intn(len(quoteSlice))].Quote)
		}
	}
}
