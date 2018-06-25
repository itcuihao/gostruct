package gostruct

import (
	"fmt"
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	gq := NewQueue(1)
	gq.Push(1)
	fmt.Println(gq.Pop())
	gq.Push(time.Now())
	gq.Push(2)
	fmt.Println(gq.Pop())
	fmt.Println(gq.Pop())
}
