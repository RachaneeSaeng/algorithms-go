package main

import (
	"../lnklist"
)

// reverse linkedlist in group
// 1 > 2 > 3 > 4 > 5 and x = 2
// then result = 2 > 1 4 > 3 > 5
func reverseLinkedList(node *lnklist.Node, x int) *lnklist.Node {
	if node == nil {
		return nil
	}

	var firstNode *lnklist.Node
	var lastNode *lnklist.Node

	for node != nil {
		startNode, endNode, nextNode := reverse(node, x)

		if firstNode == nil && lastNode == nil {
			firstNode = startNode
			lastNode = endNode
		} else if lastNode != nil {
			lastNode.Next = startNode
			lastNode = endNode
		}
		node = nextNode
	}

	return firstNode
}

func reverse(node *lnklist.Node, x int) (*lnklist.Node, *lnklist.Node, *lnklist.Node) {
	if x <= 0 || node == nil {
		return nil, nil, node
	}

	endNode := node
	preNode := node
	nextNode := node.Next
	endNode.Next = nil

	for nextNode != nil && x > 1 {
		node = nextNode
		nextNode = node.Next
		node.Next = preNode
		preNode = node
		x--
	}

	return node, endNode, nextNode
}

// func reverse(node *Node, x int) *Node {
// 	if x == 1 || node == nil {
// 		return node
// 	}
//     halfNode := reverse(node, x/2)
//     res := (halfNode.Next, x-1)

// 	//res := reverse(node.Next, x-1)

//     lastNode := gotoNode(res, x-2)
// 	lastNode.Next = node
// 	node.Next = nil
// 	return res
// }

// func reverse(node *lnklist.Node, x int) *lnklist.Node {
// 	if x <= 0 || node == nil {
// 		return node
// 	}
// 	halfNode := reverse(node, x/2)
// 	temp := node
// 	node.Data = halfNode.Next.Data
// 	halfNode.Next.Data = temp.Data

// 	return node
// }

// func reverse(node *lnklist.Node, x int) *lnklist.Node {
// 	if x == 1 || node == nil {
// 		return node
// 	}
// 	res := reverse(node.Next, x-1)
// 	lastNode := gotoNode(res, x-2)
// 	lastNode.Next = node
// 	node.Next = nil
// 	return res
// }

func gotoNode(node *lnklist.Node, n int) *lnklist.Node {
	if node == nil {
		return nil
	}

	for i := 0; i < n; i++ {
		node = node.Next
	}

	return node
}

func main() {

}
