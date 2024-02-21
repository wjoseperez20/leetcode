package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	// Base Condition
	if p != nil && q == nil {
		return false
	}

	if p == nil && q != nil {
		return false
	}

	if p == nil && q == nil {
		return true
	}

	// Problem Condition
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	testCases := []struct {
		p      *TreeNode
		q      *TreeNode
		result bool
	}{
		{&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}, true},
		{&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}, &TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}, true},
	}

	for _, testCase := range testCases {
		result := isSameTree(testCase.p, testCase.q)
		if result != testCase.result {
			println("Expected", testCase.result, "but got", result)
		}
	}
}
