package quotes

import (
	"os"
	"runtime"
	"testing"
)

func TestParse(t *testing.T) {
	_, err := Parse()
	if err != nil {
		t.Fatalf("%s", err)
	}
}

func TestGetPath(t *testing.T) {
	path := getPath()
	var wantedPath string
	if runtime.GOOS == "windows" {
		wantedPath = os.Getenv("GOPATH") + "\\src\\github.com\\bruno-chavez\\ancestorquotes\\quotes\\quotes.json"
	} else {
		wantedPath = os.Getenv("GOPATH") + "/src/github.com/bruno-chavez/ancestorquotes/quotes/quotes.json"
	}
	if path != wantedPath {
		t.Fatalf("When testing either linux or windows enviroment. The path returned was not correct. The path received was \"%s\"", path)
	}

	_ = os.Setenv("DOCKER", "SMEG")
	path = getPath()
	if path != "./quotes/quotes.json" {
		t.Fatalf("When testing Docker enviroment. The path returned was not correct. The path received was \"%s\"", path)
	}
}
