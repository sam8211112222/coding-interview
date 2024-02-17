package LinkedList

func RunLinkedList() {
	DllRun()
}

func SllRun() {
	sll := NewSinglyLinkedList[int]()
	sll.Append(1)
	sll.PrintList()
	sll.Append(2)
	sll.PrintList()
	sll.Prepend(0)
	sll.PrintList()
	sll.InsertAtIndex(1, 11)
	sll.PrintList()
	sll.DeleteAtIndex(1)
	sll.PrintList()
	sll.DeleteAtIndex(3)
	sll.PrintList()
}

func DllRun() {
	dll := NewDoublyLinkedList[int]()
	dll.PrintList()
	dll.Append(1)
	dll.PrintList()
	dll.Append(2)
	dll.PrintList()
	dll.Prepend(0)
	dll.PrintList()
}
