package linkedlists

// Define ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

// Pattern: dummy-head linked-list construction.
// Invariant: tail always points to the last node in the built list.
// Complexity: O(n) time, O(n) space.
// Interview line: a dummy head avoids special-casing the first node.
func BuildLinkedList(values []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, val := range values {
		tail.Next = &ListNode{Val: val}
		tail = tail.Next
	}
	return dummy.Next
}

// Pattern: linear traversal to array.
// Invariant: values contains every node value already visited in order.
// Complexity: O(n) time, O(n) space.
// Interview line: conversion is handy for tests, but avoid it if the prompt asks O(1) space.
func ToSlice(head *ListNode) []int {
	var res []int
	cur := head
	for cur != nil {
		res = append(res, cur.Val)
		cur = cur.Next
	}
	return res
}

// Pattern: iterative pointer reversal.
// Invariant: prev is the reversed prefix, cur is the next node to move.
// Complexity: O(n) time, O(1) extra space.
// Interview line: save next before rewiring cur.Next.
func ReverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}

// Pattern: fast/slow pointer cycle detection.
// Invariant: fast moves twice as quickly, so it meets slow iff a cycle exists.
// Complexity: O(n) time, O(1) extra space.
// Interview line: Floyd's algorithm detects a cycle without extra memory.
func HasCycle(head *ListNode) bool {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Pattern: Floyd cycle entry detection.
// Invariant: after meeting, moving head and meeting pointer together finds cycle start.
// Complexity: O(n) time, O(1) extra space.
// Interview line: the distance math makes the second phase land at the entry node.
func DetectCycleStart(head *ListNode) *ListNode {
	slow := head
	fast := head
	var meet *ListNode

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			meet = slow
			break
		}
	}

	if meet == nil {
		return nil
	}

	p1 := head
	p2 := meet
	for p1 != p2 {
		p1 = p1.Next
		p2 = p2.Next
	}
	return p1
}

// Pattern: fast/slow pointer midpoint.
// Invariant: slow advances once for every two fast steps.
// Complexity: O(n) time, O(1) extra space.
// Interview line: when fast reaches the end, slow is at the middle.
func MiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// Pattern: deque comparison from both ends.
// Invariant: remaining deque values must still form a palindrome.
// Complexity: O(n) time, O(n) space.
// Interview line: this is the simple version; reverse second half for O(1) space.
func IsPalindrome(head *ListNode) bool {
	var vals []int
	cur := head
	for cur != nil {
		vals = append(vals, cur.Val)
		cur = cur.Next
	}

	left := 0
	right := len(vals) - 1
	for left < right {
		if vals[left] != vals[right] {
			return false
		}
		left++
		right--
	}
	return true
}
