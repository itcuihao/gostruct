package gostruct

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	gs := NewSingleLinkedList()
	fmt.Println(gs.IsEmpty())
	gs.PreAppend(1)
	gs.PreAppend(2)
	gs.Display()

	gs.Append(3)
	gs.Display()
	gs.Init()
	fmt.Println(gs.IsEmpty())

}
