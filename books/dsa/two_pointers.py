from typing import List


# Pattern: sorted array + two pointers from both ends.
# Invariant: move left/right based on whether total is too small or too large.
# Complexity: O(n) time, O(1) extra space.
# Interview line: sorted order lets us adjust the sum directionally.
def two_sum_sorted(numbers: List[int], target: int) -> List[int]:
    left, right = 0, len(numbers) - 1

    while left < right:
        total = numbers[left] + numbers[right]
        if total == target:
            return [left + 1, right + 1]
        if total < target:
            left += 1
        else:
            right -= 1

    return []


# Pattern: sort + fixed anchor + two pointers.
# Invariant: move left/right based on whether total is too small or too large.
# Complexity: O(n^2) time, O(1) extra space excluding output.
# Interview line: sorting lets us skip duplicates and adjust the sum directionally.
def three_sum(nums: List[int]) -> List[List[int]]:
    nums.sort()
    result = []

    for i, value in enumerate(nums):
        if i > 0 and value == nums[i - 1]:
            continue

        left, right = i + 1, len(nums) - 1
        while left < right:
            total = value + nums[left] + nums[right]
            if total == 0:
                result.append([value, nums[left], nums[right]])
                left += 1
                right -= 1
                while left < right and nums[left] == nums[left - 1]:
                    left += 1
                while left < right and nums[right] == nums[right + 1]:
                    right -= 1
            elif total < 0:
                left += 1
            else:
                right -= 1

    return result


# Pattern: two pointers maximizing width times limiting height.
# Invariant: moving the shorter side is the only move that can improve area.
# Complexity: O(n) time, O(1) extra space.
# Interview line: width only shrinks, so we discard the shorter wall.
def max_area(height: List[int]) -> int:
    left, right = 0, len(height) - 1
    best = 0

    while left < right:
        width = right - left
        best = max(best, width * min(height[left], height[right]))
        if height[left] < height[right]:
            left += 1
        else:
            right -= 1

    return best


# Pattern: two pointers with left/right max boundaries.
# Invariant: water on the lower side is decided by that side's max.
# Complexity: O(n) time, O(1) extra space.
# Interview line: process the smaller boundary because the opposite side already bounds it.
def trap_rain_water(height: List[int]) -> int:
    left, right = 0, len(height) - 1
    left_max = right_max = 0
    water = 0

    while left < right:
        if height[left] < height[right]:
            left_max = max(left_max, height[left])
            water += left_max - height[left]
            left += 1
        else:
            right_max = max(right_max, height[right])
            water += right_max - height[right]
            right -= 1

    return water


# Pattern: read/write pointers on a sorted array.
# Invariant: nums[:write] contains the unique prefix.
# Complexity: O(n) time, O(1) extra space.
# Interview line: read scans everything, write only advances when a new value appears.
def remove_duplicates(nums: List[int]) -> int:
    if not nums:
        return 0

    write = 1
    for read in range(1, len(nums)):
        if nums[read] != nums[read - 1]:
            nums[write] = nums[read]
            write += 1

    return write


# Pattern: stable partition with read/write pointers.
# Invariant: nums[:write] contains all non-zero values seen so far.
# Complexity: O(n) time, O(1) extra space.
# Interview line: swap each non-zero into the next write slot while preserving order.
def move_zeroes(nums: List[int]) -> None:
    write = 0
    for read, value in enumerate(nums):
        if value != 0:
            nums[write], nums[read] = nums[read], nums[write]
            write += 1


# Pattern: two pointers over two strings.
# Invariant: i is the next character of s we still need to match.
# Complexity: O(len(t)) time, O(1) extra space.
# Interview line: scan the larger string once and advance the subsequence pointer on matches.
def is_subsequence(s: str, t: str) -> bool:
    i = 0
    for ch in t:
        if i < len(s) and s[i] == ch:
            i += 1
    return i == len(s)


# Pattern: inward two pointers with filtering.
# Invariant: compare only alphanumeric characters after normalization.
# Complexity: O(n) time, O(1) extra space.
# Interview line: skip irrelevant characters, then compare the normalized ends.
def valid_palindrome(s: str) -> bool:
    left, right = 0, len(s) - 1

    while left < right:
        while left < right and not s[left].isalnum():
            left += 1
        while left < right and not s[right].isalnum():
            right -= 1

        if s[left].lower() != s[right].lower():
            return False
        left += 1
        right -= 1

    return True


# Pattern: Dutch national flag with three regions.
# Invariant: left side is 0s, middle is 1s, right side is 2s.
# Complexity: O(n) time, O(1) extra space.
# Interview line: one pass partitions the array because there are only three values.
def sort_colors(nums: List[int]) -> None:
    left = mid = 0
    right = len(nums) - 1

    while mid <= right:
        if nums[mid] == 0:
            nums[left], nums[mid] = nums[mid], nums[left]
            left += 1
            mid += 1
        elif nums[mid] == 1:
            mid += 1
        else:
            nums[mid], nums[right] = nums[right], nums[mid]
            right -= 1


if __name__ == "__main__":
    assert two_sum_sorted([2, 7, 11, 15], 9) == [1, 2]
    assert sorted(three_sum([-1, 0, 1, 2, -1, -4])) == [[-1, -1, 2], [-1, 0, 1]]
    assert max_area([1, 8, 6, 2, 5, 4, 8, 3, 7]) == 49
    assert trap_rain_water([0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1]) == 6
