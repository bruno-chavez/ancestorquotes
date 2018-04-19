//Package quotes takes care of processing the JSON file, define types and methods to be used in commands.
package quotes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

//QuoteType is the type that each quote will have once the JSON file is processed.
type QuoteType struct {
	Quote string `json:"quote"`
}

//QuoteSlice is used to parse the whole json quotes in a slice of the QuoteType type.
type QuoteSlice []QuoteType

//Parse fetches from quotes.json and puts it on a QuoteSlice type slice.
func Parse() QuoteSlice {

	rawJSON, err := os.Open("quotes/quotes.json")
	if err != nil {
		panic(err)
	}

	readJSON, err2 := ioutil.ReadAll(rawJSON)
	if err2 != nil {
		panic(err2)
	}

	parsedJSON := make(QuoteSlice, 0)
	err3 := json.Unmarshal(readJSON, &parsedJSON)
	if err3 != nil {
		panic(err3)
	}

	return parsedJSON
}

//RandomQuote method returns a random quote.
func (q QuoteSlice) RandomQuote() string {
	return q[rand.Intn(len(q))].Quote
}
