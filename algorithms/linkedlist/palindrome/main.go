package main

import (
	"../lnklist"
)

type result struct {
	matched bool
	node    *lnklist.Node
}

// check if the linkedlist is palindrome
// 2 - 1 - 1 - 2 = true
// 2 - 1 - 1 - 1 = false
func isPalindrome(node *lnklist.Node) bool {
	res := isPalindromeRecures(node, node.Length())
	return res.matched
}

// start from second half and compare out to the end
func isPalindromeRecures(node *lnklist.Node, len int) result {
	if node == nil || len <= 0 {
		return result{matched: true, node: node}
	} else if len == 1 { // Odd length
		return result{matched: true, node: node.Next} // Skip the middle node
	}

	lastResult := isPalindromeRecures(node.Next, len-2) // len will reach 0 or 1 when we are at the center of the list because len reduced by 2
	if !lastResult.matched || lastResult.node == nil {
		return lastResult
	}

	lastResult.matched = lastResult.node.Data == node.Data
	lastResult.node = lastResult.node.Next

	return lastResult
}
func main() {

}
