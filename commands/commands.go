// Package commands contains all the commands of the app.
package commands

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bruno-chavez/ancestorquotes/quotes"
)

// AllQuotes prints all quotes.
func AllQuotes(quoteSlice quotes.QuoteSlice) {

	for _, quote := range quoteSlice {
		fmt.Println(quote.Quote + "\n")
	}
}

// Persistent prints a quote every certain amount of time.
func Persistent(quoteSlice quotes.QuoteSlice, timer int, measure string) {

	// Running a goroutine makes it possible to print a quote and wait for a stop message at the same time.
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

// Chat randomly selects a quote ending with a "?" and another one ending with a "." and prints them.
func Chat(quoteSlice quotes.QuoteSlice) {
	// split quotes into questions and statements
	var questions quotes.QuoteSlice
	var statements quotes.QuoteSlice

	for _, quote := range quoteSlice {
		if strings.HasSuffix(quote.Quote, "?") {
			questions = append(questions, quote)
		} else if strings.HasSuffix(quote.Quote, ".") {
			statements = append(statements, quote)
		}
	}
	// select one for each at random
	fmt.Println(questions.RandomQuote())
	fmt.Println(statements.RandomQuote())
}

// TalkBack prints a quote every time it gets an input message.
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

// Search filters every word in the database to match an input word and prints every quote that the word is part of.
func Search(quoteSlice quotes.QuoteSlice, wordToSearch string) {
	// matched is used to know if the search got any matches.
	matched := false

	quoteMatched := make([]string, 0)

	// quote is a QuoteType type.
	// wordXSlice is a slice with one word and white-spaces that appear after a filter is applied.
	// wordXFilter is a string inside a wordXSlice.
	// This for is a candidate for been transformed into an auxiliary, recursive function.
	for _, quote := range quoteSlice {
		wordFirstSlice := strings.Split(quote.Quote, " ")

		for _, wordFirstFilter := range wordFirstSlice {
			wordSecondSlice := strings.Split(wordFirstFilter, ",")

			for _, wordSecondFilter := range wordSecondSlice {
				wordThirdSlice := strings.Split(wordSecondFilter, "!")

				for _, wordThirdFilter := range wordThirdSlice {
					filteredWord := strings.Split(wordThirdFilter, ".")

					// After all the filters are applied the filtered word is compared with the word that is been searched.
					// If a match is found, the quote that the filtered word belongs to, is printed.
					if filteredWord[0] == wordToSearch {
						quoteMatched = append(quoteMatched, quote.Quote)
						matched = true
					}
				}
			}
		}
	}
	if matched {
		// Defining a nested function makes the next part clearer to see.
		// Candidate for been taking out of the Search func and be used as an auxiliary, separate func instead.
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
			// To be able to use range on a slice, it must have at least one element.
			if len(cleanSlice) == 0 {
				cleanSlice = append(cleanSlice, match)
			}
			// Comparing the element of cleanSlice and unfilteredSlice
			if inside(match, cleanSlice) {
				continue
			} else {
				cleanSlice = append(cleanSlice, match)
			}
		}
		for _, element := range cleanSlice {
			fmt.Println(element + "\n")
		}
	} else {
		fmt.Println("'" + wordToSearch + "'" + " is not present in any quote.")

	}
}
