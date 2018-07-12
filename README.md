# gostruct

常见的数据结构，栈，队列，树，图等等

## Stack

> go test -run TestStack

```Go
func TestStack(t *testing.T) {
    gs := gostruct.NewStack()
	gs.Push(1)
	fmt.Println(gs.Pop())
	gs.Push(time.Now())
	gs.Push(nil)
	fmt.Println(gs.Pop())
	fmt.Println(gs.Pop())
}

output:
1
<nil>
2018-06-25 18:05:22.931157547 +0800 CST m=+0.000417280
PASS
ok      tests/strings   0.002s
```

## Queue 单链队列

> go test -run TestQueue

```Go
func TestStack(t *testing.T) {
	gs := gostruct.NewQueue(1)
	gs.Push(1)
	fmt.Println(gs.Pop())
	gs.Push(time.Now())
	gs.Push(2)
	fmt.Println(gs.Pop())
    fmt.Println(gs.Pop())
}

output:
1
2018-06-25 18:02:58.915803906 +0800 CST m=+0.000442400
<nil>
PASS
ok      tests/strings   0.002s
```