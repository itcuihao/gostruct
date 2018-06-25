package gostruct

import (
	"fmt"
	"testing"
	"time"
)

func TestStack(t *testing.T) {
	gs := NewStack()
	gs.Push(1)
	fmt.Println(gs.Pop())
	gs.Push(time.Now())
	gs.Push(nil)
	fmt.Println(gs.Pop())
	fmt.Println(gs.Pop())
}

// go test -bench=. -run=none
func BenchmarkStack(b *testing.B) {
	gs := NewStack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gs.Push(1)
		gs.Pop()
		gs.Push(2)
		gs.Pop()
	}
}

// go test -bench=. -benchmem -run=none
func BenchmarkStack1(b *testing.B) {
	gs := NewStack()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gs.Push(1)
		gs.Push(2)
		gs.Pop()
		gs.Pop()
	}
}
