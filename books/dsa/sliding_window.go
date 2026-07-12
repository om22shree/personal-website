package dsa

// Pattern: fixed-size sliding window with rolling sum.
// Invariant: window always contains exactly k elements after the first k values.
// Complexity: O(n) time, O(1) extra space.
// Interview line: subtract the outgoing element and add the incoming one instead of recomputing.
func FindMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}

// Pattern: variable sliding window with last-seen index.
// Invariant: current window has no duplicate characters.
// Complexity: O(n) time, O(min(n, charset)) space.
// Interview line: when a duplicate appears inside the window, jump left past its last position.
func LengthOfLongestSubstring(s string) int {
	lastSeen := make(map[rune]int)
	maxLen := 0
	left := 0
	for right, char := range s {
		if idx, ok := lastSeen[char]; ok && idx >= left {
			left = idx + 1
		}
		lastSeen[char] = right
		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// Pattern: variable sliding window with required character counts.
// Invariant: shrink only while the window satisfies every needed count.
// Complexity: O(len(s) + len(t)) time, O(len(t)) space.
// Interview line: expand to become valid, then shrink greedily to find the smallest valid window.
func MinWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	window := make(map[byte]int)
	have := 0
	required := len(need)

	bestLeft := -1
	minLen := len(s) + 1

	left := 0
	for right := 0; right < len(s); right++ {
		char := s[right]
		window[char]++

		if count, ok := need[char]; ok && window[char] == count {
			have++
		}

		for have == required {
			if right-left+1 < minLen {
				minLen = right - left + 1
				bestLeft = left
			}
			// pop left
			leftChar := s[left]
			window[leftChar]--
			if count, ok := need[leftChar]; ok && window[leftChar] < count {
				have--
			}
			left++
		}
	}

	if bestLeft == -1 {
		return ""
	}
	return s[bestLeft : bestLeft+minLen]
}

// Pattern: monotonic deque for fixed-size window maximum.
// Invariant: deque stores candidate indices in decreasing value order.
// Complexity: O(n) time, O(k) space.
// Interview line: smaller elements behind a larger incoming value can never become the max.
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return nil
	}
	var q []int // stores indices
	var res []int

	for i, num := range nums {
		// remove indices out of window
		if len(q) > 0 && q[0] < i-k+1 {
			q = q[1:]
		}
		// maintain monotonic decreasing deque
		for len(q) > 0 && nums[q[len(q)-1]] < num {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		// window is fully formed
		if i >= k-1 {
			res = append(res, nums[q[0]])
		}
	}
	return res
}

// Pattern: fixed-size sliding window with frequency maps.
// Invariant: window length never exceeds len(p), and matching counts mark an anagram.
// Complexity: O(n) time with bounded alphabet, O(len(p)) space.
// Interview line: keep the window exactly the pattern length, then compare character counts.
func FindAnagrams(s string, p string) []int {
	var res []int
	if len(s) < len(p) {
		return res
	}

	var pCounts, sCounts [26]int
	for i := 0; i < len(p); i++ {
		pCounts[p[i]-'a']++
		sCounts[s[i]-'a']++
	}

	if sCounts == pCounts {
		res = append(res, 0)
	}

	for i := len(p); i < len(s); i++ {
		sCounts[s[i]-'a']++
		sCounts[s[i-len(p)]-'a']--
		if sCounts == pCounts {
			res = append(res, i-len(p)+1)
		}
	}
	return res
}

// Pattern: variable sliding window with most frequent character tracking.
// Invariant: window is valid when replacements needed <= k.
// Complexity: O(n) time, O(1) space for uppercase English letters.
// Interview line: the best target character is the most frequent one already in the window.
func CharacterReplacement(s string, k int) int {
	counts := make(map[byte]int)
	maxFreq := 0
	left := 0
	maxLen := 0

	for right := 0; right < len(s); right++ {
		counts[s[right]]++
		if counts[s[right]] > maxFreq {
			maxFreq = counts[s[right]]
		}

		// window is valid if: width - maxFreq <= k
		for (right-left+1)-maxFreq > k {
			counts[s[left]]--
			left++
			// note: we do not strictly need to decrement maxFreq here because
			// a larger valid window would require an even larger maxFreq anyway.
		}

		if right-left+1 > maxLen {
			maxLen = right - left + 1
		}
	}
	return maxLen
}
