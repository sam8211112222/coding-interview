package LinkedList

import (
	"fmt"
	"strings"
)

type SinglyNode[T any] struct {
	Value T
	Next  *SinglyNode[T]
}

type SinglyLinkedList[T any] struct {
	Head   *SinglyNode[T]
	Length int
}

func NewSinglyLinkedList[T any]() *SinglyLinkedList[T] {
	return &SinglyLinkedList[T]{}
}

func (sll *SinglyLinkedList[T]) Append(value T) {
	newNode := &SinglyNode[T]{Value: value, Next: nil}

	if sll.Head == nil {
		sll.Head = newNode
	} else {
		current := sll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	sll.Length++
}

func (sll *SinglyLinkedList[T]) Prepend(value T) {
	newNode := &SinglyNode[T]{Value: value, Next: sll.Head}
	sll.Head = newNode
	sll.Length++
}

func (sll SinglyLinkedList[T]) FindByIndex(index int) (*SinglyNode[T], error) {
	currentIndex := 0
	current := sll.Head
	err := fmt.Errorf("index: %v out of range", index)
	if index < 0 || index >= sll.Length {
		return nil, err
	}
	for current != nil {
		if index == currentIndex {
			return current, nil
		}
		currentIndex++
		current = current.Next
	}
	return nil, err
}

func (sll *SinglyLinkedList[T]) InsertAtIndex(index int, value T) error {
	newNode := &SinglyNode[T]{Value: value}
	err := fmt.Errorf("index: %v out of range", index)
	if index < 0 || index >= sll.Length {
		return err
	}
	if index == 0 {
		prevNode := sll.Head
		sll.Head = newNode
		newNode.Next = prevNode
		return nil
	}

	nodeByIndex, err := sll.FindByIndex(index - 1)
	if err != nil {
		return err
	}
	next := nodeByIndex.Next
	nodeByIndex.Next = newNode
	newNode.Next = next
	sll.Length++
	return nil
}

func (sll *SinglyLinkedList[T]) DeleteAtIndex(index int) error {

	if index < 0 || index >= sll.Length {
		err := fmt.Errorf("index: %v out of range", index)
		return err
	}

	if index == 0 {
		next := sll.Head.Next
		sll.Head = next
		return nil
	}

	nodeByIndex, err := sll.FindByIndex(index - 1)
	if err != nil {
		return err
	}
	prev := nodeByIndex
	next := nodeByIndex.Next.Next
	prev.Next = next
	sll.Length--
	return nil
}

func (sll *SinglyLinkedList[T]) Reverse() {
	var prevNode *SinglyNode[T]
	currentNode := sll.Head
	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = prevNode
		prevNode = currentNode
		currentNode = nextNode
	}
	sll.Head = prevNode
}

func (sll *SinglyLinkedList[T]) PrintList() {
	var sb strings.Builder
	sb.WriteString("[")

	current := sll.Head
	for current != nil {
		sb.WriteString(fmt.Sprintf("%v", current.Value))
		if current.Next != nil {
			sb.WriteString(", ")
		}
		current = current.Next
	}
	sb.WriteString("]")
	fmt.Println(sb.String())
}
