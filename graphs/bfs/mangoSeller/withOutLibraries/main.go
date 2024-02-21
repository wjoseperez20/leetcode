package main

import "fmt"

// Queue represents a queue data structure.
type Queue struct {
	items []interface{}
}

// Enqueue adds an item to the end of the queue.
func (q *Queue) Enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// Dequeue removes and returns the item at the front of the queue.
func (q *Queue) Dequeue() interface{} {
	if len(q.items) == 0 {
		return nil
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

// Len returns the number of items in the queue.
func (q *Queue) Len() int {
	return len(q.items)
}

var graph = map[string][]string{
	"you":    {"alice", "bob", "claire"},
	"bob":    {"anuj", "peggy"},
	"alice":  {"peggy"},
	"claire": {"thom", "jonny"},
	"anuj":   {},
	"peggy":  {},
	"thom":   {},
	"jonny":  {},
}

func personIsSeller(person string) bool {
	return person == "thom" // Assuming "thom" is the mango seller in this example
}

func search(name string) bool {
	searchQueue := Queue{}

	// Push the starting node onto the queue
	searchQueue.Enqueue(graph[name])

	searched := make(map[string]bool)

	for searchQueue.Len() > 0 {
		// Pop the front of the queue
		personElement := searchQueue.Dequeue().([]string)
		persons := personElement

		for _, person := range persons {
			if _, ok := searched[person]; !ok {
				if personIsSeller(person) {
					fmt.Println(person, "is a mango seller!")
					return true
				} else {
					// Push the neighbors of the current person onto the queue
					searchQueue.Enqueue(graph[person])
					searched[person] = true
				}
			}
		}
	}
	return false
}

func main() {
	search("you")
}
