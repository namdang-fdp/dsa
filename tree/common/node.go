package common

type Node struct {
	Data     int
	Children []*Node
}

func NewNode(x int) *Node {
	return &Node{
		Data:     x,
		Children: []*Node{},
	}
}
