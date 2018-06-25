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
