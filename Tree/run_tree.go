package Tree

func RunTree() {
	RunBST()
}

func RunBST() {
	binarySearchTree := NewBinarySearchTree[int]()
	binarySearchTree.Insert(10)
	binarySearchTree.Insert(1)
	binarySearchTree.Insert(15)
	binarySearchTree.Insert(11)
	binarySearchTree.Insert(17)
	binarySearchTree.Insert(20)
	binarySearchTree.Insert(5)
	binarySearchTree.Insert(3)
	binarySearchTree.Search(20)
	binarySearchTree.Search(3)

}
