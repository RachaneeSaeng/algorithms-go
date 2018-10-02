package main

import (
	"../linkedlist"
)

// Get kth to the last node
// by let p1 go Kth times first and then p1 and p2 go with the same pace untile p1 reach the end
// that means p2 is the Kth to the last node
func kthToTheLast(head *linkedlist.Node, k int) *linkedlist.Node {
	pointer1, pointer2 := head, head

	for i := 0; i < k; i++ {
		if pointer1 == nil {
			return nil
		}
		pointer1 = pointer1.Next
	}

	for pointer1 != nil {
		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
	}
	return pointer2
}

// Get nth to the last node
func nthToTheLast(head *linkedlist.Node, k int) *linkedlist.Node {
	var i int
	return nthToTheLastWithIndex(head, k, &i)
}

// Get nth to the last node with recusive approach
func nthToTheLastWithIndex(head *linkedlist.Node, k int, i *int) *linkedlist.Node {
	if head == nil {
		return nil
	}
	nd := nthToTheLastWithIndex(head.Next, k, i)
	*i = *i + 1
	if *i == k {
		return head
	}
	return nd
}

func main() {

}
