from collections import Counter
import heapq
from typing import List, Optional, Tuple


class ListNode:
    def __init__(self, val: int = 0, next: Optional["ListNode"] = None):
        self.val = val
        self.next = next


# Pattern: heap selection for kth largest.
# Invariant: nlargest keeps the k best values according to heap ordering.
# Complexity: O(n log k) time, O(k) space.
# Interview line: use a heap when you need the best k without fully sorting.
def kth_largest(nums: List[int], k: int) -> int:
    return heapq.nlargest(k, nums)[-1]


# Pattern: frequency count + top-k extraction.
# Invariant: Counter maps each value to its frequency.
# Complexity: O(n log k) conceptually, O(n) space.
# Interview line: count first, then select by frequency rather than by value.
def top_k_frequent(nums: List[int], k: int) -> List[int]:
    return [num for num, _ in Counter(nums).most_common(k)]


# Pattern: heap selection with custom distance key.
# Invariant: squared distance preserves ordering without sqrt.
# Complexity: O(n log k) time, O(k) space.
# Interview line: compare squared distances to avoid unnecessary floating-point work.
def k_closest_points(points: List[List[int]], k: int) -> List[List[int]]:
    return heapq.nsmallest(k, points, key=lambda p: p[0] * p[0] + p[1] * p[1])


# Pattern: k-way merge with min-heap.
# Invariant: heap stores the next unmerged value from each array.
# Complexity: O(n log k) time, O(k) space.
# Interview line: always emit the smallest current head, then advance that source.
def merge_k_sorted_arrays(arrays: List[List[int]]) -> List[int]:
    heap = []
    for array_id, arr in enumerate(arrays):
        if arr:
            heapq.heappush(heap, (arr[0], array_id, 0))

    merged = []
    while heap:
        value, array_id, idx = heapq.heappop(heap)
        merged.append(value)
        next_idx = idx + 1
        if next_idx < len(arrays[array_id]):
            heapq.heappush(heap, (arrays[array_id][next_idx], array_id, next_idx))

    return merged


# Pattern: k-way merge with min-heap and tie-breaker index.
# Invariant: heap stores the current node from each linked list.
# Complexity: O(n log k) time, O(k) space.
# Interview line: add the list index so equal node values do not require comparing ListNode objects.
def merge_k_lists(lists: List[Optional[ListNode]]) -> Optional[ListNode]:
    heap = []
    for idx, node in enumerate(lists):
        if node:
            heapq.heappush(heap, (node.val, idx, node))

    dummy = ListNode()
    tail = dummy
    while heap:
        _, idx, node = heapq.heappop(heap)
        tail.next = node
        tail = tail.next
        if node.next:
            heapq.heappush(heap, (node.next.val, idx, node.next))

    return dummy.next


# Pattern: k-way merge for sorted timestamp streams.
# Invariant: heap stores the next log event from each stream.
# Complexity: O(n log k) time, O(k) space.
# Interview line: this is merge-k-lists framed as ordered log aggregation.
def merge_sorted_log_streams(streams: List[List[Tuple[int, str]]]) -> List[Tuple[int, str]]:
    heap = []
    for stream_id, stream in enumerate(streams):
        if stream:
            ts, msg = stream[0]
            heapq.heappush(heap, (ts, stream_id, 0, msg))

    merged = []
    while heap:
        ts, stream_id, idx, msg = heapq.heappop(heap)
        merged.append((ts, msg))
        next_idx = idx + 1
        if next_idx < len(streams[stream_id]):
            next_ts, next_msg = streams[stream_id][next_idx]
            heapq.heappush(heap, (next_ts, stream_id, next_idx, next_msg))

    return merged


# Pattern: max-heap simulation plus cooldown batches.
# Invariant: each cycle schedules up to n + 1 most frequent remaining tasks.
# Complexity: O(t log u) time, O(u) space where u is unique tasks.
# Interview line: greedily run the most frequent tasks first to minimize idle slots.
def least_interval(tasks: List[str], n: int) -> int:
    counts = [-count for count in Counter(tasks).values()]
    heapq.heapify(counts)
    time = 0

    while counts:
        cooldown = []
        slots = n + 1

        while slots and counts:
            count = heapq.heappop(counts) + 1
            if count:
                cooldown.append(count)
            time += 1
            slots -= 1

        for count in cooldown:
            heapq.heappush(counts, count)

        if counts:
            time += slots

    return time


class MedianFinder:
    # Pattern: two heaps for streaming median.
    # Invariant: small has the lower half as negatives, large has the upper half.
    # Complexity: O(1) time, O(1) space.
    # Interview line: initialize both heaps so inserts can rebalance around the median.
    def __init__(self):
        self.small = []
        self.large = []

    # Pattern: push to max side, move one to min side, then rebalance sizes.
    # Invariant: len(small) is either equal to len(large) or one larger.
    # Complexity: O(log n) time, O(n) total space.
    # Interview line: balancing heaps keeps the median at the heap tops.
    def add_num(self, num: int) -> None:
        heapq.heappush(self.small, -num)
        heapq.heappush(self.large, -heapq.heappop(self.small))
        if len(self.large) > len(self.small):
            heapq.heappush(self.small, -heapq.heappop(self.large))

    # Pattern: read median from heap tops.
    # Invariant: heap sizes determine whether median is one top or average of two tops.
    # Complexity: O(1) time, O(1) extra space.
    # Interview line: after every insert, the median is immediately available.
    def find_median(self) -> float:
        if len(self.small) > len(self.large):
            return float(-self.small[0])
        return (-self.small[0] + self.large[0]) / 2


if __name__ == "__main__":
    assert kth_largest([3, 2, 1, 5, 6, 4], 2) == 5
    assert merge_k_sorted_arrays([[1, 4], [1, 3], [2]]) == [1, 1, 2, 3, 4]
