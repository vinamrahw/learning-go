package main

import "fmt"

// Creating a structure for linked list
type LinkedList struct {
	name string
	next *LinkedList
}
//Add data to list
func (a *LinkedList) FillList(nm string) {
	nxt := &LinkedList{
		name: nm,
		next: nil,
	}
	// 1st Record to list
	if a.next == nil && a.name == "" {
		a.name = nm
	// 2nd Record to list
	} else if a.next == nil {
		a.next = nxt
	} else {
		curr := a.next
		fmt.Println(curr.next)
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = nxt
		fmt.Println("cur", curr.next)
	}
}
//Print data in the list
func (a *LinkedList) printList() {
	if a.next == nil {
		fmt.Println("Empty list")
	} else {
		fmt.Println(a.name)
		curr := a.next
		for curr.next != nil {
			fmt.Println(curr.name)
			curr = curr.next
		}
		fmt.Println(curr.name)
	}
}

// Main function
func main() {
	a := &LinkedList{}
	a.FillList("Data1")
	a.FillList("Data2")
	a.FillList("Data3")
	a.FillList("Data4")
	//Print the list
	a.printList()
}
