from typing import List


# Pattern: 2D longest common subsequence DP.
# Invariant: dp[r][c] is LCS length for text1[:r] and text2[:c].
# Complexity: O(mn) time, O(mn) space.
# Interview line: matching chars extend diagonal; otherwise keep best without one char.
def longest_common_subsequence(text1: str, text2: str) -> int:
    rows, cols = len(text1), len(text2)
    dp = [[0] * (cols + 1) for _ in range(rows + 1)]
    for r in range(1, rows + 1):
        for c in range(1, cols + 1):
            if text1[r - 1] == text2[c - 1]:
                dp[r][c] = 1 + dp[r - 1][c - 1]
            else:
                dp[r][c] = max(dp[r - 1][c], dp[r][c - 1])
    return dp[rows][cols]


# Pattern: edit distance DP.
# Invariant: dp[r][c] is min edits from word1[:r] to word2[:c].
# Complexity: O(mn) time, O(mn) space.
# Interview line: insert, delete, replace are the three ways to repair the last character.
def edit_distance(word1: str, word2: str) -> int:
    rows, cols = len(word1), len(word2)
    dp = [[0] * (cols + 1) for _ in range(rows + 1)]
    for r in range(rows + 1):
        dp[r][0] = r
    for c in range(cols + 1):
        dp[0][c] = c
    for r in range(1, rows + 1):
        for c in range(1, cols + 1):
            if word1[r - 1] == word2[c - 1]:
                dp[r][c] = dp[r - 1][c - 1]
            else:
                dp[r][c] = 1 + min(dp[r - 1][c], dp[r][c - 1], dp[r - 1][c - 1])
    return dp[rows][cols]


# Pattern: grid path-count DP with blockers.
# Invariant: dp[c] counts ways to reach current row and column c.
# Complexity: O(mn) time, O(n) space.
# Interview line: each open cell receives paths from top and left.
def unique_paths_with_obstacles(obstacle_grid: List[List[int]]) -> int:
    cols = len(obstacle_grid[0]) if obstacle_grid else 0
    dp = [0] * cols
    dp[0] = 1
    for row in obstacle_grid:
        for c in range(cols):
            if row[c] == 1:
                dp[c] = 0
            elif c > 0:
                dp[c] += dp[c - 1]
    return dp[-1] if cols else 0


# Pattern: grid min-cost DP.
# Invariant: dp[c] is min cost to reach current cell in column c.
# Complexity: O(mn) time, O(n) space.
# Interview line: each cell adds its cost to the cheaper of top or left.
def min_path_sum(grid: List[List[int]]) -> int:
    if not grid:
        return 0
    cols = len(grid[0])
    dp = [float("inf")] * cols
    dp[0] = 0
    for row in grid:
        for c in range(cols):
            left = dp[c - 1] if c > 0 else float("inf")
            dp[c] = row[c] + min(dp[c], left)
    return int(dp[-1])


# Pattern: expand around centers.
# Invariant: each palindrome is discovered from its middle.
# Complexity: O(n^2) time, O(1) extra space.
# Interview line: palindromes grow symmetrically from one or two centers.
def longest_palindromic_substring(s: str) -> str:
    best = (0, 0)

    def expand(left: int, right: int) -> None:
        nonlocal best
        while left >= 0 and right < len(s) and s[left] == s[right]:
            left -= 1
            right += 1
        if right - left - 2 > best[1] - best[0]:
            best = (left + 1, right - 1)

    for i in range(len(s)):
        expand(i, i)
        expand(i, i + 1)

    return s[best[0] : best[1] + 1]


# Pattern: bottom-up palindrome table.
# Invariant: dp[l][r] means s[l:r+1] is a palindrome.
# Complexity: O(n^2) time, O(n^2) space.
# Interview line: a substring is palindromic if ends match and the inside is palindromic.
def count_palindromic_substrings(s: str) -> int:
    n = len(s)
    dp = [[False] * n for _ in range(n)]
    count = 0
    for length in range(1, n + 1):
        for left in range(n - length + 1):
            right = left + length - 1
            if s[left] == s[right] and (length <= 2 or dp[left + 1][right - 1]):
                dp[left][right] = True
                count += 1
    return count


# Pattern: 2D DP over transaction state.
# Invariant: hold/sold/rest track best profit ending in each state after current day.
# Complexity: O(n) time, O(1) space.
# Interview line: cooldown means a buy can only come from yesterday's rest state.
def max_profit_with_cooldown(prices: List[int]) -> int:
    hold = float("-inf")
    sold = 0
    rest = 0
    for price in prices:
        prev_hold, prev_sold, prev_rest = hold, sold, rest
        hold = max(prev_hold, prev_rest - price)
        sold = prev_hold + price
        rest = max(prev_rest, prev_sold)
    return int(max(sold, rest))
