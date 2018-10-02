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

func (node *Node) InsertBefore(data int) *Node {
	return &Node{Data: data, Next: node}
}

func (node *Node) InsertAfter(data int) *Node {
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
