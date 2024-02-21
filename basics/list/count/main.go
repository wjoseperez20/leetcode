package main

type Node struct {
	Value int
	Next  *Node
}

func count(head *Node) int {
	// Base Case
	if head.Next == nil {
		return 1
	}

	// Recursion Case
	return 1 + count(head.Next)

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
	}

	for _, t := range testCases {
		actual := count(t.input)
		if actual != t.expected {
			panic("Test failed")
		}
	}
}
