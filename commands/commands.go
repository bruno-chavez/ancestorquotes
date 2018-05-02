//Package commands contains all the commands of the app.
package commands

import (
	"bufio"
	"fmt"
	"github.com/bruno-chavez/ancestorquotes/quotes"
	"os"
	"strings"
	"time"
)

//AllQuotes prints all quotes to standard output.
func AllQuotes(quoteSlice quotes.QuoteSlice) {

	for _, quote := range quoteSlice {
		fmt.Println(quote.Quote + "\n")
	}
}

//Persistent makes the app print out a quote every certain amount of time.
func Persistent(quoteSlice quotes.QuoteSlice, timer int, measure string) {

	//Running a goroutine allowing the command to print put quotes and read for a stop at the same time.
	switch measure {
	case "minute":
		go func() {
			ticking := time.Tick(time.Duration(timer) * time.Minute)
			for range ticking {
				fmt.Println(quoteSlice.RandomQuote())
			}
		}()

	case "second":
		go func() {
			ticking := time.Tick(time.Duration(timer) * time.Second)
			for range ticking {
				fmt.Println(quoteSlice.RandomQuote())
			}
		}()
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "stop" {
			return
		}
	}
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

	fmt.Println("Type 'stop' to stop talking with the Ancestor")
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

//Search filters every word in the database to match an input word and prints every quote that the word is part of.
func Search(quoteSlice quotes.QuoteSlice, wordToSearch string) {
	//flag is used to know if the search got any matches.
	flag := true

	quoteMatched := make([]string, 0)

	//quote is a QuoteType type.
	//wordXSlice is a slice with one word and white-spaces that appear after a filter is applied.
	//wordXFilter is a string inside a wordXSlice.

	for _, quote := range quoteSlice {
		wordFirstSlice := strings.Split(quote.Quote, " ")

		for _, wordFirstFilter := range wordFirstSlice {
			wordSecondSlice := strings.Split(wordFirstFilter, ",")

			for _, wordSecondFilter := range wordSecondSlice {
				wordThirdSlice := strings.Split(wordSecondFilter, "!")

				for _, wordThirdFilter := range wordThirdSlice {
					filteredWord := strings.Split(wordThirdFilter, ".")

					//After all the filters are applied the filtered word is compared with the word that is been searched.
					//If a match is found, the quote that the filtered word belongs to, is printed.
					if filteredWord[0] == wordToSearch {
						quoteMatched = append(quoteMatched, quote.Quote)
						flag = false
					}
				}
			}
		}
	}
	if flag {
		fmt.Println("'" + wordToSearch + "'" + " is not in any quote.")
	} else {
		//Defining a nested function helps to keep the code clean in the later parts of it.
		inside := func(comparingQuote string, unfilteredSlice []string) bool {
			for _, i := range unfilteredSlice {
				if i == comparingQuote {
					return true
				}
			}
			return false
		}

		cleanSlice := make([]string, 0)

		for _, match := range quoteMatched {
			//To be able to use range on a slice, it must have at least one element.
			if len(cleanSlice) == 0 {
				cleanSlice = append(cleanSlice, match)
			}
			//Comparing the element of cleanSlice and unfilteredSlice
			if inside(match, cleanSlice) {
				continue
			} else {
				cleanSlice = append(cleanSlice, match)
			}
		}
		for _, element := range cleanSlice {
			fmt.Println(element + "\n")
		}
	}
}
