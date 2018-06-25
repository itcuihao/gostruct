package gostruct

import "fmt"

// Display view
func (s *SingleLinkedList) Display() {
	list := s.Front
	for list != nil {
		fmt.Printf("%+v ->", list.Value)
		list = list.Next
	}
	fmt.Println()
}

// NewSingleLinkedList new single linkedlist
func NewSingleLinkedList() *SingleLinkedList {
	return &SingleLinkedList{}
}

// Init init
func (s *SingleLinkedList) Init() *SingleLinkedList {
	s.Len = 0
	return s
}

// PreAppend 添加头结点
func (s *SingleLinkedList) PreAppend(i interface{}) {
	n := &Nodes{
		Value: i,
	}
	if s.Front == nil {
		s.Front = n
	} else {
		n.Next = s.Front
		s.Front = n
	}
	s.Len++
}

// Append 添加尾节点
func (s *SingleLinkedList) Append(i interface{}) {
	n := &Nodes{
		Value: i,
	}
	if s.Front == nil {
		s.Front = n
	} else {
		tmp := s.Front
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = n
	}
	s.Len++
}

// IsEmpty is empty
func (s *SingleLinkedList) IsEmpty() bool {

	if s.Len == 0 {
		return true
	}

	return false
}
