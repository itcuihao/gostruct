package gostruct

import "fmt"

// Node node
type Node struct {
	value interface{}
}

func (n *Node) String() string {
	return fmt.Sprint(n.value)
}

// Stack stack
type Stack struct {
	Nodes []*Node
	Count int
}

// Push push
func (s *Stack) Push(i interface{}) {
	n := &Node{
		value: i,
	}
	nodes := make([]*Node, 0, s.Count+1)
	nodes = append(s.Nodes[:s.Count], n)

	s.Nodes = nodes
	s.Count++
	return
}

// Pop pop
func (s *Stack) Pop() *Node {
	if s.Count == 0 {
		return nil
	}

	s.Count--
	return s.Nodes[s.Count]
}
