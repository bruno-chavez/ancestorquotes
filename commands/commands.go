//Package commands contains all the commands of the app.
package commands

import (
	"bufio"
	"fmt"
	"github.com/bruno-chavez/ancestorquotes/quotes"
	"os"
	"strconv"
	"strings"
)

//AllQuotes prints all quotes to standard output.
func AllQuotes(quoteSlice quotes.QuoteSlice) {

	for _, quote := range quoteSlice {
		fmt.Println(quote.Quote + "\n")
	}
}

//SliceSecondsMinutes populates a slice with the number of seconds in a minute and the number of minutes in an hour and returns it.
func SliceSecondsMinutes() []string {
	slice := make([]string, 60)

	for i := 1; i < 60; i++ {
		slice[i] = strconv.Itoa(i)
	}

	return slice
}

//Chat randomly selects a quote ending with a "?" and another one ending with a "." and prints them.
func Chat(quoteSlice quotes.QuoteSlice) {
	//split quotes into questions and statements
	var questions quotes.QuoteSlice
	var statements quotes.QuoteSlice

	for _, quote := range quoteSlice {
		if strings.HasSuffix(quote.Quote, "?") {
			questions = append(questions, quote)
		} else if strings.HasSuffix(quote.Quote, ".") {
			statements = append(statements, quote)
		}
	}
	//select one for each at random
	fmt.Println(questions.RandomQuote())
	fmt.Println(statements.RandomQuote())
}

//TalkBack talks to the Ancestor and the Ancestor replies back in a random manner
func TalkBack(quoteSlice quotes.QuoteSlice) {
	userName := ""
	userReply := ""
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name:")
	fmt.Scanln(&userName)

	if userName == "stop" {
		fmt.Println("Good bye nameless fiend!")
		fmt.Println("Bear in mind my last quote")
		fmt.Println("Ancestor says: " + quoteSlice.RandomQuote())
		return
	}
	fmt.Println("Hi " + userName)

	for {
		fmt.Println("What do you wanna say?")
		userReply, _ = reader.ReadString('\n')
		if userReply == "stop\n" {
			fmt.Println("Good bye " + userName)
			fmt.Println("Bear in mind my last quote")
			fmt.Println("Ancestor says: " + quoteSlice.RandomQuote())
			break
		}
		fmt.Println("Ancestor says: " + quoteSlice.RandomQuote())
	}
}
