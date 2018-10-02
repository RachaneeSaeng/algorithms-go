package main

import (
	"../lnklist"
)

type partialSum struct {
	head  *lnklist.Node
	carry int
}

// sum data in 2 linkedlist and return another linkedlist
// (6 -> 1 -> 7) + (2-> 9 -> 5)
// (9 -> 1 -> 2)
func sumList(l1 *lnklist.Node, l2 *lnklist.Node) *lnklist.Node {
	// if l1 == nil && l2 == nil {
	// 	return nil
	// } else if l1 == nil {
	// 	return l2
	// } else if l2 == nil {
	// 	return l1
	// }

	len1 := l1.Length()
	len2 := l2.Length()

	if len1 < len2 {
		l1 = l1.PadLeft(len2 - len1)
	} else if len2 < len1 {
		l2 = l2.PadLeft(len1 - len2)
	}

	sum := sum(l1, l2)
	if sum.carry > 0 {
		sum.head = sum.head.InsertBefore(sum.carry)
	}

	return sum.head
}

func sum(l1 *lnklist.Node, l2 *lnklist.Node) partialSum {
	if l1 == nil && l2 == nil {
		return partialSum{}
	}
	latestSum := sum(l1.Next, l2.Next)
	val := latestSum.carry + l1.Data + l2.Data

	latestSum.head = latestSum.head.InsertBefore(val % 10)
	latestSum.carry = val / 10
	return latestSum
}

func main() {

}
