package Tree

func RunTree() {
	//RunBST()
	//RunBFS()
	isValidBSTTest()
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
	binarySearchTree.Insert(2)
	binarySearchTree.Insert(7)
	binarySearchTree.Insert(6)
	binarySearchTree.Insert(12)
	binarySearchTree.Insert(16)
	binarySearchTree.Insert(18)
	binarySearchTree.Delete(5)

}

func RunBFS() {
	binarySearchTree := NewBinarySearchTree[int]()
	binarySearchTree.Insert(9)
	binarySearchTree.Insert(4)
	binarySearchTree.Insert(6)
	binarySearchTree.Insert(20)
	binarySearchTree.Insert(170)
	binarySearchTree.Insert(15)
	binarySearchTree.Insert(1)

	//binarySearchTree.BreadthFirstSearch()
	binarySearchTree.DFSInOrder()
	binarySearchTree.DFSPreOrder()
	binarySearchTree.DFSPostOrder()

}

func isValidBSTTest() {

}
