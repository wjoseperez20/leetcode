package main

import (
	"container/list"
	"fmt"
	"github.com/golang-collections/collections/queue"
	"reflect"
	"time"
)

// Global Graph
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

// List
func searchL(name string) bool {
	searchQueue := list.New()
	searchQueue.PushBack(graph[name])

	searched := make(map[string]bool)

	for searchQueue.Len() > 0 {
		personElement := searchQueue.Front()
		searchQueue.Remove(personElement)
		persons := personElement.Value.([]string)

		for _, person := range persons {

			if _, ok := searched[person]; !ok {

				if personIsSeller(person) {
					return true
				} else {
					searchQueue.PushBack(graph[person])
					searched[person] = true
				}
			}
		}
	}

	return false
}

// Queue
func searchQ(name string) (bool, int, []string) {
	searchQueue := queue.New()
	searchQueue.Enqueue([]string{name})
	visited := make(map[string]bool)

	// Start the breadth-first search
	for searchQueue.Len() > 0 {

		// Dequeue the first element from the queue
		currentPath := searchQueue.Dequeue().([]string)
		currentPerson := currentPath[len(currentPath)-1]

		if _, ok := visited[currentPerson]; !ok {
			visited[currentPerson] = true

			if personIsSeller(currentPerson) {
				return true, len(currentPath) - 1, currentPath
			}

			// If the current person is not the mango seller,
			// enqueue its neighbors to explore further
			for _, neighbor := range graph[currentPerson] {
				newPath := append(currentPath, neighbor)
				searchQueue.Enqueue(newPath)
			}
		}
	}

	return false, 0, nil
}

func main() {
	testCases := []struct {
		name         string
		expected     bool
		path         []string
		shortestPath int
	}{
		{"you", true, []string{"you", "claire", "thom"}, 2},
		{"bob", false, []string{}, 2},
		{"alice", false, []string{}, 1},
		{"claire", true, []string{"claire", "thom"}, 1},
		{"thom", true, []string{"thom"}, 0},
	}

	start := time.Now()
	for _, testCase := range testCases {
		searchL(testCase.name)
	}
	elapsed := time.Since(start)
	fmt.Println("List Search took", elapsed)

	start = time.Now()
	for _, testCase := range testCases {

		result, short, path := searchQ(testCase.name)
		if result != testCase.expected && short != testCase.shortestPath && reflect.DeepEqual(path, testCase.path) {
			fmt.Printf("Test failed for %v, expected %v, got %v\n", testCase.name, testCase.expected, result)
		}
	}
	elapsed = time.Since(start)
	fmt.Println("Queue Search took", elapsed)
}
