package main

import (
	"../lnklist"
)

type lstInfo struct {
	tail *lnklist.Node
	size int
}

func findIntersaction(l1 *lnklist.Node, l2 *lnklist.Node) *lnklist.Node {
	if l1 == nil || l2 == nil {
		return nil
	}

	info1 := getTailAndSize(l1)
	info2 := getTailAndSize(l2)

	// if the last items is not the same, that means they are not intesect
	if info1.tail != info2.tail {
		return nil
	}

	if info1.size > info2.size {
		l1 = getKthNode(l1, info1.size-info2.size)
	} else if info2.size > info1.size {
		l2 = getKthNode(l2, info2.size-info1.size)
	}

	for l1 != l2 {
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1
}

func getTailAndSize(head *lnklist.Node) lstInfo {
	i := 1
	tail := head
	for tail.Next != nil {
		i++
		tail = tail.Next
	}
	return lstInfo{tail: tail, size: i}
}

func getKthNode(node *lnklist.Node, k int) *lnklist.Node {
	for i := 0; i < k; i++ {
		node = node.Next
	}
	return node
}

func main() {

}
