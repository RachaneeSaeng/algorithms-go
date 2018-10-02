package main

import (
	"../lnklist"
)

func findLoopBeginning(head *lnklist.Node) *lnklist.Node {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // dollision spot (loopsize - k or k steps to loop start)
			break
		}
	}

	// no loop (if it is loop next loop never be null)
	if fast == nil || fast.Next == nil {
		return nil
	}

	// if slow and fast move the same pace, they will meet at the loop start after k steps
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}

	return slow
}

func main() {

}
