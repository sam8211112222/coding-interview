package Tree

// NumberType 要自己宣告泛型裡面的type, 不然如果用any, comparable 都沒辦法編譯
// comparable 约束在 Go 泛型中用于确保类型可以进行等值比较（即 == 和 != 操作），这对于很多泛型数据结构和函数来说是足够的。
// 然而，comparable 并不保证类型支持顺序比较操作（如 <、>、<=、>=）。这是因为 comparable 约束的设计目的是为了确保类型可以用在需要等值比较的场合，比如作为 map 的键。
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

func (bst *BinarySearchTree[T]) Search(value T) *T {
	if bst.Root == nil {
		return nil
	}
	return bst.Root.SearchNode(value)
}

func (n *Node[T]) SearchNode(value T) *T {
	if *n.Value == value {
		return n.Value
	}
	if value < *n.Value {
		if n.Left != nil {
			return n.Left.SearchNode(value)
		}
		return nil
	} else {
		if n.Right != nil {
			return n.Right.SearchNode(value)
		}
		return nil
	}
}
