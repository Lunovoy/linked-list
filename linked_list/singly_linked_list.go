package linked_list

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type SinglyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l *SinglyLinkedList) Prepend(value int) {
	newNode := Node{Value: value}
	if l.Head != nil {
		newNode.Next = l.Head
		l.Head = &newNode
		l.Length++
		return
	}
	l.Head = &newNode
	l.Tail = &newNode
	l.Length++

}

func (l *SinglyLinkedList) Append(value int) {
	newNode := Node{Value: value}
	tail := l.Tail
	tail.Next = &newNode
	l.Tail = &newNode
	l.Length++
}

func (l *SinglyLinkedList) InsertAfterValue(givenValue, newValue int) {

	if l.Tail.Value == givenValue {
		l.Append(newValue)
		return
	}

	currentNode := l.Head
	newNode := Node{Value: newValue}

	for currentNode != nil {
		if currentNode.Value == givenValue {
			newNode.Next = currentNode.Next
			currentNode.Next = &newNode
			l.Length++
			fmt.Printf("Node with value %v inserted after %v\n", newValue, givenValue)
			return
		}
		currentNode = currentNode.Next
	}

	fmt.Printf("Given value %v not found\n", givenValue)
}

func (l SinglyLinkedList) PrintList() {
	if l.Length == 0 {
		fmt.Println("Length is 0")
		return
	}

	currentNode := l.Head
	for currentNode != nil {
		fmt.Printf("%v -> ", currentNode)
		currentNode = currentNode.Next
	}
	fmt.Println("Length is ", l.Length)
}

func (l *SinglyLinkedList) DeleteWithValue(value int) {
	if l.Length == 0 {
		fmt.Println("Error occured while deliting List with 0 length")
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		l.Length--
		fmt.Printf("First Node with value %v Deleted\n", value)
		return
	}

	currentNode := l.Head

	for currentNode != nil && currentNode.Next != nil {
		if currentNode.Next.Value == value {
			fmt.Printf("Node %v Deleted\n", currentNode.Next)
			currentNode.Next = currentNode.Next.Next
			l.Length--
			return
		}
		currentNode = currentNode.Next
	}
}
