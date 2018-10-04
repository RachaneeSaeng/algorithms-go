package main

import (
	"errors"

	"../../linkedlist/lnklist"
)

type queue struct {
	first *lnklist.Node
	last  *lnklist.Node
}

func (stk *queue) add(n int) {
	newNode := lnklist.New(n, nil)
	if stk.last != nil {
		stk.last.Next = newNode
	}
	stk.last = newNode
	if stk.first == nil {
		stk.first = stk.last
	}
}

func (stk *queue) remove() (int, error) {
	if stk.first != nil {
		data := stk.first.Data
		stk.first = stk.first.Next
		return data, nil
	}
	return 0, errors.New("Empty queue")
}

func (stk *queue) peek() (int, error) {
	if stk.first != nil {
		return stk.first.Data, nil
	}
	return 0, errors.New("Empty queue")
}

func (stk *queue) isEmpty() bool {
	return stk.first == nil
}

func main() {

}
