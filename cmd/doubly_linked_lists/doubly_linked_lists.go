package main

import (
	"fmt"

	"github.com/fatih/color"
)

type Node struct {
	data int
	prev *Node
	next *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

// Insert at head
func (list *LinkedList) InsertAtBeginning(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	// link the new node
	newNode.next = list.head
	list.head.prev = newNode

	// assign the new node
	list.head = newNode
}

// Insert at tail
func (list *LinkedList) InsertAtEnd(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		return
	}

	newNode.prev = list.tail
	list.tail.next = newNode
	list.tail = newNode
}

func (list *LinkedList) TraverseForward() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		prev := -1
		if current.prev != nil {
			prev = current.prev.data
		}
		fmt.Printf(" <-[ %d ", prev)
		fmt.Printf("%s", color.GreenString(fmt.Sprintf("| %d |", current.data)))
		next := -1
		if current.next != nil {
			next = current.next.data
		}
		fmt.Printf(" %d ]-> ", next)

		current = current.next
	}
	fmt.Println()
}

func main() {

	dll := LinkedList{}
	dll.InsertAtBeginning(10)
	dll.InsertAtBeginning(20)
	dll.InsertAtBeginning(30)
	dll.InsertAtBeginning(40)
	dll.TraverseForward()

	dll2 := LinkedList{}
	dll2.InsertAtEnd(10)
	dll2.InsertAtEnd(20)
	dll2.InsertAtEnd(30)
	dll2.InsertAtEnd(40)
	dll2.TraverseForward()

}
