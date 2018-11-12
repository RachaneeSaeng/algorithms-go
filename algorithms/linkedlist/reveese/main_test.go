package main

import (
	"fmt"
	"testing"

	"../lnklist"
)

func Test_Partition_Multiple_MatchFirstItem(t *testing.T) {
	node := lnklist.New(3, lnklist.New(5, lnklist.New(8, nil)))
	//expectedResult := lnklist.New(5, lnklist.New(3, lnklist.New(8, nil)))

	result := reverseLinkedList(node, 2)

	for result != nil {
		fmt.Println(result.Data)
		result = result.Next
	}
}
