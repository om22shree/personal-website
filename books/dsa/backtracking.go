package dsa

import "sort"

// Pattern: include/exclude backtracking.
// Invariant: path contains chosen values from nums[:i].
// Complexity: O(n 2^n) time, O(n) recursion space excluding output.
// Interview line: every element has two choices: take it or skip it.
func Subsets(nums []int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		// Decision 1: include nums[i]
		path = append(path, nums[i])
		backtrack(i + 1)
		path = path[:len(path)-1]

		// Decision 2: exclude nums[i]
		backtrack(i + 1)
	}

	backtrack(0)
	return res
}

// Pattern: choice tree with reusable candidates.
// Invariant: path sum is current, and choices start at index start.
// Complexity: O(branches^target) time, O(target) recursion space.
// Interview line: stay on the same index when a candidate can be reused.
func CombinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	var backtrack func(start int, currentSum int)
	backtrack = func(start int, currentSum int) {
		if currentSum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if currentSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			path = append(path, candidates[i])
			// Stay at index `i` because we can reuse elements
			backtrack(i, currentSum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}

// Pattern: sorted backtracking with duplicate skip.
// Invariant: each candidate index is used at most once in a path.
// Complexity: O(2^n) time, O(n) recursion space excluding output.
// Interview line: skip equal candidates at the same recursion depth to avoid duplicate sets.
func CombinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	var path []int

	var backtrack func(start int, currentSum int)
	backtrack = func(start int, currentSum int) {
		if currentSum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		if currentSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue // Skip duplicates at the same level
			}
			path = append(path, candidates[i])
			backtrack(i+1, currentSum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return res
}

// Pattern: permutation by used flags.
// Invariant: path contains each used value exactly once.
// Complexity: O(n n!) time, O(n) recursion space excluding output.
// Interview line: permutations choose any unused value for the next slot.
func Permute(nums []int) [][]int {
	var res [][]int
	var path []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(path) == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack()
			path = path[:len(path)-1]
			used[i] = false
		}
	}

	backtrack()
	return res
}

// Pattern: duplicate-aware subset generation.
// Invariant: equal values are only skipped at the same choice depth.
// Complexity: O(n 2^n) time, O(n) recursion space excluding output.
// Interview line: sort first so duplicates are adjacent and easy to skip.
func SubsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	var path []int

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == len(nums) {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		// Take the element
		path = append(path, nums[i])
		backtrack(i + 1)
		path = path[:len(path)-1]

		// Skip all subsequent duplicates
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
		backtrack(i + 1)
	}

	backtrack(0)
	return res
}

// Pattern: grid DFS with temporary marking.
// Invariant: each board cell is used at most once in the current path.
// Complexity: O(mn 4^L) time, O(L) recursion space.
// Interview line: mark before exploring neighbors and restore before returning.
func Exist(board [][]byte, word string) bool {
	rows := len(board)
	if rows == 0 {
		return false
	}
	cols := len(board[0])

	var dfs func(r int, c int, wIdx int) bool
	dfs = func(r int, c int, wIdx int) bool {
		if wIdx == len(word) {
			return true
		}
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[wIdx] {
			return false
		}

		// Save original value and mark cell as visited
		temp := board[r][c]
		board[r][c] = '#'

		// Explore neighbors
		found := dfs(r+1, c, wIdx+1) ||
			dfs(r-1, c, wIdx+1) ||
			dfs(r, c+1, wIdx+1) ||
			dfs(r, c-1, wIdx+1)

		// Backtrack
		board[r][c] = temp
		return found
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if dfs(r, c, 0) {
				return true
			}
		}
	}
	return false
}

func isPalindromeString(s string, left int, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

// Pattern: palindrome partition backtracking.
// Invariant: path contains palindromic chunks covering s[:start].
// Complexity: O(n 2^n) time, O(n) recursion space excluding output.
// Interview line: choose the next cut only if the chosen substring is a palindrome.
func PalindromePartition(s string) [][]string {
	var res [][]string
	var path []string

	var backtrack func(start int)
	backtrack = func(start int) {
		if start == len(s) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}

		for i := start; i < len(s); i++ {
			if isPalindromeString(s, start, i) {
				path = append(path, s[start:i+1])
				backtrack(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	backtrack(0)
	return res
}
