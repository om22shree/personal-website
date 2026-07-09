from typing import List


# Pattern: include/exclude backtracking.
# Invariant: path contains chosen values from nums[:i].
# Complexity: O(n 2^n) time, O(n) recursion space excluding output.
# Interview line: every element has two choices: take it or skip it.
def subsets(nums: List[int]) -> List[List[int]]:
    ans = []

    def dfs(i: int, path: List[int]) -> None:
        if i == len(nums):
            ans.append(path[:])
            return
        dfs(i + 1, path)
        path.append(nums[i])
        dfs(i + 1, path)
        path.pop()

    dfs(0, [])
    return ans


# Pattern: choice tree with reusable candidates.
# Invariant: path sum is current, and choices start at index start.
# Complexity: O(branches^target) time, O(target) recursion space.
# Interview line: stay on the same index when a candidate can be reused.
def combination_sum(candidates: List[int], target: int) -> List[List[int]]:
    ans = []
    candidates.sort()

    def dfs(start: int, remaining: int, path: List[int]) -> None:
        if remaining == 0:
            ans.append(path[:])
            return
        for i in range(start, len(candidates)):
            if candidates[i] > remaining:
                break
            path.append(candidates[i])
            dfs(i, remaining - candidates[i], path)
            path.pop()

    dfs(0, target, [])
    return ans


# Pattern: sorted backtracking with duplicate skip.
# Invariant: each candidate index is used at most once in a path.
# Complexity: O(2^n) time, O(n) recursion space excluding output.
# Interview line: skip equal candidates at the same recursion depth to avoid duplicate sets.
def combination_sum2(candidates: List[int], target: int) -> List[List[int]]:
    ans = []
    candidates.sort()

    def dfs(start: int, remaining: int, path: List[int]) -> None:
        if remaining == 0:
            ans.append(path[:])
            return
        prev = None
        for i in range(start, len(candidates)):
            if candidates[i] == prev:
                continue
            if candidates[i] > remaining:
                break
            prev = candidates[i]
            path.append(candidates[i])
            dfs(i + 1, remaining - candidates[i], path)
            path.pop()

    dfs(0, target, [])
    return ans


# Pattern: permutation by used flags.
# Invariant: path contains each used value exactly once.
# Complexity: O(n n!) time, O(n) recursion space excluding output.
# Interview line: permutations choose any unused value for the next slot.
def permutations(nums: List[int]) -> List[List[int]]:
    ans = []
    used = [False] * len(nums)

    def dfs(path: List[int]) -> None:
        if len(path) == len(nums):
            ans.append(path[:])
            return
        for i, num in enumerate(nums):
            if used[i]:
                continue
            used[i] = True
            path.append(num)
            dfs(path)
            path.pop()
            used[i] = False

    dfs([])
    return ans


# Pattern: duplicate-aware subset generation.
# Invariant: equal values are only skipped at the same choice depth.
# Complexity: O(n 2^n) time, O(n) recursion space excluding output.
# Interview line: sort first so duplicates are adjacent and easy to skip.
def subsets_with_dup(nums: List[int]) -> List[List[int]]:
    ans = []
    nums.sort()

    def dfs(start: int, path: List[int]) -> None:
        ans.append(path[:])
        for i in range(start, len(nums)):
            if i > start and nums[i] == nums[i - 1]:
                continue
            path.append(nums[i])
            dfs(i + 1, path)
            path.pop()

    dfs(0, [])
    return ans


# Pattern: grid DFS with temporary marking.
# Invariant: each board cell is used at most once in the current path.
# Complexity: O(mn 4^L) time, O(L) recursion space.
# Interview line: mark before exploring neighbors and restore before returning.
def word_exists(board: List[List[str]], word: str) -> bool:
    rows = len(board)
    cols = len(board[0]) if rows else 0

    def dfs(r: int, c: int, i: int) -> bool:
        if i == len(word):
            return True
        if r < 0 or c < 0 or r == rows or c == cols or board[r][c] != word[i]:
            return False
        saved = board[r][c]
        board[r][c] = "#"
        found = (
            dfs(r + 1, c, i + 1)
            or dfs(r - 1, c, i + 1)
            or dfs(r, c + 1, i + 1)
            or dfs(r, c - 1, i + 1)
        )
        board[r][c] = saved
        return found

    for r in range(rows):
        for c in range(cols):
            if dfs(r, c, 0):
                return True
    return False


# Pattern: palindrome partition backtracking.
# Invariant: path contains palindromic chunks covering s[:start].
# Complexity: O(n 2^n) time, O(n) recursion space excluding output.
# Interview line: choose the next cut only if the chosen substring is a palindrome.
def palindrome_partition(s: str) -> List[List[str]]:
    ans = []

    def is_pal(left: int, right: int) -> bool:
        while left < right:
            if s[left] != s[right]:
                return False
            left += 1
            right -= 1
        return True

    def dfs(start: int, path: List[str]) -> None:
        if start == len(s):
            ans.append(path[:])
            return
        for end in range(start, len(s)):
            if is_pal(start, end):
                path.append(s[start : end + 1])
                dfs(end + 1, path)
                path.pop()

    dfs(0, [])
    return ans
