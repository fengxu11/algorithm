package example

import "fmt"

// Node 节点
type Node struct {
	data interface{}
	next *Node
}

// NewNode 创建node
func NewNode(data interface{}, next *Node) *Node {
	return &Node{
		data: data,
		next: next,
	}
}

func (s *Node) String() string {
	return fmt.Sprint(s.data)
}
