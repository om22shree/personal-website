package dsa

// Pattern: 2D longest common subsequence DP.
// Invariant: dp[r][c] is LCS length for text1[:r] and text2[:c].
// Complexity: O(mn) time, O(mn) space.
// Interview line: matching chars extend diagonal; otherwise keep best without one char.
func LongestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1)
	n := len(text2)
	dp := make([][]int, m+1)
	for r := 0; r <= m; r++ {
		dp[r] = make([]int, n+1)
	}

	for r := 1; r <= m; r++ {
		for c := 1; c <= n; c++ {
			if text1[r-1] == text2[c-1] {
				dp[r][c] = dp[r-1][c-1] + 1
			} else {
				dp[r][c] = dp[r-1][c]
				if dp[r][c-1] > dp[r][c] {
					dp[r][c] = dp[r][c-1]
				}
			}
		}
	}
	return dp[m][n]
}

// Pattern: edit distance DP.
// Invariant: dp[r][c] is min edits from word1[:r] to word2[:c].
// Complexity: O(mn) time, O(mn) space.
// Interview line: insert, delete, replace are the three ways to repair the last character.
func MinDistance(word1 string, word2 string) int {
	m := len(word1)
	n := len(word2)
	dp := make([][]int, m+1)
	for r := 0; r <= m; r++ {
		dp[r] = make([]int, n+1)
		dp[r][0] = r
	}
	for c := 0; c <= n; c++ {
		dp[0][c] = c
	}

	for r := 1; r <= m; r++ {
		for c := 1; c <= n; c++ {
			if word1[r-1] == word2[c-1] {
				dp[r][c] = dp[r-1][c-1]
			} else {
				// edit operations: replace, delete, insert
				minVal := dp[r-1][c-1] // replace
				if dp[r-1][c] < minVal {
					minVal = dp[r-1][c] // delete
				}
				if dp[r][c-1] < minVal {
					minVal = dp[r][c-1] // insert
				}
				dp[r][c] = minVal + 1
			}
		}
	}
	return dp[m][n]
}

// Pattern: grid path-count DP with blockers.
// Invariant: dp[c] counts ways to reach current row and column c.
// Complexity: O(mn) time, O(n) space.
// Interview line: each open cell receives paths from top and left.
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || obstacleGrid[0][0] == 1 {
		return 0
	}
	n := len(obstacleGrid[0])
	dp := make([]int, n)
	dp[0] = 1

	for r := 0; r < len(obstacleGrid); r++ {
		for c := 0; c < n; c++ {
			if obstacleGrid[r][c] == 1 {
				dp[c] = 0
			} else if c > 0 {
				dp[c] += dp[c-1]
			}
		}
	}
	return dp[n-1]
}

// Pattern: grid min-cost DP.
// Invariant: dp[c] is min cost to reach current cell in column c.
// Complexity: O(mn) time, O(n) space.
// Interview line: each cell adds its cost to the cheaper of top or left.
func MinPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	cols := len(grid[0])
	dp := make([]int, cols)
	for i := range dp {
		dp[i] = 1e9 // infinite sentinel
	}
	dp[0] = 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < cols; c++ {
			if c == 0 {
				dp[c] += grid[r][c]
			} else {
				minVal := dp[c]
				if dp[c-1] < minVal {
					minVal = dp[c-1]
				}
				dp[c] = minVal + grid[r][c]
			}
		}
	}
	return dp[cols-1]
}

// Pattern: expand around centers.
// Invariant: each palindrome is discovered from its middle.
// Complexity: O(n^2) time, O(1) extra space.
// Interview line: palindromes grow symmetrically from one or two centers.
func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	start := 0
	end := 0

	expand := func(l int, r int) int {
		for l >= 0 && r < len(s) && s[l] == s[r] {
			l--
			r++
		}
		return r - l - 1
	}

	for i := 0; i < len(s); i++ {
		len1 := expand(i, i)
		len2 := expand(i, i+1)
		maxLen := len1
		if len2 > maxLen {
			maxLen = len2
		}
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

// Pattern: bottom-up palindrome table.
// Invariant: dp[l][r] means s[l:r+1] is a palindrome.
// Complexity: O(n^2) time, O(n^2) space.
// Interview line: a substring is palindromic if ends match and the inside is palindromic.
func CountSubstrings(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	count := 0

	for length := 1; length <= n; length++ {
		for l := 0; l <= n-length; l++ {
			r := l + length - 1
			if s[l] == s[r] {
				if length <= 2 || dp[l+1][r-1] {
					dp[l][r] = true
					count++
				}
			}
		}
	}
	return count
}

// Pattern: 2D DP over transaction state.
// Invariant: hold/sold/rest track best profit ending in each state after current day.
// Complexity: O(n) time, O(1) space.
// Interview line: cooldown means a buy can only come from yesterday's rest state.
func MaxProfitWithCooldown(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	hold := -prices[0]
	sold := 0
	rest := 0

	for i := 1; i < len(prices); i++ {
		newHold := hold
		if rest-prices[i] > newHold {
			newHold = rest - prices[i]
		}

		newSold := hold + prices[i]

		newRest := rest
		if sold > newRest {
			newRest = sold
		}

		hold = newHold
		sold = newSold
		rest = newRest
	}

	if sold > rest {
		return sold
	}
	return rest
}
