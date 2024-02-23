package Tree

import (
	"errors"
	"fmt"
)

// NumberType 要自己宣告泛型裡面的type, 不然如果用any, comparable 都沒辦法編譯
// comparable 约束在 Go 泛型中用于确保类型可以进行等值比较（即 == 和 != 操作），这对于很多泛型数据结构和函数来说是足够的。
// 然而，comparable 并不保证类型支持顺序比较操作（如 <、>、<=、>=）。这是因为 comparable 约束的设计目的是为了确保类型可以用在需要等值比较的场合，比如作为 map 的键。

var errNotFound = errors.New("node not found")

type NumberType interface {
	int | float64
}

type Node[T NumberType] struct {
	Value *T
	Right *Node[T]
	Left  *Node[T]
}

type BinarySearchTree[T NumberType] struct {
	Root   *Node[T]
	Length int
}

func NewBinarySearchTree[T NumberType]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) Insert(insertValue T) {
	if bst.Root == nil {
		bst.Root = &Node[T]{Value: &insertValue}
		bst.Length++
		return
	}
	// 遞迴調用insertNode
	bst.insertNode(bst.Root, insertValue)
}

func (bst *BinarySearchTree[T]) insertNode(current *Node[T], insertValue T) {
	if insertValue < *current.Value {
		if current.Left == nil {
			current.Left = &Node[T]{Value: &insertValue}
			bst.Length++
		} else {
			bst.insertNode(current.Left, insertValue)
		}
	} else {
		if current.Right == nil {
			current.Right = &Node[T]{Value: &insertValue}
			bst.Length++
		} else {
			bst.insertNode(current.Right, insertValue)
		}
	}
}

func (bst *BinarySearchTree[T]) Search(value T) (*T, error) {

	if bst.Root == nil {
		return nil, errNotFound
	}
	return bst.Root.SearchNode(value)
}

func (n *Node[T]) SearchNode(value T) (*T, error) {
	if *n.Value == value {
		return n.Value, nil
	}
	if value < *n.Value {
		if n.Left != nil {
			return n.Left.SearchNode(value)
		}
		return nil, errNotFound
	} else {
		if n.Right != nil {
			return n.Right.SearchNode(value)
		}
		return nil, errNotFound
	}
}

func (bst *BinarySearchTree[T]) SearchChildNodeAndParentNode(value T) (*Node[T], *Node[T], error) {
	var parent *Node[T]
	node := bst.Root
	for node != nil {
		if value == *node.Value {
			return node, parent, nil
		} else if value < *node.Value {
			parent = node
			node = node.Left
		} else {
			parent = node
			node = node.Right
		}
	}
	return nil, nil, errNotFound
}

func (bst *BinarySearchTree[T]) Delete(value T) error {
	nodeToDelete, parent, err := bst.SearchChildNodeAndParentNode(value)
	if err != nil {
		return err
	}
	bst.deleteNode(nodeToDelete, parent)

	return nil
}

func (bst *BinarySearchTree[T]) deleteNode(nodeToDelete, parent *Node[T]) {
	// 情况1：無子節點
	//只需要考慮原本指向此節點的parent node，將其 left child/right child 指向NULL即可。
	if nodeToDelete.Left == nil && nodeToDelete.Right == nil {
		if parent == nil { // 要删除的是根节点
			bst.Root = nil
		} else if parent.Left == nodeToDelete {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		//情況 2：只有一個子節點
		//如果nodeToDelete只有一個子節點，那麽在刪除nodeToDelete後，需要將其唯一的子節點連接到nodeToDelete的父節點上，以保持樹的連續性。
		//根據nodeToDelete是其父節點的左子節點還是右子節點，相應地更新parent.Left或parent.Right。
		//如果nodeToDelete是根節點（即parent為nil），則將BST的根節點更新為nodeToDelete的子節點。
	} else if nodeToDelete.Left == nil || nodeToDelete.Right == nil {
		var child *Node[T]
		if nodeToDelete.Left != nil {
			child = nodeToDelete.Left
		} else {
			child = nodeToDelete.Right
		}

		if parent == nil { // 要删除的是根节点
			bst.Root = child
		} else if parent.Left == nodeToDelete {
			parent.Left = child
		} else {
			parent.Right = child
		}
	} else {
		//有兩個child pointer
		//須先從「欲刪除之節點」的左子樹中尋找最大值或是右子樹中尋找最小值作為取代點(假設此節點為X)，來取代「欲刪除之節點」的data，再來對於X進行刪除的動作。
		successor, succParent := bst.findMin(nodeToDelete.Right, nodeToDelete)
		*nodeToDelete.Value = *successor.Value // 替换值
		bst.deleteNode(successor, succParent)  // 删除后继节点
	}
}

// findMin 找到以node為根的右子數中的最小node及其parent node
func (bst *BinarySearchTree[T]) findMin(node, parent *Node[T]) (*Node[T], *Node[T]) {
	current := node
	for current.Left != nil {
		parent = current
		current = current.Left
	}
	return current, parent
}

func (bst *BinarySearchTree[T]) BreadthFirstSearch() {
	var list []*Node[T]
	var queue []*Node[T]
	queue = append(queue, bst.Root)
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		list = append(list, currentNode)
		if currentNode.Left != nil {
			queue = append(queue, currentNode.Left)
		}
		if currentNode.Right != nil {
			queue = append(queue, currentNode.Right)
		}
	}
	for _, node := range list {
		fmt.Println(*node.Value)
	}
}

func (bst *BinarySearchTree[T]) DFSInOrder() {
	list := []*Node[T]{}
	list = traverseInOrder(bst.Root, list) // 接收返回的list
	for _, v := range list {
		fmt.Print(*v.Value, " ")
	}
	fmt.Println() // 输出换行，美化打印结果
}

func (bst *BinarySearchTree[T]) DFSPreOrder() {
	list := []*Node[T]{}
	list = traversePreOrder(bst.Root, list) // 接收返回的list
	for _, v := range list {
		fmt.Print(*v.Value, " ")
	}
	fmt.Println() // 输出换行，美化打印结果

}

func (bst *BinarySearchTree[T]) DFSPostOrder() {
	list := []*Node[T]{}
	list = traversePostOrder(bst.Root, list) // 接收返回的list
	for _, v := range list {
		fmt.Print(*v.Value, " ")
	}
	fmt.Println() // 输出换行，美化打印结果
}

func traverseInOrder[T NumberType](root *Node[T], list []*Node[T]) []*Node[T] {
	if root != nil {
		if root.Left != nil {
			list = traverseInOrder(root.Left, list)
		}
		list = append(list, root)
		if root.Right != nil {
			list = traverseInOrder(root.Right, list)
		}
	}
	return list
}

func traversePreOrder[T NumberType](root *Node[T], list []*Node[T]) []*Node[T] {
	if root != nil {
		list = append(list, root)
		if root.Left != nil {
			list = traversePreOrder(root.Left, list)
		}
		if root.Right != nil {
			list = traversePreOrder(root.Right, list)
		}
	}
	return list
}
func traversePostOrder[T NumberType](root *Node[T], list []*Node[T]) []*Node[T] {
	if root != nil {
		if root.Left != nil {
			list = traversePostOrder(root.Left, list)
		}
		if root.Right != nil {
			list = traversePostOrder(root.Right, list)
		}
		list = append(list, root)
	}
	return list
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func isValidBST(root *TreeNode) bool {
//
//}
