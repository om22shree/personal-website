package heaps

import (
	"container/heap"
)

// Array element tracker
type ArrayElement struct {
	val      int
	arrIdx   int
	elemIdx  int
}

type ArrayElementHeap []ArrayElement

func (h ArrayElementHeap) Len() int           { return len(h) }
func (h ArrayElementHeap) Less(i, j int) bool { return h[i].val < h[j].val }
func (h ArrayElementHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ArrayElementHeap) Push(x interface{}) {
	*h = append(*h, x.(ArrayElement))
}
func (h *ArrayElementHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: k-way merge with min-heap.
// Invariant: heap stores the next unmerged value from each array.
// Complexity: O(n log k) time, O(k) space.
// Interview line: always emit the smallest current head, then advance that source.
func MergeKSortedArrays(arrays [][]int) []int {
	h := &ArrayElementHeap{}
	heap.Init(h)

	for i, arr := range arrays {
		if len(arr) > 0 {
			heap.Push(h, ArrayElement{val: arr[0], arrIdx: i, elemIdx: 0})
		}
	}

	var res []int
	for h.Len() > 0 {
		curr := heap.Pop(h).(ArrayElement)
		res = append(res, curr.val)

		nextElemIdx := curr.elemIdx + 1
		if nextElemIdx < len(arrays[curr.arrIdx]) {
			heap.Push(h, ArrayElement{
				val:     arrays[curr.arrIdx][nextElemIdx],
				arrIdx:  curr.arrIdx,
				elemIdx: nextElemIdx,
			})
		}
	}
	return res
}

// Local ListNode definition to avoid cross-package imports
type ListNode struct {
	Val  int
	Next *ListNode
}

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

// Pattern: k-way merge with min-heap and tie-breaker index.
// Invariant: heap stores the current node from each linked list.
// Complexity: O(n log k) time, O(k) space.
// Interview line: add the list index so equal node values do not compare ListNode objects.
func MergeKLists(lists []*ListNode) *ListNode {
	h := &ListNodeHeap{}
	heap.Init(h)

	for _, l := range lists {
		if l != nil {
			heap.Push(h, l)
		}
	}

	dummy := &ListNode{}
	tail := dummy

	for h.Len() > 0 {
		curr := heap.Pop(h).(*ListNode)
		tail.Next = curr
		tail = tail.Next

		if curr.Next != nil {
			heap.Push(h, curr.Next)
		}
	}
	return dummy.Next
}

// Log event
type LogEvent struct {
	Timestamp int64
	Message   string
	StreamIdx int
	EventIdx  int
}

type LogEventHeap []LogEvent

func (h LogEventHeap) Len() int           { return len(h) }
func (h LogEventHeap) Less(i, j int) bool { return h[i].Timestamp < h[j].Timestamp }
func (h LogEventHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *LogEventHeap) Push(x interface{}) {
	*h = append(*h, x.(LogEvent))
}
func (h *LogEventHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: k-way merge for sorted timestamp streams.
// Invariant: heap stores the next log event from each stream.
// Complexity: O(n log k) time, O(k) space.
// Interview line: this is merge-k-lists framed as ordered log aggregation.
func MergeSortedLogStreams(streams [][]LogEvent) []LogEvent {
	h := &LogEventHeap{}
	heap.Init(h)

	for i, stream := range streams {
		if len(stream) > 0 {
			heap.Push(h, LogEvent{
				Timestamp: stream[0].Timestamp,
				Message:   stream[0].Message,
				StreamIdx: i,
				EventIdx:  0,
			})
		}
	}

	var res []LogEvent
	for h.Len() > 0 {
		curr := heap.Pop(h).(LogEvent)
		res = append(res, curr)

		nextIdx := curr.EventIdx + 1
		if nextIdx < len(streams[curr.StreamIdx]) {
			heap.Push(h, LogEvent{
				Timestamp: streams[curr.StreamIdx][nextIdx].Timestamp,
				Message:   streams[curr.StreamIdx][nextIdx].Message,
				StreamIdx: curr.StreamIdx,
				EventIdx:  nextIdx,
			})
		}
	}
	return res
}
