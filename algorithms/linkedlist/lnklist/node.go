package lnklist

type Node struct {
	Data int
	Next *Node
}

func New(data int, next *Node) *Node {
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
