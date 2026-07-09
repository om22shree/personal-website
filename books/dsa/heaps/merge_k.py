import heapq
from typing import List, Optional, Tuple


class ListNode:
    # Pattern: linked-list node container.
    # Invariant: each node stores a value and pointer to the next node.
    # Complexity: O(1) time, O(1) space.
    # Interview line: simple node structure is enough for merge-k linked-list prompts.
    def __init__(self, val: int = 0, next: Optional["ListNode"] = None):
        self.val = val
        self.next = next


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
# Interview line: add the list index so equal node values do not compare ListNode objects.
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
