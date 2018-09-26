// Package quotes takes care of processing the JSON file, define types and methods that are needed for commands.
package quotes

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
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
func Parse() (qs QuoteSlice) {

	var path = ""
	if os.Getenv("DOCKER") != "" {
		path = "./quotes/quotes.json"

	} else {
		pwd, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		path = filepath.Join(pwd, "quotes", "quotes.json")
	}

	jsonFile, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	// Get size of JSON file in bytes
	sizeInBytes, _ := jsonFile.Stat()
	data := make([]byte, sizeInBytes.Size())

	_, _ = jsonFile.Read(data) // Read JSON data into byte slice
	err = json.Unmarshal(data, &qs)
	if err != nil {
		log.Fatal(err)
	}

	return qs
}

// RandomQuote is a method of the QuoteSlice type that returns a random quote.
func (q QuoteSlice) RandomQuote() string {
	return q[rand.Intn(len(q))].Quote
}
