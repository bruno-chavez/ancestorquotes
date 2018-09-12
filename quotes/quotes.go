// Package quotes takes care of processing the JSON file, define types and methods that are needed for commands.
package quotes

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"runtime"
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

func getPath() (path string) {
	path = os.Getenv("GOPATH") + "/src/github.com/bruno-chavez/ancestorquotes/quotes/quotes.json"
	if runtime.GOOS == "windows" {
		path = os.Getenv("GOPATH") + "\\src\\github.com\\bruno-chavez\\ancestorquotes\\quotes\\quotes.json"
	}
	if os.Getenv("DOCKER") != "" {
		path = "./quotes/quotes.json"

	}
	return
}

// Parse fetches quotes.json and puts it on a QuoteSlice type.
func Parse() (QuoteSlice, error) {
	rawJSON, err := os.Open(getPath())
	if err != nil {
		return nil, err
	}

	readJSON, err := ioutil.ReadAll(rawJSON)
	if err != nil {
		return nil, err
	}

	// The capacity is 393 because it is the total number of quotes, subject to change.
	// The capacity is manually set for a more optimized program.
	parsedJSON := make(QuoteSlice, 0, 393)
	err = json.Unmarshal(readJSON, &parsedJSON)
	if err != nil {
		return nil, err
	}

	return parsedJSON, nil
}

// RandomQuote is a method of the QuoteSlice type that returns a random quote.
func (q QuoteSlice) RandomQuote() string {
	return q[rand.Intn(len(q))].Quote
}
