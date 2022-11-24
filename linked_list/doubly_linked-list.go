package linked_list

import "fmt"

type dNode struct {
	Value int
	Prev  *dNode
	Next  *dNode
}

type DoublyLinkedList struct {
	Head   *dNode
	Tail   *dNode
	Length int
}

func (dl *DoublyLinkedList) Prepend(value int) {
	newNode := dNode{Value: value}
	if dl.Head == nil {
		dl.Head = &newNode
		dl.Tail = &newNode
		dl.Length++
		return
	}
	head := dl.Head
	newNode.Next = head
	head.Prev = &newNode
	dl.Head = &newNode
	dl.Length++
}

func (dl *DoublyLinkedList) Append(value int) {
	newNode := dNode{Value: value}
	if dl.Head == nil {
		dl.Head = &newNode
		dl.Tail = &newNode
		dl.Length++
		return
	}

	tail := dl.Tail
	newNode.Prev = tail
	tail.Next = &newNode
	dl.Tail = &newNode
	dl.Length++
}

func (dl *DoublyLinkedList) InsertAfterValue(givenValue, newValue int) {

	if dl.Length == 0 {
		fmt.Printf("List length is 0\n")
		return
	}

	if dl.Tail.Value == givenValue {
		dl.Append(newValue)
		return
	}

	newNode := dNode{Value: newValue}
	currentNode := dl.Head

	for currentNode.Value != givenValue && currentNode.Next.Next != nil {
		currentNode = currentNode.Next
	}

	if currentNode.Value != givenValue {
		fmt.Printf("Given value %v not found\n", givenValue)
		return
	}

	newNode.Prev = currentNode
	newNode.Next = currentNode.Next
	newNode.Next.Prev = &newNode
	currentNode.Next = &newNode
	dl.Length++
}

func (dl *DoublyLinkedList) InsertBeforeValue(givenValue, newValue int) {
	if dl.Length == 0 {
		fmt.Printf("List length is 0\n")
		return
	}

	if dl.Head.Value == givenValue {
		dl.Prepend(newValue)
		return
	}

	newNode := dNode{Value: newValue}
	currentNode := dl.Head
	for currentNode.Next.Value != givenValue && currentNode.Next.Next.Next != nil {
		currentNode = currentNode.Next
	}

	if currentNode.Next.Value != givenValue {
		fmt.Printf("Given value %v not found\n", givenValue)
		return
	}

	newNode.Prev = currentNode
	newNode.Next = currentNode.Next
	newNode.Next.Prev = &newNode
	currentNode.Next = &newNode
	dl.Length++
}

func (dl DoublyLinkedList) PrintDoublyList() {
	if dl.Length == 0 {
		fmt.Println("Length is 0")
		return
	}

	currentNode := dl.Head
	for currentNode != nil {
		fmt.Printf("%v ->", currentNode)
		currentNode = currentNode.Next
	}
	fmt.Println("Length is ", dl.Length)
}
