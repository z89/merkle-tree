package main

// sha256("test123") = "ecd71870d1963316a97e3ac3408c9835ad8cf0f3c1bc703527c30265534f75ae"

func main() {
	var leaf1 = createNode(nil, nil, []byte("test123")) // create leaf node
	var leaf2 = createNode(nil, nil, []byte("test123")) // create leaf node

	createNode(leaf1, leaf2, []byte("test123")) // create parent node

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
