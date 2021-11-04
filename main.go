package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func readFile(filepath string) string {
	defer timeTrack(time.Now(), "readFile")

	body, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	return string(body)
}

func main() {
	defer timeTrack(time.Now(), "main")

	var treeObjects [][]byte

	// real files for testing

	json := readFile("./data/generated.json")
	pic := readFile("./data/test.jpg")
	text := readFile("./data/text.txt")

	treeObjects = append(treeObjects, []byte(json))
	treeObjects = append(treeObjects, []byte(text))
	treeObjects = append(treeObjects, []byte(pic))

	treeObjects = append(treeObjects, []byte("hello"))
	treeObjects = append(treeObjects, []byte("world"))

	treeObjects = append(treeObjects, []byte("lorem"))
	treeObjects = append(treeObjects, []byte("ipsum"))

	tree := createTree(treeObjects)

	fmt.Printf("merkle tree: \n")
	fmt.Printf("root: %s\n\n", hex.EncodeToString(tree.root.hash))
	fmt.Printf("tree depth: %d\n\n", tree.depth)

	// display data from second appended file (text.txt) for testing
	fmt.Printf("leaf node test	****\n")
	fmt.Printf("leaf node hash	: %s\n", hex.EncodeToString(tree.root.left.left.right.hash))
	fmt.Printf("leaf node data	: %s\n", tree.root.left.left.right.data)
	fmt.Printf("original data	: %s\n\n", text)

	// display data from second appended piece of data ("lorem") for testing
	fmt.Printf("leaf node test	****\n")
	fmt.Printf("leaf node hash	: %s\n", hex.EncodeToString(tree.root.right.left.right.hash))
	fmt.Printf("leaf node data	: %s\n", tree.root.right.left.right.data)
	fmt.Printf("original data	: %s\n\n", []byte("lorem"))
}
