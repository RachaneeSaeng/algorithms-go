package main

type node struct {
	data int32
	next *node
}

// Get kth to the last node
// by let p1 go Kth times first and then p1 and p2 go with the same pace untile p1 reach the end
// that means p2 is the Kth to the last node
func kthToTheLast(head *node, k int) *node {
	pointer1, pointer2 := head, head

	for i := 0; i < k; i++ {
		if pointer1 == nil {
			return nil
		}
		pointer1 = pointer1.next
	}

	for pointer1 != nil {
		pointer1 = pointer1.next
		pointer2 = pointer2.next
	}
	return pointer2
}

// Get nth to the last node
func nthToTheLast(head *node, k int) *node {
	var i int
	return nthToTheLastWithIndex(head, k, &i)
}

// Get nth to the last node with recusive approach
func nthToTheLastWithIndex(head *node, k int, i *int) *node {
	if head == nil {
		return nil
	}
	nd := nthToTheLastWithIndex(head.next, k, i)
	*i = *i + 1
	if *i == k {
		return head
	}
	return nd
}

func main() {

}
