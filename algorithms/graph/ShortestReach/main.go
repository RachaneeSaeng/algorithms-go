package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Data interface{}
	Next *Node
}

func New(data interface{}, next *Node) *Node {
	return &Node{
		Data: data,
		Next: next,
	}
}

func (node *Node) Length() int {
	if node == nil {
		return 0
	}
	return 1 + node.Next.Length()
}

func (node *Node) InsertBefore(data interface{}) *Node {
	return &Node{Data: data, Next: node}
}

func (node *Node) InsertAfter(data interface{}) *Node {
	newNode := &Node{Data: data, Next: nil}
	node.Next = newNode
	return node
}

func (node *Node) PadLeft(padding int) *Node {
	for i := 0; i < padding; i++ {
		node = node.InsertBefore(0)
	}
	return node
}

func (node *Node) PadRight(padding int) *Node {
	for i := 0; i < padding; i++ {
		node = node.InsertAfter(0)
	}
	return node
}

type Queue struct {
	first *Node
	last  *Node
}

func (stk *Queue) Add(n interface{}) {
	newNode := New(n, nil)
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

func (stk *Queue) Length() int {
	return stk.first.Length()
}

type node struct {
	nodeNum int32
	level   int32
}

// Complete the bfs function below.
func bfs(n int32, m int32, edges [][]int32, s int32) []int32 {
	adjacentList := buildAdjacentList(edges)
	result := make([]int32, n+1)
	q := &Queue{}
	q.Add(node{s, 0})
	result[s] = -1

	for !q.IsEmpty() {
		iCurrent, _ := q.Remove()
		current, _ := iCurrent.(node)
		result[current.nodeNum] = current.level

		level := current.level + 1
		if adjList, exist := adjacentList[current.nodeNum]; exist {
			for _, nodeNum := range adjList {
				if nodeNum != s && result[nodeNum] == 0 { //never visited, except the first node (result of first node = 0)
					q.Add(node{nodeNum, level})
					result[nodeNum] = -1
				}
			}
		}
	}

	result = removeElement(result, 0)
	result = removeElement(result, s-1)

	for i := range result {
		if result[i] == 0 {
			result[i] = -1
		} else {
			result[i] = result[i] * 6
		}
	}

	return result
}

func removeElement(arr []int32, index int32) []int32 {
	if index+1 < int32(len(arr)) {
		arr = append(arr[:index], arr[index+1:]...)
	} else {
		arr = append(arr[:index])
	}
	return arr
}
func buildAdjacentList(edges [][]int32) map[int32][]int32 {
	adjacentList := make(map[int32][]int32)
	for _, e := range edges {
		adjacentList[e[0]] = append(adjacentList[e[0]], e[1])
		adjacentList[e[1]] = append(adjacentList[e[1]], e[0])
	}
	return adjacentList
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		nm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(nm[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		mTemp, err := strconv.ParseInt(nm[1], 10, 64)
		checkError(err)
		m := int32(mTemp)

		var edges [][]int32
		for i := 0; i < int(m); i++ {
			edgesRowTemp := strings.Split(readLine(reader), " ")

			var edgesRow []int32
			for _, edgesRowItem := range edgesRowTemp {
				edgesItemTemp, err := strconv.ParseInt(edgesRowItem, 10, 64)
				checkError(err)
				edgesItem := int32(edgesItemTemp)
				edgesRow = append(edgesRow, edgesItem)
			}

			if len(edgesRow) != int(2) {
				panic("Bad input")
			}

			edges = append(edges, edgesRow)
		}

		sTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		s := int32(sTemp)

		result := bfs(n, m, edges, s)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
