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

func (dll *DoublyLinkedList[T]) FindByIndex(index int) (*DoublyLinkedListNode[T], error) {
	err := fmt.Errorf("index: %v out of range", index)
	if index < 0 || index > dll.Length-1 {
		return nil, err
	}
	var currentIndex int
	// 如果index比一半的length大, 從後往前找比較快
	if index > dll.Length/2 {
		currentIndex = dll.Length - 1
		currentNode := dll.Tail
		for currentNode != nil {
			if currentIndex == index {
				return currentNode, nil
			}
			currentNode = currentNode.Prev
			currentIndex--
		}
	} else {
		currentNode := dll.Head
		for currentNode != nil {
			if currentIndex == index {
				return currentNode, nil
			}
			currentNode = currentNode.Next
			currentIndex++
		}
	}
	return nil, err
}

func (dll *DoublyLinkedList[T]) DeleteAtIndex(index int) error {
	nodeByIndex, err := dll.FindByIndex(index)
	if err != nil {
		return err
	}

	if index == 0 {
		dll.Head = nodeByIndex.Next
		if dll.Head != nil {
			dll.Head.Prev = nil
		} else {
			dll.Tail = nil
		}
	} else if index == dll.Length-1 {
		dll.Tail = nodeByIndex.Prev
		if dll.Tail != nil {
			dll.Tail.Next = nil
		} else {
			dll.Head = nil
		}
	} else {
		prev := nodeByIndex.Prev
		next := nodeByIndex.Next
		prev.Next = next
		next.Prev = prev
	}
	dll.Length--
	return nil
}

// InsertAtIndex 容許的index範圍只從0~length-1
func (dll *DoublyLinkedList[T]) InsertAtIndex(index int, value T) error {
	nodeByIndex, err := dll.FindByIndex(index)
	if err != nil {
		return err
	}
	newNode := &DoublyLinkedListNode[T]{Value: value}

	if index == 0 {
		nextNode := dll.Head
		dll.Head = newNode
		newNode.Next = nextNode
		if nextNode != nil {
			nextNode.Prev = newNode
		} else {
			dll.Tail = newNode
		}
	} else {
		prev := nodeByIndex.Prev
		next := nodeByIndex
		prev.Next = newNode
		newNode.Prev = prev
		next.Prev = newNode
		newNode.Next = next
	}
	dll.Length++
	return nil
}

func (dll *DoublyLinkedList[T]) Reverse() {
	var prevNode *DoublyLinkedListNode[T]
	currentNode := dll.Head
	dll.Tail = currentNode
	for currentNode != nil {
		prevNode = currentNode.Prev
		nextNode := currentNode.Next
		currentNode.Prev = nextNode
		currentNode.Next = prevNode
		if nextNode == nil {
			dll.Head = currentNode
			return
		} else {
			currentNode = nextNode
		}
	}
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
