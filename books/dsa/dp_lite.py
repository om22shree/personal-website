import bisect
from functools import lru_cache
from typing import List


# Pattern: Fibonacci-style rolling DP.
# Invariant: prev and cur hold ways for the previous two step counts.
# Complexity: O(n) time, O(1) extra space.
# Interview line: only the last two states matter, so no full DP array is needed.
def climb_stairs(n: int) -> int:
    prev, cur = 1, 1
    for _ in range(n):
        prev, cur = cur, prev + cur
    return prev


# Pattern: choose/skip rolling DP.
# Invariant: skip excludes current house, take includes current house.
# Complexity: O(n) time, O(1) extra space.
# Interview line: each house depends only on the best states before it.
def house_robber(nums: List[int]) -> int:
    skip = take = 0
    for num in nums:
        skip, take = max(skip, take), skip + num
    return max(skip, take)


# Pattern: bottom-up unbounded coin DP.
# Invariant: dp[value] is the fewest coins needed for that amount.
# Complexity: O(amount * len(coins)) time, O(amount) space.
# Interview line: build answers from smaller amounts and try every coin as the last coin.
def coin_change(coins: List[int], amount: int) -> int:
    dp = [float("inf")] * (amount + 1)
    dp[0] = 0

    for value in range(1, amount + 1):
        for coin in coins:
            if value >= coin:
                dp[value] = min(dp[value], dp[value - coin] + 1)

    return -1 if dp[amount] == float("inf") else int(dp[amount])


# Pattern: memoized DFS over string index.
# Invariant: dp(i) means s[i:] can be segmented into dictionary words.
# Complexity: O(n * max_word_len) states/checks, O(n) memo space.
# Interview line: cache the start index so repeated suffix checks are computed once.
def word_break(s: str, word_dict: List[str]) -> bool:
    words = set(word_dict)
    max_len = max(map(len, words), default=0)

    @lru_cache(None)
    def dp(i: int) -> bool:
        if i == len(s):
            return True
        for j in range(i + 1, min(len(s), i + max_len) + 1):
            if s[i:j] in words and dp(j):
                return True
        return False

    return dp(0)


# Pattern: patience sorting with binary search.
# Invariant: tails[i] is the smallest possible tail for an increasing subsequence of length i + 1.
# Complexity: O(n log n) time, O(n) space.
# Interview line: smaller tails leave more room for future numbers.
def longest_increasing_subsequence(nums: List[int]) -> int:
    tails = []
    for num in nums:
        idx = bisect.bisect_left(tails, num)
        if idx == len(tails):
            tails.append(num)
        else:
            tails[idx] = num
    return len(tails)


# Pattern: 2D grid DP compressed to one row.
# Invariant: row[col] stores paths to the current row and column.
# Complexity: O(m * n) time, O(n) space.
# Interview line: each cell depends only on top and left, so one row is enough.
def unique_paths(m: int, n: int) -> int:
    row = [1] * n
    for _ in range(m - 1):
        for col in range(1, n):
            row[col] += row[col - 1]
    return row[-1]


if __name__ == "__main__":
    assert house_robber([2, 7, 9, 3, 1]) == 12
    assert coin_change([1, 2, 5], 11) == 3
    assert word_break("leetcode", ["leet", "code"]) is True
