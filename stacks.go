package main

import (
	"errors"
	"fmt"
)

type Node struct {
	item string
	next *Node
}

type stack struct {
	top  *Node
	size int
}

func (p *stack) push(name string) error {
	newNode := &Node{
		item: name,
		next: nil,
	}
	if p.top == nil {
		p.top = newNode
	} else {
		newNode.next = p.top // points newNode to previous newNode's address
		p.top = newNode      // creates newNode ontop
	}
	p.size++
	return nil
}

func (p *stack) pop() (string, error) {
	var item string

	if p.top == nil {
		return "", errors.New("Stack is empty!")
	}
	
	item = p.top.item
	if p.size == 1 {
		p.top = nil
	} else {
		p.top = p.top.next
	}
	p.size--
	return item, nil
}

// func (p *stack) pop(num ...int) ([]Node, error) { // how to make pop more than one item?
// 	items := []string {}

// 	if p.top == nil {
// 		return nil, errors.New("Stack is empty!")
// 	}
	
// 	if num < 1 && p.size == 1 {
// 		items = append(items, p.top)
// 		p.top = nil
// 		return items[0], errors.New("Only one item in stack.")
// 	} else {
// 		items := []string {}
// 		for i:= 0 ; i <= num[0]; i++{
// 			items = append(items, p.top.item)
// 			p.top = p.top.next
// 		}
// 		return items, nil
// 	}
// 	p.size = p.size - num[0]
// }

func (p *stack) printAllNodes() error {
	currentNode := p.top
	if currentNode == nil {
		fmt.Println("Stack is empty.")
		return nil
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}
	return nil
}

func main() {
	myList := &stack{nil, 0} //init stack

	myList.push("Peter")
	myList.push("Parker")
	myList.push("Mary Jane")
	myList.push("SpooderMan")

	fmt.Println("\nPrinting all the nodes in the stack...")
	myList.printAllNodes()

	myList.pop()

	fmt.Println("\nPrinting all the nodes in the stack...")
	myList.printAllNodes()
}