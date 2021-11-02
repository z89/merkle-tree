package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

// sha256("test123") = "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae"

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func readFile() string {
	defer timeTrack(time.Now(), "readFile")

	body, err := ioutil.ReadFile("generated.json")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return string(body)
}

func main() {
	json := readFile()
	reader := strings.NewReader(json)
	defer timeTrack(time.Now(), "main")

	// fmt.Printf("%T", reader)

	// createNode(nil, nil, []byte(json)) // create leaf node

	count, err := lineCounter(reader)
	if err != nil {
		log.Fatalf("line counter problem!")
	}

	fmt.Printf("%d\n", count)
}
