package main

import "fmt"

// Node represents an element in the linked list
type Node struct {
	data int
	next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
	head *Node
}

// Add adds a new node with the given data to the end of the list
func (ll *LinkedList) Add(data int) {
	newNode := &Node{data: data}
	if ll.head == nil {
		ll.head = newNode
		fmt.Printf("Added head: %d\n", data)
		return
	}
	current := ll.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
	fmt.Printf("Added: %d\n", data)
}

// Remove removes the node with the given data from the list
func (ll *LinkedList) Remove(data int) {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}
	if ll.head.data == data {
		fmt.Printf("Removed head: %d\n", data)
		ll.head = ll.head.next
		return
	}
	current := ll.head
	for current.next != nil && current.next.data != data {
		current = current.next
	}
	if current.next == nil {
		fmt.Printf("Node with data %d not found\n", data)
		return
	}
	fmt.Printf("Removed: %d\n", data)
	current.next = current.next.next
}

// Traverse prints all the elements in the list
func (ll *LinkedList) Traverse() {
	if ll.head == nil {
		fmt.Println("List is empty")
		return
	}
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	ll := &LinkedList{}

	ll.Add(10)
	ll.Add(20)
	ll.Add(30)
	ll.Traverse() // Output: 10 -> 20 -> 30 -> nil

	ll.Remove(20)
	ll.Traverse() // Output: 10 -> 30 -> nil

	ll.Remove(40) // Output: Node with data 40 not found
	ll.Traverse() // Output: 10 -> 30 -> nil

	ll.Remove(10)
	ll.Traverse() // Output: 30 -> nil

	ll.Remove(30)
	ll.Traverse() // Output: List is empty
}
