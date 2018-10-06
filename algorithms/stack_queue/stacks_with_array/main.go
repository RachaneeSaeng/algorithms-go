package main

import (
	"errors"
)

type stacks struct {
	stackCount int
	tops       []int
	arr        []int
}

func new(stackCount int) *stacks {
	tops := make([]int, stackCount)
	arr := make([]int, stackCount)
	for i := 0; i < stackCount; i++ {
		tops[i] = -1
	}
	return &stacks{stackCount: stackCount, tops: tops, arr: arr}
}

func (stks *stacks) push(stackNum int, n int) error {
	if stackNum > len(stks.tops) {
		return errors.New("Stack number is out of range")
	}

	topIdx := stks.tops[stackNum-1]
	if topIdx < 0 {
		topIdx = stackNum - 1
	} else {
		topIdx += stks.stackCount
	}

	if topIdx >= len(stks.arr) {
		paddingValues := make([]int, topIdx-len(stks.arr)+1)
		stks.arr = append(stks.arr, paddingValues...)
	}
	stks.arr[topIdx] = n
	stks.tops[stackNum-1] = topIdx
	return nil
}

func (stks *stacks) pop(stackNum int) (int, error) {
	if stackNum > len(stks.tops) {
		return 0, errors.New("Stack number is out of range")
	}

	topIdx := stks.tops[stackNum-1]
	if topIdx < 0 {
		return 0, errors.New("Empty stack")
	}
	if topIdx >= len(stks.arr) {
		return 0, errors.New("Index out of range")
	}

	stks.tops[stackNum-1] = topIdx - stks.stackCount
	return stks.arr[topIdx], nil
}

func (stks *stacks) peek(stackNum int) (int, error) {
	if stackNum > len(stks.tops) {
		return 0, errors.New("Stack number is out of range")
	}

	topIdx := stks.tops[stackNum-1]
	if topIdx < 0 {
		return 0, errors.New("Empty stack")
	}
	if topIdx >= len(stks.arr) {
		return 0, errors.New("Index out of range")
	}

	return stks.arr[topIdx], nil
}

func (stks *stacks) isEmpty(stackNum int) (bool, error) {
	if stackNum > len(stks.tops) {
		return false, errors.New("Stack number is out of range")
	}

	topIdx := stks.tops[stackNum-1]
	return topIdx < 0, nil
}

func main() {

}
