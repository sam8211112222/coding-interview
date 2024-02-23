package Tree

import "fmt"

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
	// 构建一个较大的BST
	root := &TreeNode{10,
		&TreeNode{5,
			&TreeNode{2, nil, nil},
			&TreeNode{7, nil, nil}},
		&TreeNode{15,
			&TreeNode{12, nil, nil},
			&TreeNode{20, nil, nil}}}

	fmt.Println("Checking if the tree is a valid BST:", isValidBST(root)) // 应输出：true

	// 修改树使其违反BST的条件
	// 例如，将节点值15的左子节点改为值大于15的节点
	root.Right.Left = &TreeNode{17, nil, nil}

	fmt.Println("Checking if the modified tree is a valid BST:", isValidBST(root)) // 应输出：false
}
