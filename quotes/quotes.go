// Package quotes takes care of processing the JSON file, define types and methods that are needed for commands.
package quotes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"path/filepath"
	"time"
)

// init is needed to seed the rand package.
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// QuoteType is the type that each quote will have once the JSON file is processed.
type QuoteType struct {
	Quote string `json:"quote"`
}

// QuoteSlice is used to parse the JSON file under a single, usable type.
type QuoteSlice []QuoteType

// Parse fetches quotes.json and puts it on a QuoteSlice type.
func Parse() QuoteSlice {

	absPath, _ := filepath.Abs("../quotes/quotes.json")
	readJSON, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	// The capacity is 393 because it is the total number of quotes, subject to change.
	// The capacity is manually set for a more optimized program.
	parsedJSON := make(QuoteSlice, 0, 393)
	err2 := json.Unmarshal(readJSON, &parsedJSON)
	if err2 != nil {
		panic(err2)
	}

	return parsedJSON
}

// RandomQuote is a method of the QuoteSlice type that returns a random quote.
func (q QuoteSlice) RandomQuote() string {
	return q[rand.Intn(len(q))].Quote
}
