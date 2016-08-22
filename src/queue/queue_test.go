package queue

import "testing"

func TestNewQueue(t *testing.T) {
	q := NewQueue()

	if q.count != 0 {
		t.Error("should creat empty Queue")
	}
}

func TestQueue_Add(t *testing.T) {
	q := NewQueue()

	q.Add(1)

	if q.count != 1 {
		t.Error("Queue count should should be equals 1")
	}

	q.Add(2)

	if q.count != 2 {
		t.Error("Queue count should should be equals 2")
	}
}

func TestQueue_Pop(t *testing.T) {
	q := NewQueue()

	q.Add(1)
	q.Add(24)
	q.Add(33)

	v, error := q.Pop()

	if error != nil {
		t.Error("shouldn't get Error")
	}

	if v != 1 {
		t.Error("v should be equals 1")
	}

	q.Pop()

	v, error = q.Pop()

	if error != nil {
		t.Error("shouldn't get Error")
	}

	if v != 33 {
		t.Error("v should be equals 33")
	}
}
