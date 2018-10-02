package main

import (
	"../lnklist"
)

// partition linked list around a value x
// such that all nodes < x come before all nodes >= x
func partition(node *lnklist.Node, x int) *lnklist.Node {
	if node == nil {
		return nil
	}

	head := node
	tail := node
	for node != nil {
		next := node.Next
		if node.Data < x {
			node.Next = head
			head = node
		} else {
			tail.Next = node
			tail = node
		}
		node = next
	}
	tail.Next = nil

	return head
}

func main() {

}
