package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
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
	defer timeTrack(time.Now(), "main")

	fmt.Printf(hex.EncodeToString(generateHash(readFile())) + "\n")

	// var leaf1 = createNode(nil, nil, []byte("test123")) // create leaf node
	// var leaf2 = createNode(nil, nil, []byte("test123")) // create leaf node

	// var leaf3 = createNode(nil, nil, []byte("test123")) // create leaf node
	// var leaf4 = createNode(nil, nil, []byte("test123")) // create leaf node

	// var leaf5 = createNode(nil, nil, []byte("test123")) // create leaf node
	// var leaf6 = createNode(nil, nil, []byte("test123")) // create leaf node

	// var leaf7 = createNode(nil, nil, []byte("test123")) // create leaf node
	// var leaf8 = createNode(nil, nil, []byte("test123")) // create leaf node

	// var parent1 = createNode(leaf1, leaf2, []byte("test123")) // create parent node one
	// var parent2 = createNode(leaf3, leaf4, []byte("test123")) // create parent node two

	// var parent3 = createNode(leaf5, leaf6, []byte("test123")) // create parent node three
	// var parent4 = createNode(leaf7, leaf8, []byte("test123")) // create parent node four

	// var parent2_1 = createNode(parent1, parent2, []byte("test123")) // create second parent node one
	// var parent2_2 = createNode(parent3, parent4, []byte("test123")) // create second parent node two

	// var root = createNode(parent2_1, parent2_2, []byte("test123")) // create root node

	// fmt.Printf(hex.EncodeToString(root.hash))
	/*
		console result:
			+----------------------------+
			+     created leaf node      +
			+----------------------------+
			hash: ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae
			data: test123
			rightChild: nil
			leftChild: nil

			+----------------------------+
			+     created leaf node      +
			+----------------------------+
			hash: ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae
			data: test123
			rightChild: nil
			leftChild: nil

			+----------------------------+
			+    created parent node     +
			+----------------------------+
			hash: c9ff8ca5c3b73ca4e30ea5ed4570a75a4fe235bd40479ceb695316c779166b78
			data: test123
			rightChild: ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae
			leftChild: ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae
	*/
}
