package main

type MerkleTree struct {
	depth uint32
	root  *MerkleNode
}

type MerkleNode struct {
	hash  []byte
	data  []byte
	left  *MerkleNode
	right *MerkleNode
}

func createNode(left, right *MerkleNode, data []byte) *MerkleNode {
	node := MerkleNode{}

	if left == nil && right == nil {
		hash := generateArgon2Hash(data)

		node.hash = hash[:]
		node.data = data
		node.left = nil
		node.right = nil

	} else {
		childrenHashes := append(left.hash, right.hash...)
		rawData := append(childrenHashes, data...)

		hash := generateArgon2Hash(rawData)
		node.hash = hash[:]
		node.data = data
		node.left = left
		node.right = right
	}

	return &node
}

func createTree(data [][]byte) *MerkleTree {
	var nodes []MerkleNode

	for i := 0; i < len(data); i++ {
		node := createNode(nil, nil, data[i])
		nodes = append(nodes, *node)
	}

	if len(nodes)%2 != 0 {
		node := createNode(nil, nil, []byte("empty data leaf node"))
		nodes = append(nodes, *node)
	}

	var temp []MerkleNode = nodes

	var depth uint32 = 1
	for i := 0; i < (len(nodes)/2)-1; i++ {
		var round []MerkleNode

		if len(temp)%2 != 0 {
			node := createNode(&temp[0], &temp[0], []byte("empty data parent node (odd)"))
			temp = append(temp, *node)
		}

		for i := 0; i < len(temp); i += 2 {
			node := createNode(&temp[i], &temp[i+1], []byte("empty data parent node"))
			round = append(round, *node)
		}

		depth++
		temp = round
	}

	// debugging
	root := temp[0]

	// fmt.Printf("root hash of tree of root buffer: %s value: %s\n\n", hex.EncodeToString(root.hash), root.data)

	// fmt.Printf("left: values of root.left: %s value: %s\n", hex.EncodeToString(root.left.hash), root.left.data)
	// fmt.Printf("right: values of root.right:  %s value: %s\n\n", hex.EncodeToString(root.right.hash), root.right.data)

	// fmt.Printf("left left:  values of root.left.left:  %s value: %s\n", hex.EncodeToString(root.left.left.hash), root.left.left.data)
	// fmt.Printf("left right:  values of root.left.right:  %s value: %s\n", hex.EncodeToString(root.left.right.hash), root.left.right.data)
	// fmt.Printf("right left:  values of root.right.left:  %s value: %s\n", hex.EncodeToString(root.right.left.hash), root.right.left.data)
	// fmt.Printf("right right:  values of root.right.right:  %s value: %s\n\n", hex.EncodeToString(root.right.right.hash), root.right.right.data)

	// fmt.Printf("left left left:  values of root.left.left.left:  %s value: %s\n", hex.EncodeToString(root.left.left.left.hash), root.left.left.left.data)
	// fmt.Printf("left left right:  values of root.left..left.right:  %s value: %s\n", hex.EncodeToString(root.left.left.right.hash), root.left.left.right.data)
	// fmt.Printf("left right left:  values of root.left.right.left:  %s value: %s\n", hex.EncodeToString(root.left.right.left.hash), root.left.right.left.data)
	// fmt.Printf("left right right:  values of root.left.right.right:  %s value: %s\n\n", hex.EncodeToString(root.left.right.right.hash), root.left.right.right.data)

	// fmt.Printf("right left left:  values of root.right.left.left:  %s value: %s\n", hex.EncodeToString(root.right.left.left.hash), root.right.left.left.data)
	// fmt.Printf("right left right:  values of root.right.left.right:  %s value: %s\n", hex.EncodeToString(root.right.left.right.hash), root.right.left.right.data)
	// fmt.Printf("right right left:  values of root.right.right.left:  %s value: %s\n", hex.EncodeToString(root.right.right.left.hash), root.right.right.left.data)
	// fmt.Printf("right right right:  values of root.right.right.right:  %s value: %s\n\n", hex.EncodeToString(root.right.right.right.hash), root.right.right.right.data)

	tree := MerkleTree{depth, &root}

	return &tree

}
