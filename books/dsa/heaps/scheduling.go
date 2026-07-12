package heaps

import (
	"container/heap"
)

// IntMaxHeap is a max-heap of integers
type IntMaxHeap []int

func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] } // Max-heap
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: max-heap simulation plus cooldown batches.
// Invariant: each cycle schedules up to n + 1 most frequent remaining tasks.
// Complexity: O(t log u) time, O(u) space where u is unique tasks.
// Interview line: greedily run the most frequent tasks first to minimize idle slots.
func LeastInterval(tasks []byte, n int) int {
	countsMap := make(map[byte]int)
	for _, task := range tasks {
		countsMap[task]++
	}

	h := &IntMaxHeap{}
	heap.Init(h)
	for _, count := range countsMap {
		heap.Push(h, count)
	}

	time := 0

	for h.Len() > 0 {
		var cooldown []int
		slots := n + 1

		for slots > 0 && h.Len() > 0 {
			count := heap.Pop(h).(int) - 1
			if count > 0 {
				cooldown = append(cooldown, count)
			}
			time++
			slots--
		}

		for _, count := range cooldown {
			heap.Push(h, count)
		}

		if h.Len() > 0 {
			time += slots
		}
	}

	return time
}

// MedianFinder using min-heap and max-heap
type MedianFinder struct {
	small *IntMaxHeap // lower half, max-heap
	large *IntMinHeap // upper half, min-heap
}

func NewMedianFinder() MedianFinder {
	small := &IntMaxHeap{}
	large := &IntMinHeap{}
	heap.Init(small)
	heap.Init(large)
	return MedianFinder{
		small: small,
		large: large,
	}
}

// Pattern: push to max side, move one to min side, then rebalance sizes.
// Invariant: len(small) is either equal to len(large) or one larger.
// Complexity: O(log n) time, O(n) total space.
// Interview line: balancing heaps keeps the median at the heap tops.
func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.small, num)
	heap.Push(this.large, heap.Pop(this.small).(int))

	if this.large.Len() > this.small.Len() {
		heap.Push(this.small, heap.Pop(this.large).(int))
	}
}

// Pattern: read median from heap tops.
// Invariant: heap sizes determine whether median is one top or average of two tops.
// Complexity: O(1) time, O(1) extra space.
// Interview line: after every insert, the median is immediately available.
func (this *MedianFinder) FindMedian() float64 {
	if this.small.Len() > this.large.Len() {
		return float64((*this.small)[0])
	}
	return float64((*this.small)[0]+(*this.large)[0]) / 2.0
}
