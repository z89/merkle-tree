package main

import (
	"encoding/hex"
	"fmt"
)

type MerkleTree struct {
	RootNode *MerkleNode
}

type MerkleNode struct {
	hash       []byte
	data       []byte
	rightChild *MerkleNode
	leftChild  *MerkleNode
}

func displayNode(node MerkleNode) {
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

func createNode(rightChild, leftChild *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{} // initalise node

	if rightChild == nil && leftChild == nil {
		/** creating child node (a "leaf" type node which does not contain any child nodes) **/
		// hash := sha256.Sum256(data)
		hash := generateArgon2Hash(data)

		node.hash = hash[:]
		node.data = data
		node.rightChild = nil
		node.leftChild = nil

		displayNode(node)
	} else {
		/** creating parent node (a "branch" type node which links the "root" and "leaf" nodes) **/
		childrenHashes := append(rightChild.hash, leftChild.hash...)
		rawData := append(childrenHashes, data...)

		hash := generateArgon2Hash(rawData)

		node.hash = hash[:]
		node.data = data
		node.rightChild = rightChild
		node.leftChild = leftChild

		displayNode(node)
	}

	return &node
}

func createTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	if len(data)%2 != 0 {
		data = append(data, data[len(data)-1])
	}

	for _, dat := range data {
		node := createNode(nil, nil, dat)
		nodes = append(nodes, *node)
	}

	for i := 0; i < len(data)/2-1; i++ {
		var level []MerkleNode

		for j := 0; j < len(nodes); j += 2 {
			node := createNode(&nodes[j], &nodes[j+1], nil)
			level = append(level, *node)
		}

		nodes = level
	}

	tree := MerkleTree{&nodes[0]}

	return &tree
}
