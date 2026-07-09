import bisect
import math
from typing import List


# Pattern: classic binary search on sorted array.
# Invariant: target, if present, remains inside [left, right].
# Complexity: O(log n) time, O(1) extra space.
# Interview line: each comparison discards half the remaining search space.
def binary_search(nums: List[int], target: int) -> int:
    left, right = 0, len(nums) - 1

    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid
        if nums[mid] < target:
            left = mid + 1
        else:
            right = mid - 1

    return -1


# Pattern: lower-bound binary search via bisect.
# Invariant: returned index is the first position where target can be inserted.
# Complexity: O(log n) time, O(1) extra space.
# Interview line: bisect_left gives the first candidate index for target.
def lower_bound(nums: List[int], target: int) -> int:
    return bisect.bisect_left(nums, target)


# Pattern: upper-bound binary search via bisect.
# Invariant: returned index is the first position after all target values.
# Complexity: O(log n) time, O(1) extra space.
# Interview line: bisect_right steps past duplicates, which is useful for counts.
def upper_bound(nums: List[int], target: int) -> int:
    return bisect.bisect_right(nums, target)


# Pattern: duplicate count from lower and upper bounds.
# Invariant: all target values live in [lower_bound, upper_bound).
# Complexity: O(log n) time, O(1) extra space.
# Interview line: two boundary searches avoid scanning repeated values.
def count_occurrences(nums: List[int], target: int) -> int:
    return upper_bound(nums, target) - lower_bound(nums, target)


# Pattern: binary search with one sorted half.
# Invariant: at least one side of mid is sorted each iteration.
# Complexity: O(log n) time, O(1) extra space.
# Interview line: identify the sorted side, then decide if target lies inside it.
def search_rotated(nums: List[int], target: int) -> int:
    left, right = 0, len(nums) - 1

    while left <= right:
        mid = (left + right) // 2
        if nums[mid] == target:
            return mid

        if nums[left] <= nums[mid]:
            if nums[left] <= target < nums[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if nums[mid] < target <= nums[right]:
                left = mid + 1
            else:
                right = mid - 1

    return -1


# Pattern: binary search on answer space.
# Invariant: feasible speeds form a monotonic true suffix.
# Complexity: O(n log max(piles)) time, O(1) extra space.
# Interview line: if a speed works, every larger speed also works.
def min_eating_speed(piles: List[int], h: int) -> int:
    def feasible(speed: int) -> bool:
        return sum(math.ceil(pile / speed) for pile in piles) <= h

    left, right = 1, max(piles)
    while left < right:
        mid = (left + right) // 2
        if feasible(mid):
            right = mid
        else:
            left = mid + 1
    return left


# Pattern: binary search on minimum capacity.
# Invariant: feasible capacities form a monotonic true suffix.
# Complexity: O(n log sum(weights)) time, O(1) extra space.
# Interview line: if a capacity ships in time, any larger capacity also ships in time.
def ship_within_days(weights: List[int], days: int) -> int:
    def feasible(capacity: int) -> bool:
        needed = 1
        load = 0
        for weight in weights:
            if load + weight > capacity:
                needed += 1
                load = 0
            load += weight
        return needed <= days

    left, right = max(weights), sum(weights)
    while left < right:
        mid = (left + right) // 2
        if feasible(mid):
            right = mid
        else:
            left = mid + 1
    return left


if __name__ == "__main__":
    assert binary_search([1, 3, 5], 3) == 1
    assert count_occurrences([1, 2, 2, 2, 3], 2) == 3
    assert min_eating_speed([3, 6, 7, 11], 8) == 4
