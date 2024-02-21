package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func mergeTwoLists(list1 *Node, list2 *Node) *Node {
	merge := &Node{}
	current := merge

	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 == nil {
		current.Next = list2
	} else {
		current.Next = list1
	}

	return merge.Next
}

func displayList(head *Node) []int {
	var result []int
	current := head

	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}

	return result
}

func main() {
	testCases := []struct {
		l1 Node
		l2 Node
	}{
		{Node{1, &Node{2, &Node{4, nil}}}, Node{1, &Node{3, &Node{4, nil}}}},
		{Node{}, Node{}},
		{Node{}, Node{0, nil}},
	}

	for _, value := range testCases {
		fmt.Printf("Merge List = %v\n", displayList(mergeTwoLists(&value.l1, &value.l2)))
	}

}
