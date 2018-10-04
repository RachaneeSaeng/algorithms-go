package main

import (
	"errors"

	"../../linkedlist/lnklist"
)

type stack struct {
	top *lnklist.Node
}

func (stk *stack) push(n int) {
	newNode := lnklist.New(n, stk.top)
	stk.top = newNode
}

func (stk *stack) pop() (int, error) {
	if stk.top != nil {
		data := stk.top.Data
		stk.top = stk.top.Next
		return data, nil
	}
	return 0, errors.New("Empty stack")
}

func (stk *stack) peek() (int, error) {
	if stk.top != nil {
		return stk.top.Data, nil
	}
	return 0, errors.New("Empty stack")
}

func (stk *stack) isEmpty() bool {
	return stk.top == nil
}

func main() {

}
