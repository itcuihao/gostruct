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

// go test -test.bench=".*"
func BenchmarkQueue(b *testing.B) {
	b.ResetTimer()
	q := NewQueue(1)
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}
func BenchmarkQueue1(b *testing.B) {
	b.ResetTimer()
	q := NewSimpleQueue(1)
	for i := 0; i < b.N; i++ {
		q.Push(i)
		q.Pop()
	}
}
