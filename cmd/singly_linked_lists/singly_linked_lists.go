package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (list *LinkedList) Display() {
	current := list.head

	if current == nil {
		fmt.Println("Linked list is empty")
		return
	}

	fmt.Print("Linked list: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}

	list.Display()
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Display()

}
