package LinkedList

func RunLinkedList() {
	sll := NewSinglyLinkedList[int]()
	sll.Append(1)
	sll.Print()
	sll.Append(2)
	sll.Print()
	sll.Prepend(0)
	sll.Print()
	sll.InsertAtIndex(1, 11)
	sll.Print()
	sll.DeleteAtIndex(1)
	sll.Print()
	sll.DeleteAtIndex(3)
	sll.Print()
}
