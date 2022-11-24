package main

import (
	"fmt"

	"github.com/Lunovoy/linked-list/linked_list"
)

func main() {
	fmt.Println("=========SinglyLinkedList=========")
	sList := linked_list.SinglyLinkedList{}

	sList.Prepend(6)
	sList.Prepend(5)
	sList.Prepend(4)
	sList.Prepend(3)
	sList.Prepend(2)
	sList.Prepend(1)
	sList.Append(777)

	sList.PrintList()

	sList.DeleteWithValue(3)

	sList.PrintList()

	sList.InsertAfterValue(5, 444)

	sList.PrintList()

	sList.InsertAfterValue(777, 444)

	sList.PrintList()

	fmt.Println("Head: ", sList.Head)
	fmt.Println("Tail: ", sList.Tail)
	fmt.Println("Length: ", sList.Length)

	fmt.Println("==================================")
	fmt.Println("=========DoublyLinkedList=========")

	dList := linked_list.DoublyLinkedList{}

	dList.Prepend(5)
	dList.Prepend(4)
	dList.Prepend(3)
	dList.Prepend(2)
	dList.Prepend(1)

	dList.Append(666)

	dList.PrintDoublyList()

	dList.InsertAfterValue(2, 333)

	dList.PrintDoublyList()

	dList.InsertAfterValue(666, 777)

	dList.InsertBeforeValue(2, 111)

	dList.PrintDoublyList()

	dList.InsertBeforeValue(1, 100)

	dList.PrintDoublyList()

	fmt.Println("Head: ", dList.Head)
	fmt.Println("Tail: ", dList.Tail)
	fmt.Println("Length: ", dList.Length)

	fmt.Println("==================================")

}
