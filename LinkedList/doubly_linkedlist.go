package LinkedList

import (
	"fmt"
	"strings"
)

type DoublyLinkedListNode[T any] struct {
	Value T
	Prev  *DoublyLinkedListNode[T]
	Next  *DoublyLinkedListNode[T]
}

type DoublyLinkedList[T any] struct {
	Head   *DoublyLinkedListNode[T]
	Tail   *DoublyLinkedListNode[T]
	Length int
}

func NewDoublyLinkedList[T any]() *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{}
}

func (dll *DoublyLinkedList[T]) Append(value T) {
	newNode := &DoublyLinkedListNode[T]{Value: value}
	if dll.Length == 0 {
		dll.Head = newNode
		dll.Tail = newNode
	} else {
		dll.Tail.Next = newNode
		newNode.Prev = dll.Tail
		dll.Tail = newNode
	}
	dll.Length++
}

func (dll *DoublyLinkedList[T]) Prepend(value T) {
	newNode := &DoublyLinkedListNode[T]{Value: value, Next: dll.Head}
	if dll.Length == 0 {
		dll.Tail = newNode
	} else {
		// 舊的head prev指針指向newNode
		dll.Head.Prev = newNode
	}
	// 把新的 head 指向 newNode
	dll.Head = newNode
	dll.Length++

	//比較繁瑣的寫法是這樣
	//newNode := &DoublyLinkedListNode[T]{Value: value}
	//if dll.Length == 0 {
	//	// 如果鏈表為空，則新節點既是頭節點也是尾節點
	//	dll.Head = newNode
	//	dll.Tail = newNode
	//} else {
	//	// 如果鏈表不為空，更新現有頭節點的Prev指針
	//	dll.Head.Prev = newNode
	//	// 將新節點的Next指針指向原頭節點
	//	newNode.Next = dll.Head
	//	// 將新節點設為頭節點
	//	dll.Head = newNode
	//}
	//dll.Length++
}

func (dll *DoublyLinkedList[T]) PrintList() {
	var sb strings.Builder
	sb.WriteString("[")
	current := dll.Head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.Value))
		current = current.Next
		if current != nil {
			sb.WriteString(", ")
		}
	}
	sb.WriteString("]")
	fmt.Println(sb.String())
}
