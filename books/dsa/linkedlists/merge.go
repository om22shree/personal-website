package linkedlists

import "container/heap"

// Pattern: two sorted lists with dummy tail.
// Invariant: tail points to the end of the merged sorted prefix.
// Complexity: O(m + n) time, O(1) extra space.
// Interview line: repeatedly attach the smaller head and advance that list.
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}

	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return dummy.Next
}

// Heap for ListNode elements
type ListNodeHeap []*ListNode

func (h ListNodeHeap) Len() int           { return len(h) }
func (h ListNodeHeap) Less(i, j int) bool { return h[i].Val < h[j].Val }
func (h ListNodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ListNodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}
func (h *ListNodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: k-way merge with min-heap.
// Invariant: heap stores the current smallest candidate from each list.
// Complexity: O(n log k) time, O(k) space.
// Interview line: use index as a tie-breaker so equal values do not compare nodes (not required in Go because Less just compares .Val, and heaps are stable/unstable doesn't crash on ListNode pointers).
func MergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	for _, list := range lists {
		if list != nil {
			heap.Push(h, list)
		}
	}

	dummy := &ListNode{}
	tail := dummy

	for h.Len() > 0 {
		node := heap.Pop(h).(*ListNode)
		tail.Next = node
		tail = tail.Next

		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}

	return dummy.Next
}

// Pattern: two pointers with an n-node gap.
// Invariant: fast is n nodes ahead of slow before both move together.
// Complexity: O(n) time, O(1) extra space.
// Interview line: a dummy node makes deleting the head node clean.
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow := dummy
	fast := dummy

	for i := 0; i <= n; i++ {
		if fast == nil {
			return head
		}
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next
	return dummy.Next
}

// Pattern: split, reverse second half, then weave.
// Invariant: first and reversed second halves are alternately connected.
// Complexity: O(n) time, O(1) extra space.
// Interview line: find middle, reverse the back half, then merge one node at a time.
func ReorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	// Step 1: Find middle
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Step 2: Reverse second half
	var prev *ListNode
	cur := slow.Next
	slow.Next = nil // Split the list

	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	// Step 3: Weave the lists
	first := head
	second := prev

	for second != nil {
		temp1 := first.Next
		temp2 := second.Next

		first.Next = second
		second.Next = temp1

		first = temp1
		second = temp2
	}
}
