package main

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func orderTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	result = append(result, root.Value)
	result = append(result, orderTree(root.Left)...)
	result = append(result, orderTree(root.Right)...)

	// BubbleSort
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}

	return result
}

func main() {
	testCases := []struct {
		root   *TreeNode
		result []int
	}{
		{&TreeNode{1, &TreeNode{3, nil, nil}, &TreeNode{2, nil, nil}}, []int{1, 2, 3}},
	}

	for _, testCase := range testCases {
		result := orderTree(testCase.root)
		if !reflect.DeepEqual(result, testCase.result) {
			fmt.Printf("Expected %v, but got %v\n", testCase.result, result)
		}

	}
}
