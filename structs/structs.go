//Package structs takes care of processing the whole json top level array.
package structs

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/bruno-chavez/ancestorquotes/quotes"
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

// Quotes holds all quotes.
var Quotes Picker

//Quote is used to contain quotes
type Quote struct {
	Quote string `json:"quote"`
}

// Print outputs the quote to stdout.
func (q Quote) Print() {
	fmt.Printf("%v", q.Quote+"\n")
}

// A Filter returns true if the quote meets a condition.
type Filter func(Quote) bool

// Contains creates a filter that tests if the quote contains a substring.
func Contains(s string) Filter {
	return func(q Quote) bool {
		return strings.Contains(q.Quote, s)
	}
}

type Picker []Quote

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

func (p Picker) Random() Quote {
	if len(p) < 1 {
		return Quote{}
	}
	return p[rand.Intn(len(p))]
}

//AllQuotes prints all quotes to standard output
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
