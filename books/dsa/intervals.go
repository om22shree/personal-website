package dsa

import (
	"container/heap"
	"sort"
)

// Define Interval struct
type Interval struct {
	Start int
	End   int
}

// Pattern: sort by start and merge overlapping intervals.
// Invariant: merged holds non-overlapping intervals covering everything processed.
// Complexity: O(n log n) time, O(n) space for output.
// Interview line: after sorting, only the last merged interval can overlap the next one.
func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	var merged []Interval
	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		lastIdx := len(merged) - 1
		if intervals[i].Start <= merged[lastIdx].End {
			if intervals[i].End > merged[lastIdx].End {
				merged[lastIdx].End = intervals[i].End
			}
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}

// Pattern: three-phase interval insertion.
// Invariant: append intervals before, merge overlaps, then append intervals after.
// Complexity: O(n) time, O(n) space for output.
// Interview line: because input is sorted, we only merge the contiguous overlap block.
func InsertInterval(intervals []Interval, newInterval Interval) []Interval {
	var res []Interval
	i := 0
	n := len(intervals)

	// Phase 1: Append all intervals that end before the new one starts
	for i < n && intervals[i].End < newInterval.Start {
		res = append(res, intervals[i])
		i++
	}

	// Phase 2: Merge overlapping intervals
	for i < n && intervals[i].Start <= newInterval.End {
		if intervals[i].Start < newInterval.Start {
			newInterval.Start = intervals[i].Start
		}
		if intervals[i].End > newInterval.End {
			newInterval.End = intervals[i].End
		}
		i++
	}
	res = append(res, newInterval)

	// Phase 3: Append all intervals that start after the new one ends
	for i < n {
		res = append(res, intervals[i])
		i++
	}

	return res
}

// Pattern: greedy by earliest end time.
// Invariant: keep the interval that leaves the most room for future intervals.
// Complexity: O(n log n) time, O(1) extra space after sorting.
// Interview line: when two intervals overlap, dropping the one with later end is optimal.
func EraseOverlapIntervals(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].End < intervals[j].End
	})

	count := 0
	lastEnd := intervals[0].End

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < lastEnd {
			count++ // drop this one
		} else {
			lastEnd = intervals[i].End
		}
	}
	return count
}

// Pattern: sorted adjacent overlap check.
// Invariant: after sorting, any conflict must be between neighboring intervals.
// Complexity: O(n log n) time, O(1) extra space after sorting.
// Interview line: if every meeting starts after the previous ends, one person can attend all.
func CanAttendMeetings(intervals []Interval) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			return false
		}
	}
	return true
}

// Heap of integers for meeting room end times
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: sweep line with min-heap of meeting end times.
// Invariant: heap contains end times for rooms currently in use.
// Complexity: O(n log n) time, O(n) space.
// Interview line: free the earliest-ending room before allocating a new one.
func MinMeetingRooms(intervals []Interval) int {
	if len(intervals) == 0 {
		return 0
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	rooms := &IntHeap{}
	heap.Init(rooms)

	heap.Push(rooms, intervals[0].End)

	for i := 1; i < len(intervals); i++ {
		// If earliest room is free, reuse it
		if intervals[i].Start >= (*rooms)[0] {
			heap.Pop(rooms)
		}
		heap.Push(rooms, intervals[i].End)
	}

	return rooms.Len()
}

// Pattern: two pointers over two sorted interval lists.
// Invariant: advance the interval that ends first because it cannot overlap future intervals.
// Complexity: O(m + n) time, O(1) extra space excluding output.
// Interview line: intersection is max(starts) to min(ends), if that range is valid.
func IntervalIntersection(listA []Interval, listB []Interval) []Interval {
	var res []Interval
	i := 0
	j := 0

	for i < len(listA) && j < len(listB) {
		start := listA[i].Start
		if listB[j].Start > start {
			start = listB[j].Start
		}

		end := listA[i].End
		if listB[j].End < end {
			end = listB[j].End
		}

		if start <= end {
			res = append(res, Interval{Start: start, End: end})
		}

		if listA[i].End < listB[j].End {
			i++
		} else {
			j++
		}
	}
	return res
}
