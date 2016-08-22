package queue

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
}

type Queue struct {
	nodes []*Node
	head  *Node
	tail  *Node
	count int
}

func NewQueue() *Queue {
	return &Queue{}
}

func newNode(value int) *Node {
	return &Node{value}
}

func (q *Queue) Pop() (int, error) {
	head := q.head

	if head == nil {
		return 0, errors.New("Empty Stack")
	}

	q.nodes = q.nodes[1:]

	if head == q.tail {
		q.head = nil
		q.count--
		return head.value, nil
	}

	q.head = q.nodes[0]
	q.count--

	return head.value, nil
}

func (q *Queue) Add(value int) {
	node := newNode(value)

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
