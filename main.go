package main

import (
	"./src/queue/"
)

func main() {
	q := queue.NewQueue()
	q.Add(1);
	q.Add(2);
	q.Add(3);
	q.Pop()
	q.PrintAll();
}
