package lnklist

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
