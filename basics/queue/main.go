package main

import "fmt"

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type queue struct {
	items []interface{}
}

// QueueOperations is an interface for queue operations.
type QueueOperations interface {
	Push(item interface{})
	Pop() interface{}
	Len() int
}

func (q *queue) Push(item interface{}) {
	q.items = append(q.items, item)
}

func (q *queue) Pop() interface{} {
	item := q.items[0]
	q.items = q.items[1:len(q.items)]
	return item
}

func (q *queue) Len() int {
	return len(q.items)
}

func main() {
	// Create a new queue and put some items in it
	q := &queue{}

	// Create an interface variable and assign the queue to it
	var qo QueueOperations
	qo = q

	// Add some items to the queue using the interface
	qo.Push("foo")
	qo.Push("bar")
	qo.Push("baz")

	// Pop and print items from the queue using the interface
	for qo.Len() > 0 {
		fmt.Println(qo.Pop())
	}
}
