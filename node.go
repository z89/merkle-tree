package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Node struct {
	hash       []byte
	data       []byte
	rightChild *Node
	leftChild  *Node
}

func displayNode(node Node) {
	/** display node info **/
	if node.rightChild != nil && node.leftChild != nil {
		fmt.Printf("+----------------------------+\n+    created parent node     +\n+----------------------------+\n")

		fmt.Printf("hash: %s\ndata: %s\nrightChild: %s\nleftChild: %s\n\n",
			hex.EncodeToString(node.hash),
			node.data,
			hex.EncodeToString(node.rightChild.hash),
			hex.EncodeToString(node.leftChild.hash),
		)
	} else {
		fmt.Printf("+----------------------------+\n+     created leaf node      +\n+----------------------------+\n")

		fmt.Printf("hash: %s\ndata: %s\nrightChild: %s\nleftChild: %s\n\n",
			hex.EncodeToString(node.hash),
			node.data,
			"nil",
			"nil",
		)
	}
}

func createNode(rightChild, leftChild *Node, data []byte) *Node {
	node := Node{} // initalise node

	if rightChild == nil && leftChild == nil {
		/** creating child node (a "leaf" type node which does not contain any child nodes) **/
		hash := sha256.Sum256(data)

		node.hash = hash[:]
		node.data = data
		node.rightChild = nil
		node.leftChild = nil

		displayNode(node)
	} else {
		/** creating parent node (a "branch" type node which links the "root" and "leaf" nodes) **/
		childrenHashes := append(rightChild.hash, leftChild.hash...)
		rawData := append(childrenHashes, data...)

		hash := sha256.Sum256(rawData)

		node.hash = hash[:]
		node.data = data
		node.rightChild = rightChild
		node.leftChild = leftChild

		displayNode(node)
	}

	return &node
}
