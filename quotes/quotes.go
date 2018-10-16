// Package quotes takes care of processing the JSON file, define types and methods that are needed for commands.
package quotes

import (
	"encoding/json"
	"io/ioutil"
	"log"
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

// Parse fetches quotes.json and puts it on a QuoteSlice type.
func Parse() QuoteSlice {
	// Extremely tedious way to always find the json file.
	// Its pretty horrible to look at, but it works, somebody please do it better than me.

	var path = ""
	// if operating system is windows
	if runtime.GOOS == "windows" {
		if os.Getenv("DOCKER") != "" {
			path = ".\\quotes\\quotes.json"
		} else {
			path = os.Getenv("GOPATH") + "\\src\\github.com\\bruno-chavez\\ancestorquotes\\quotes\\quotes.json"
		}
	} else { // if operating system isn't windows
		if os.Getenv("DOCKER") != "" {
			path = "./quotes/quotes.json"
		} else {
			path = os.Getenv("GOPATH") + "/src/github.com/bruno-chavez/ancestorquotes/quotes/quotes.json"
		}
	}

	rawJSON, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	readJSON, err2 := ioutil.ReadAll(rawJSON)
	if err2 != nil {
		log.Fatal(err2)
	}

	// The capacity is 393 because it is the total number of quotes, subject to change.
	// The capacity is manually set for a more optimized program.
	parsedJSON := make(QuoteSlice, 0, 393)

	err3 := json.Unmarshal(readJSON, &parsedJSON)
	if err3 != nil {
		log.Fatal(err3)
	}

	return parsedJSON
}

// RandomQuote is a method of the QuoteSlice type that returns a random quote.
func (q QuoteSlice) RandomQuote() string {
	return q[rand.Intn(len(q))].Quote
}
