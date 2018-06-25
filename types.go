package gostruct

import "fmt"

// Node node
type Node struct {
	value interface{}
}

func (n Node) String() string {
	return fmt.Sprint(n.value)
}

// NewNode new node
func NewNode(v interface{}) *Node {
	return &Node{
		value: v,
	}
}

// Nodes nodes
type Nodes struct {
	Value interface{}
	Next  *Nodes
}

// SingleLinkedList 单链表
type SingleLinkedList struct {
	Front *Nodes
	Len   int
}
