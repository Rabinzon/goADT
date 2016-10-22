package stack

import (
	"errors"
	"fmt"
)

type Node struct {
	value interface{}
}

func newNode(value interface{}) *Node {
	return &Node{value}
}

type Stack struct {
	nodes []*Node
	top   *Node
	count int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(value interface{}) {
	node := newNode(value)
	s.nodes = append(s.nodes, node)
	s.top = node
	s.count++
}

func (s *Stack) Empty() bool {
	return s.count == 0
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return 0, errors.New("Empty Stack")
	}

	top := s.top
	s.nodes = s.nodes[:s.count]
	s.count--
	s.top = s.nodes[s.count]

	return top.value, nil
}

func (s *Stack) PrintStack() {
	i := s.count

	for i != 0 {
		fmt.Println(s.nodes[i-1].value)
		i--
	}
}
