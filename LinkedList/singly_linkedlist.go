package LinkedList

import "fmt"

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
	err := fmt.Errorf("index: %v out of range: %v", index, sll.Length)
	if index >= sll.Length {
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
