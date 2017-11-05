//Package random takes care of proccesing the whole json top level array.
package random

import (
	"encoding/json"
	"fmt"
	"github.com/bruno-chavez/ancestorquotes/quotes"
	"math/rand"
	"time"
)

type Quote struct {
	Quote string `json:"quote"`
}

//Returns a single, randomly picked quote
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
