# gostruct

## Stack

> go test -run TestStack

```Go
func TestStack(t *testing.T) {
	gs := &gostruct.Stack{}
	gs.Push(1)
	fmt.Println(gs.Pop())
	gs.Push(time.Now())
	gs.Push(nil)
	fmt.Println(gs.Pop())
	fmt.Println(gs.Pop())
}
```

