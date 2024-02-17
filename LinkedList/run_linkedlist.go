package LinkedList

func RunLinkedList() {
	//DllRun()
	//SllRevserse()
	DllReverse()
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
	dll.DeleteAtIndex(2)
	dll.DeleteAtIndex(1)
	dll.PrintList()
	dll.InsertAtIndex(0, 12345)
	dll.PrintList()
}

func SllRevserse() {
	sll := NewSinglyLinkedList[int]()
	sll.Append(1)
	sll.Append(2)
	sll.Append(3)
	sll.Reverse()
	sll.PrintList()
}

func DllReverse() {
	dll := NewDoublyLinkedList[int]()
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	dll.PrintList()
	dll.Reverse()
	dll.PrintList()

}
