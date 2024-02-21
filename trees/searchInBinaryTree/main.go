package main

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func searchInBinaryTree(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Value == target {
		return true
	}

	return searchInBinaryTree(root.Left, target) || searchInBinaryTree(root.Right, target)
}

func main() {
	testCases := []struct {
		root   *TreeNode
		target int
		result bool
	}{
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 2, true},
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 3, true},
		{&TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}, 4, false},
	}

	for _, testCase := range testCases {
		result := searchInBinaryTree(testCase.root, testCase.target)
		if result != testCase.result {
			panic("failed")
		}
	}
}
