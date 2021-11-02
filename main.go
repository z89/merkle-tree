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

	json := readFile("./data/generated.json")
	pic := readFile("./data/test.jpg")
	text := readFile("./data/text.txt")

	// placeholder data for leaf nodes
	treeObjects = append(treeObjects, []byte(json))
	treeObjects = append(treeObjects, []byte(text))
	treeObjects = append(treeObjects, []byte(pic))
	treeObjects = append(treeObjects, []byte("assd"))

	treeObjects = append(treeObjects, []byte("assd"))
	treeObjects = append(treeObjects, []byte("assd"))
	treeObjects = append(treeObjects, []byte("assd"))
	treeObjects = append(treeObjects, []byte("assd"))

	tree := createTree(treeObjects)

	fmt.Printf("Accessing through tree object:\n")
	fmt.Printf("root: %s\n", hex.EncodeToString(tree.RootNode.hash))
	fmt.Printf("right child hash: %s\n", hex.EncodeToString(tree.RootNode.rightChild.hash))
	fmt.Printf("left child hash: %s\n", hex.EncodeToString(tree.RootNode.leftChild.hash))
	// fmt.Printf("right right child hash: %s\n", hex.EncodeToString(tree.RootNode.rightChild.rightChild.hash))
	// fmt.Printf("left left child hash: %s\n\n", hex.EncodeToString(tree.RootNode.leftChild.leftChild.hash))
}
