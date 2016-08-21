package queue

import "fmt"

type Node struct {
	value int
}

type Queue struct {
	nodes []*Node
	head  *Node
	tail  *Node
	count int
}

func (q *Queue) Pop() *Node {
	head := q.head

	if head == nil {
		return nil
	}

	q.nodes = q.nodes[1:]

	if head == q.tail {
		q.head = nil
		q.count--
		return head
	}
	q.head = q.nodes[0]
	q.count--

	return head
}

func (q *Queue) Add(value int) {
	node := NewNode(value)

	if q.count == 0 {
		q.head = node
	}

	q.tail = node
	q.nodes = append(q.nodes, node)
	q.count++
}

func (q *Queue) PrintAll() {
	if q.count == 0 {
		fmt.Println("Queue is empty.")
		return
	}

	i := 0

	for i != q.count {
		fmt.Println(q.nodes[i].value)
		i++
	}

}

func NewQueue() *Queue {
	return &Queue{}
}

func NewNode(value int) *Node {
	return &Node{value}
}
