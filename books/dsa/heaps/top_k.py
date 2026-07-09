from collections import Counter
import heapq
from typing import List


# Pattern: heap selection for kth largest.
# Invariant: nlargest keeps the k best values under heap ordering.
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


# Pattern: fixed-size min-heap over frequencies.
# Invariant: heap keeps only the k highest-frequency candidates seen so far.
# Complexity: O(n log k) time, O(k) space.
# Interview line: pop the smallest frequency whenever the candidate set grows past k.
def top_k_frequent_heap(nums: List[int], k: int) -> List[int]:
    heap = []
    for num, count in Counter(nums).items():
        heapq.heappush(heap, (count, num))
        if len(heap) > k:
            heapq.heappop(heap)
    return [num for _, num in heap]


# Pattern: heap selection with custom distance key.
# Invariant: squared distance preserves ordering without sqrt.
# Complexity: O(n log k) time, O(k) space.
# Interview line: compare squared distances to avoid unnecessary floating-point work.
def k_closest_points(points: List[List[int]], k: int) -> List[List[int]]:
    return heapq.nsmallest(k, points, key=lambda p: p[0] * p[0] + p[1] * p[1])
