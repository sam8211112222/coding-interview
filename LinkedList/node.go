package LinkedList

import "fmt"

type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// 定義一個泛型鏈表
type LinkedList[T any] struct {
	Head *Node[T]
}

// 添加元素到鏈表末尾
func (l *LinkedList[T]) Append(value T) {
	newNode := &Node[T]{Value: value}
	if l.Head == nil {
		l.Head = newNode
		return
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}

// 遍歷鏈表並打印所有元素
func (l *LinkedList[T]) Print() {
	current := l.Head
	for current != nil {
		fmt.Println(current.Value)
		current = current.Next
	}
}
