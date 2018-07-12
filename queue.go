package gostruct

// Queue queue
type Queue struct {
	Nodes []*Node
	Size  int
	Count int
	Head  int
	Tail  int
}

// Push push
func (q *Queue) Push(v interface{}) {
	if q.Count == q.Size {
		return
	}

	n := NewNode(v)

	if q.Head == q.Tail && q.Count > 0 {
		nodes := make([]*Node, 0, len(q.Nodes)+q.Size)
		nodes = append(nodes, q.Nodes...)
		nodes = append(nodes, n)

		q.Head = 0
		q.Tail = len(q.Nodes)
		q.Nodes = nodes
	}
	q.Nodes[q.Tail] = n
	q.Tail = (q.Tail + 1) % len(q.Nodes)
	q.Count++
}

// Pop pop
func (q *Queue) Pop() *Node {
	if q.Count == 0 {
		return nil
	}

	node := q.Nodes[q.Head]
	q.Head = (q.Head + 1) % len(q.Nodes)
	q.Count--

	return node
}

// NewQueue new queue
func NewQueue(size int) *Queue {
	return &Queue{
		Nodes: make([]*Node, size),
		Size:  size,
	}
}

// Queue1 queue
type Queue1 struct {
	Nodes []*Node
	Size  int
}

// Push push
func (q *Queue1) Push(v interface{}) {

	if q.Len() == q.Size {
		return
	}

	n := NewNode(v)
	nodes := make([]*Node, 0, len(q.Nodes)+1)
	nodes = append(nodes, q.Nodes...)
	nodes = append(nodes, n)

	q.Nodes = nodes

}

// Pop pop
func (q *Queue1) Pop() *Node {
	if q.Len() == 0 {
		return nil
	}

	pnode := q.Nodes[0]
	node := q.Nodes[1:]
	q.Nodes = node

	return pnode
}

// Len 队列长度
func (q *Queue1) Len() int {
	return len(q.Nodes)
}

// IsEmpty 是否为空
func (q *Queue1) IsEmpty() bool {
	return len(q.Nodes) == 0
}

// NewSimpleQueue new queue
func NewSimpleQueue(size int) *Queue1 {
	return &Queue1{
		Nodes: make([]*Node, 0, size),
		Size:  size,
	}
}
