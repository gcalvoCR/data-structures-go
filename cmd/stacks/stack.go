package main

import "fmt"

// Stack structure using a slice
type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
	fmt.Println("element added:", item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		fmt.Println("The stack is empty")
		return -1
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	stack := Stack{}

	// Pushing elements to the stack
	stack.Push(2)
	stack.Push(4)
	stack.Push(6)

	// Popping elements from the stack
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
