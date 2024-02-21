package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}

	return isSymmetricHelper(root.Left, root.Right)
}

func isSymmetricHelper(p *TreeNode, q *TreeNode) bool {
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

	return isSymmetricHelper(p.Right, q.Left) && isSymmetricHelper(p.Left, q.Right)
}

func main() {
	testCases := []struct {
		tree   *TreeNode
		result bool
	}{
		{&TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}, true},
		{&TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}, false},
	}

	for _, testCase := range testCases {
		result := isSymmetric(testCase.tree)
		if result != testCase.result {
			println("Expected", testCase.result, "but got", result)
		}
	}
}
