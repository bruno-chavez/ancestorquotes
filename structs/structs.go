//Package structs takes care of processing the whole json top level array.
package structs

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/bruno-chavez/ancestorquotes/quotes"
)

//Quote is used to contain quotes
type Quote struct {
	Quote string `json:"quote"`
}

//RandomQuote returns a single, randomly picked quote
func RandomQuote() {

	rand.Seed(time.Now().UTC().UnixNano())

	quoteSlice := make([]Quote, 0)

	err := json.Unmarshal(quotes.Q(), &quoteSlice)
	if err != nil {
		panic(err)

	} else {
		selectedQuote := quoteSlice[rand.Intn(len(quoteSlice))]
		fmt.Printf("%v", selectedQuote.Quote+"\n")
	}
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
