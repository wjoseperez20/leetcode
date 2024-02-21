package main

import (
	"fmt"
	"math"
)

// Dijkstra function finds the shortest path in a graph using Dijkstra's algorithm
func dijkstra(graph map[string]map[string]float64, start, end string) float64 {
	costs := make(map[string]float64)
	parents := make(map[string]string)

	for node := range graph {
		costs[node] = math.Inf(1)
		parents[node] = ""
	}

	costs[start] = 0
	processed := make(map[string]bool)

	for {
		node := findLowestCostNode(costs, processed)
		if node == "" {
			break
		}

		cost := costs[node]
		neighbors := graph[node]

		for n, w := range neighbors {
			newCost := cost + w
			if newCost < costs[n] {
				costs[n] = newCost
				parents[n] = node
			}
		}

		processed[node] = true
	}

	return costs[end]
}

// findLowestCostNode finds the node with the lowest cost
func findLowestCostNode(costs map[string]float64, processed map[string]bool) string {
	lowestCost := math.Inf(1)
	var lowestCostNode string

	for node, cost := range costs {
		if cost < lowestCost && !processed[node] {
			lowestCost = cost
			lowestCostNode = node
		}
	}

	return lowestCostNode
}

func main() {
	testCases := map[string]struct {
		graph    map[string]map[string]float64
		expected float64
	}{
		"case1": {
			graph: map[string]map[string]float64{
				"start": {"a": 6, "b": 2},
				"a":     {"fin": 1},
				"b":     {"a": 3, "fin": 5},
				"fin":   {},
			},
			expected: 6,
		},
		"case2": {
			graph: map[string]map[string]float64{
				"start": {"a": 2},
				"a":     {"fin": 2},
				"fin":   {},
			},
			expected: 4,
		},
		"case3": {
			graph: map[string]map[string]float64{
				"start": {"a": 5, "b": 2},
				"a":     {"c": 4, "d": 2},
				"b":     {"a": 8, "d": 7},
				"c":     {"fin": 3, "d": 6},
				"d":     {"fin": 1},
				"fin":   {},
			},
			expected: 8,
		},
		"case4": {
			graph: map[string]map[string]float64{
				"start": {"a": 10},
				"a":     {"b": 20},
				"b":     {"c": 1, "fin": 30},
				"c":     {"a": 1},
				"fin":   {},
			},
			expected: 60,
		},
	}

	for key, value := range testCases {
		result := dijkstra(value.graph, "start", "fin")
		if result != value.expected {
			fmt.Printf("Test case %s failed: expected %f, got %f\n", key, value.expected, result)
		} else {
			fmt.Printf("Test case %s passed\n", key)
		}
	}
}
