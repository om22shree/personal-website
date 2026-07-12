package dsa

// Pattern: Fibonacci-style rolling DP.
// Invariant: prev and cur hold ways for the previous two step counts.
// Complexity: O(n) time, O(1) extra space.
// Interview line: only the last two states matter, so no full DP array is needed.
func ClimbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prev := 1
	cur := 2
	for i := 3; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}

// Pattern: choose/skip rolling DP.
// Invariant: skip excludes current house, take includes current house.
// Complexity: O(n) time, O(1) extra space.
// Interview line: each house depends only on the best states before it.
func HouseRobber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	skip := 0
	take := 0
	for _, num := range nums {
		newTake := skip + num
		newSkip := skip
		if take > skip {
			newSkip = take
		}
		skip = newSkip
		take = newTake
	}
	if skip > take {
		return skip
	}
	return take
}

// Pattern: bottom-up unbounded coin DP.
// Invariant: dp[value] is the fewest coins needed for that amount.
// Complexity: O(amount * len(coins)) time, O(amount) space.
// Interview line: build answers from smaller amounts and try every coin as the last coin.
func CoinChange(coins []int, amount int) int {
	// dp size is amount + 1
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1 // sentinel representing infinity
	}
	dp[0] = 0

	for a := 1; a <= amount; a++ {
		for _, c := range coins {
			if a-c >= 0 {
				if dp[a-c]+1 < dp[a] {
					dp[a] = dp[a-c] + 1
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// Pattern: memoized DFS over string index.
// Invariant: dp(i) means s[i:] can be segmented into dictionary words.
// Complexity: O(n * max_word_len) states/checks, O(n) memo space.
// Interview line: cache the start index so repeated suffix checks are computed once.
func WordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	maxLen := 0
	for _, w := range wordDict {
		wordSet[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	memo := make(map[int]bool)
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if start == len(s) {
			return true
		}
		if val, ok := memo[start]; ok {
			return val
		}

		for end := start + 1; end <= len(s) && end-start <= maxLen; end++ {
			if wordSet[s[start:end]] && dfs(end) {
				memo[start] = true
				return true
			}
		}
		memo[start] = false
		return false
	}

	return dfs(0)
}

// Pattern: patience sorting with binary search.
// Invariant: tails[i] is the smallest possible tail for an increasing subsequence of length i + 1.
// Complexity: O(n log n) time, O(n) space.
// Interview line: smaller tails leave more room for future numbers.
func LengthOfLIS(nums []int) int {
	var tails []int
	for _, num := range nums {
		// Binary search to find where to put num
		left := 0
		right := len(tails)
		for left < right {
			mid := left + (right-left)/2
			if tails[mid] >= num {
				right = mid
			} else {
				left = mid + 1
			}
		}
		if left == len(tails) {
			tails = append(tails, num)
		} else {
			tails[left] = num
		}
	}
	return len(tails)
}

// Pattern: 2D grid DP compressed to one row.
// Invariant: row[col] stores paths to the current row and column.
// Complexity: O(m * n) time, O(n) space.
// Interview line: each cell depends only on top and left, so one row is enough.
func UniquePaths(m int, n int) int {
	row := make([]int, n)
	for i := range row {
		row[i] = 1
	}

	for r := 1; r < m; r++ {
		for c := 1; c < n; c++ {
			row[c] += row[c-1]
		}
	}
	return row[n-1]
}
