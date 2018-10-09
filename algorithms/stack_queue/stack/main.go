package stack

import (
	"errors"

	"../../linkedlist/lnklist"
)

type Stack struct {
	top *lnklist.Node
}

func (stk *Stack) Push(n interface{}) {
	newNode := lnklist.New(n, stk.top)
	stk.top = newNode
}

func (stk *Stack) Pop() (interface{}, error) {
	if stk.top != nil {
		data := stk.top.Data
		stk.top = stk.top.Next
		return data, nil
	}
	return 0, errors.New("Empty stack")
}

func (stk *Stack) Peek() (interface{}, error) {
	if stk.top != nil {
		return stk.top.Data, nil
	}
	return 0, errors.New("Empty stack")
}

func (stk *Stack) IsEmpty() bool {
	return stk.top == nil
}

func main() {

}
