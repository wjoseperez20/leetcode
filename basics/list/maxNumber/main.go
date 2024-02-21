package main

type Node struct {
	Value int
	Next  *Node
}

func maxNumber(head *Node) int {
	// Base case
	if head.Next == nil {
		return head.Value
	}

	// Recursive case
	return max(head.Value, maxNumber(head.Next))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func main() {
	testCases := []struct {
		input    *Node
		expected int
	}{
		{input: &Node{Value: 1, Next: nil}, expected: 1},
		{input: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5}}}}}, expected: 5},
		{input: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6}}}}}}, expected: 6},
		{input: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6, Next: &Node{Value: 7}}}}}}}, expected: 7},
		{input: &Node{Value: 1, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6, Next: &Node{Value: 7, Next: &Node{Value: 8}}}}}}}}, expected: 8},
		{input: &Node{Value: 20, Next: &Node{Value: 2, Next: &Node{Value: 3, Next: &Node{Value: 4, Next: &Node{Value: 5, Next: &Node{Value: 6, Next: &Node{Value: 7, Next: &Node{Value: 8, Next: &Node{Value: 9}}}}}}}}}, expected: 20},
	}

	for _, t := range testCases {
		actual := maxNumber(t.input)
		if actual != t.expected {
			panic("Test failed")
		}
	}
}
