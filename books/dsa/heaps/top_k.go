package heaps

import (
	"container/heap"
)

// KthLargest struct
type KthLargest struct {
	k    int
	heap *IntMinHeap
}

// Define IntMinHeap for KthLargest
type IntMinHeap []int

func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func NewKthLargest(k int, nums []int) KthLargest {
	h := &IntMinHeap{}
	heap.Init(h)
	kl := KthLargest{k: k, heap: h}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.heap, val)
	if this.heap.Len() > this.k {
		heap.Pop(this.heap)
	}
	return (*this.heap)[0]
}

// Frequency Pair
type FreqPair struct {
	val  int
	freq int
}

type FreqMinHeap []FreqPair

func (h FreqMinHeap) Len() int           { return len(h) }
func (h FreqMinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h FreqMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *FreqMinHeap) Push(x interface{}) {
	*h = append(*h, x.(FreqPair))
}
func (h *FreqMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: frequency count + top-k extraction.
// Invariant: Counter maps each value to its frequency.
// Complexity: O(n log k) conceptually, O(n) space.
// Interview line: count first, then select by frequency rather than by value.
func TopKFrequent(nums []int, k int) []int {
	counts := make(map[int]int)
	for _, num := range nums {
		counts[num]++
	}

	h := &FreqMinHeap{}
	heap.Init(h)

	for val, freq := range counts {
		heap.Push(h, FreqPair{val: val, freq: freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(h).(FreqPair).val
	}
	return res
}

// Point with squared distance
type Point struct {
	x    int
	y    int
	dist int
}

type PointMaxHeap []Point

func (h PointMaxHeap) Len() int           { return len(h) }
func (h PointMaxHeap) Less(i, j int) bool { return h[i].dist > h[j].dist } // Max-heap
func (h PointMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *PointMaxHeap) Push(x interface{}) {
	*h = append(*h, x.(Point))
}
func (h *PointMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: heap selection with custom distance key.
// Invariant: squared distance preserves ordering without sqrt.
// Complexity: O(n log k) time, O(k) space.
// Interview line: compare squared distances to avoid unnecessary floating-point work.
func KClosest(points [][]int, k int) [][]int {
	h := &PointMaxHeap{}
	heap.Init(h)

	for _, pt := range points {
		dist := pt[0]*pt[0] + pt[1]*pt[1]
		heap.Push(h, Point{x: pt[0], y: pt[1], dist: dist})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	res := make([][]int, k)
	for i := 0; i < k; i++ {
		p := heap.Pop(h).(Point)
		res[i] = []int{p.x, p.y}
	}
	return res
}
