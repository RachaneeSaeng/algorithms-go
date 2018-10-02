package linkedlist

type Node struct {
	Data int
	Next *Node
}

func (node *Node) Length() int {
	if node == nil {
		return 0
	}
	return 1 + node.Next.Length()
}
