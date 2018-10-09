package queue

import (
	"errors"

	"../../linkedlist/lnklist"
)

type Queue struct {
	first *lnklist.Node
	last  *lnklist.Node
}

func (stk *Queue) Add(n interface{}) {
	newNode := lnklist.New(n, nil)
	if stk.last != nil {
		stk.last.Next = newNode
	}
	stk.last = newNode
	if stk.first == nil {
		stk.first = stk.last
	}
}

func (stk *Queue) Remove() (interface{}, error) {
	if stk.first != nil {
		data := stk.first.Data
		stk.first = stk.first.Next
		return data, nil
	}
	return 0, errors.New("Empty queue")
}

func (stk *Queue) Peek() (interface{}, error) {
	if stk.first != nil {
		return stk.first.Data, nil
	}
	return 0, errors.New("Empty queue")
}

func (stk *Queue) IsEmpty() bool {
	return stk.first == nil
}

func main() {

}
