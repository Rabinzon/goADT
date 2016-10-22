package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack()

	if stack.count != 0 {
		t.Error("should creat new Stack")
	}
}

func TestStack_Push(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(24)
	stack.Push(3)
	if stack.count != 3 {
		t.Error("Stack count should be equals 3")
	}
}

func TestStack_Pop(t *testing.T) {
	stack := NewStack()
	test_stack := NewStack()

	stack.Push(1)
	stack.Push(test_stack)
	stack.Push(4)

	value, error := stack.Pop()

	if error != nil {
		t.Error("shouldn't get Error")
	}

	if value != 4 {
		t.Error("value should be equals 4")
	}
}

func TestStack_Empty(t *testing.T) {
	stack := NewStack()
	stack.Push(1)

	if stack.Empty() {
		t.Error("Stack shouldn't be empty")
	}

	stack.Pop()

	if !stack.Empty() {
		t.Error("Stack should be empty")
	}
}
