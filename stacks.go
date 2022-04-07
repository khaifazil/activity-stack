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
	if p.top == nil { //checks to see if there are any nodes in Stack
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
	if p.size == 1 { //p.top.next == nil (if no size is specified)
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

func (p *stack) peek() (string, error) {
	if p.top == nil {
		return "", errors.New("Stack is empty.")
	}

	var item string = p.top.item
	return item, nil
}

func (p *stack) printAllNodes() error {
	currentNode := p.top
	if currentNode == nil {
		fmt.Println("Stack is empty.")
		return nil // minimize checks. every check is an operation
	}
	fmt.Printf("%+v\n", currentNode.item)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", currentNode.item)
	}

	// for currentNode != nil { //different way to printAllNodes
	// 	fmt.Printf("%+v\n", currentNode.item)
	// 	currentNode = currentNode.next
	// }
	return nil
}

func (p *stack) isEmpty() bool {
	return p.size == 0
}

func main() {
	myStack := &stack{nil, 0} //init stack

	myStack.push("Peter")
	myStack.push("Eddie")
	myStack.push("Mary Jane")
	myStack.push("SpooderMan")

	// fmt.Println("\nPrinting all the nodes in the stack...")
	// myStack.printAllNodes()

	// poppedItem, _ := myStack.pop()

	// fmt.Println("\nPrinting all the nodes in the stack...")
	// myStack.printAllNodes()
	// fmt.Printf("\n%v was popped", poppedItem)

	// fmt.Println("\nPeeking at the top node in the stack...")
	// peekItem, _ := myStack.peek()
	// fmt.Printf(peekItem)

	//print without using printAllNodes()


	tempStack := &stack{nil, 0}
	for myStack.isEmpty() == false {
		item, _ := myStack.pop()
		tempStack.push(item)
		fmt.Println(item)
	}
	
	for tempStack.isEmpty() == false {
		item, _ := tempStack.pop()
		myStack.push(item)
	}
}
