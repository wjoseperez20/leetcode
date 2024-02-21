package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicatesDisorder(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	current := head

	for current != nil {
		runner := current

		for runner.Next != nil {
			if current.Val == runner.Next.Val {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}

		current = current.Next
	}

	return head
}

func deleteDuplicates_(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	current := head

	for {
		if current == nil || current.Next == nil {
			break
		}

		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head
	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			current.Next = current.Next.Next
		} else {
			current = current.Next
		}
	}

	return head
}

func main() {
	testCasesOrder := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{&ListNode{1, &ListNode{1, &ListNode{2, nil}}}, &ListNode{1, &ListNode{2, nil}}},
		{&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}}}, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
		{&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, nil}}}}}, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}},
	}

	for _, testCase := range testCasesOrder {
		actual := deleteDuplicates(testCase.input)

		if actual.Val != testCase.expected.Val {
			fmt.Printf("Input: %v, expected: %v, actual: %v\n", testCase.input, testCase.expected, actual)
		}
	}

	testCasesDisorder := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{&ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{1, nil}}}}}}, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}},
		{&ListNode{1, &ListNode{1, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}}}, &ListNode{1, &ListNode{3, &ListNode{4, nil}}}},
	}

	for _, testCase := range testCasesDisorder {
		actual := deleteDuplicatesDisorder(testCase.input)

		if actual.Val != testCase.expected.Val {
			fmt.Printf("Input: %v, expected: %v, actual: %v\n", testCase.input, testCase.expected, actual)
		}
	}
}
