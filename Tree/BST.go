package Tree

type Node[T any] struct {
	Value T
	Right *Node[T]
	Left  *Node[T]
}

type BinarySearchTree[T any] struct {
	Root   *Node[T]
	Length int
}

func NewBinarySearchTree[T any]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}
