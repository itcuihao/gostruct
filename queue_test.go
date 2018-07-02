package gostruct

import (
	"fmt"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	gq := NewSimpleQueue(1)
	gq.Push(1)
	fmt.Println("len:", gq.Len())
	fmt.Println(gq.Pop())
	fmt.Println(gq.IsEmpty())
	gq.Push(time.Now())
	fmt.Println(gq.IsEmpty())
	fmt.Println(gq.Pop())
	fmt.Println(gq.IsEmpty())
	gq.Push(2)
	fmt.Println(gq.Pop())
}
