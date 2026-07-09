from collections import Counter
import heapq
from typing import List


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
