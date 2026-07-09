from collections import Counter, defaultdict
from typing import Dict, List


# Pattern: hashmap complement lookup.
# Invariant: seen maps every previous value to its index.
# Complexity: O(n) time, O(n) space.
# Interview line: for pair sums, store what you have seen and ask whether the complement exists.
def two_sum(nums: List[int], target: int) -> List[int]:
    seen: Dict[int, int] = {}
    for i, num in enumerate(nums):
        need = target - num
        if need in seen:
            return [seen[need], i]
        seen[num] = i
    return []


# Pattern: set membership duplicate check.
# Invariant: seen contains exactly the distinct values processed so far.
# Complexity: O(n) time, O(n) space.
# Interview line: a duplicate is the first value already present in the set.
def contains_duplicate(nums: List[int]) -> bool:
    seen = set()
    for num in nums:
        if num in seen:
            return True
        seen.add(num)
    return False


# Pattern: frequency equality.
# Invariant: anagrams have identical character counts.
# Complexity: O(n) time, O(1) space for fixed alphabet.
# Interview line: order does not matter, counts do.
def is_anagram(s: str, t: str) -> bool:
    return Counter(s) == Counter(t)


# Pattern: canonical sorted key grouping.
# Invariant: words with the same sorted letters share one bucket.
# Complexity: O(n k log k) time, O(n k) space.
# Interview line: choose a key that erases the irrelevant order.
def group_anagrams(strs: List[str]) -> List[List[str]]:
    groups = defaultdict(list)
    for word in strs:
        groups["".join(sorted(word))].append(word)
    return list(groups.values())


# Pattern: prefix product plus suffix product.
# Invariant: answer[i] accumulates product of all values left and right of i.
# Complexity: O(n) time, O(1) extra space excluding output.
# Interview line: multiply prefix on the way forward, suffix on the way back.
def product_except_self(nums: List[int]) -> List[int]:
    ans = [1] * len(nums)
    prefix = 1
    for i, num in enumerate(nums):
        ans[i] = prefix
        prefix *= num

    suffix = 1
    for i in range(len(nums) - 1, -1, -1):
        ans[i] *= suffix
        suffix *= nums[i]

    return ans


# Pattern: length-prefixed string encoding.
# Invariant: each encoded string is length, delimiter, then exact payload.
# Complexity: O(total chars) time, O(total chars) space.
# Interview line: length prefixes avoid delimiter ambiguity inside strings.
def encode_strings(strs: List[str]) -> str:
    return "".join(f"{len(s)}#{s}" for s in strs)


# Pattern: parse length-prefixed strings.
# Invariant: i always points to the start of the next length field.
# Complexity: O(total chars) time, O(total chars) space.
# Interview line: read the length first, then jump exactly that many characters.
def decode_strings(data: str) -> List[str]:
    ans = []
    i = 0
    while i < len(data):
        j = data.index("#", i)
        length = int(data[i:j])
        start = j + 1
        ans.append(data[start : start + length])
        i = start + length
    return ans


# Pattern: row/column/box set validation.
# Invariant: no non-dot value repeats in its row, column, or 3x3 box.
# Complexity: O(1) time, O(1) space for fixed 9x9 board.
# Interview line: each filled cell belongs to three uniqueness constraints.
def valid_sudoku(board: List[List[str]]) -> bool:
    rows = [set() for _ in range(9)]
    cols = [set() for _ in range(9)]
    boxes = [set() for _ in range(9)]

    for r in range(9):
        for c in range(9):
            val = board[r][c]
            if val == ".":
                continue
            box = (r // 3) * 3 + c // 3
            if val in rows[r] or val in cols[c] or val in boxes[box]:
                return False
            rows[r].add(val)
            cols[c].add(val)
            boxes[box].add(val)
    return True


# Pattern: hash-set sequence starts.
# Invariant: only numbers without a predecessor start counted sequences.
# Complexity: O(n) time, O(n) space.
# Interview line: count each consecutive run once, from its left edge.
def longest_consecutive(nums: List[int]) -> int:
    values = set(nums)
    best = 0
    for num in values:
        if num - 1 in values:
            continue
        cur = num
        while cur in values:
            cur += 1
        best = max(best, cur - num)
    return best


# Pattern: prefix sum frequency.
# Invariant: counts[prefix] is how many earlier prefixes had that sum.
# Complexity: O(n) time, O(n) space.
# Interview line: subarray sum equals k when current_prefix - old_prefix equals k.
def subarray_sum_equals_k(nums: List[int], k: int) -> int:
    counts = defaultdict(int)
    counts[0] = 1
    prefix = 0
    ans = 0
    for num in nums:
        prefix += num
        ans += counts[prefix - k]
        counts[prefix] += 1
    return ans


# Pattern: Boyer-Moore majority vote.
# Invariant: candidate survives pair cancellations against other values.
# Complexity: O(n) time, O(1) space.
# Interview line: the true majority cannot be fully canceled by all non-majority values.
def majority_element(nums: List[int]) -> int:
    candidate = None
    votes = 0
    for num in nums:
        if votes == 0:
            candidate = num
        votes += 1 if num == candidate else -1
    return int(candidate)


class NumMatrix:
    # Pattern: 2D prefix sums.
    # Invariant: prefix[r][c] stores sum of rectangle [0:r) x [0:c).
    # Complexity: O(mn) setup, O(mn) space.
    # Interview line: inclusion-exclusion turns rectangle queries into four prefix reads.
    def __init__(self, matrix: List[List[int]]):
        rows = len(matrix)
        cols = len(matrix[0]) if rows else 0
        self.prefix = [[0] * (cols + 1) for _ in range(rows + 1)]
        for r in range(rows):
            row_sum = 0
            for c in range(cols):
                row_sum += matrix[r][c]
                self.prefix[r + 1][c + 1] = self.prefix[r][c + 1] + row_sum

    # Pattern: 2D prefix sum query.
    # Invariant: bottom/right indexes are converted to exclusive prefix coordinates.
    # Complexity: O(1) time, O(1) space.
    # Interview line: add big rectangle, subtract strips, add back overlap.
    def sum_region(self, row1: int, col1: int, row2: int, col2: int) -> int:
        p = self.prefix
        return (
            p[row2 + 1][col2 + 1]
            - p[row1][col2 + 1]
            - p[row2 + 1][col1]
            + p[row1][col1]
        )
