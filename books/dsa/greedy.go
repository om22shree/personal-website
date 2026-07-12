package dsa

import "sort"

// Pattern: one-pass minimum price.
// Invariant: min_price is the lowest price before the current day.
// Complexity: O(n) time, O(1) space.
// Interview line: sell today against the cheapest earlier buy.
func MaxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProf := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		}
		prof := p - minPrice
		if prof > maxProf {
			maxProf = prof
		}
	}
	return maxProf
}

// Pattern: collect every upward edge.
// Invariant: every positive day-to-day gain can be taken independently.
// Complexity: O(n) time, O(1) space.
// Interview line: unlimited transactions reduce to summing all positive slopes.
func MaxProfitMany(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

// Pattern: greedy farthest reach.
// Invariant: reach is the farthest index reachable from processed positions.
// Complexity: O(n) time, O(1) space.
// Interview line: if you ever stand beyond reach, the end is impossible.
func CanJump(nums []int) bool {
	reach := 0
	for i, num := range nums {
		if i > reach {
			return false
		}
		if i+num > reach {
			reach = i + num
		}
	}
	return true
}

// Pattern: BFS levels without queue.
// Invariant: current_end bounds indexes reachable with jumps jumps.
// Complexity: O(n) time, O(1) space.
// Interview line: when you exhaust the current range, commit one jump to the farthest next range.
func JumpGameII(nums []int) int {
	jumps := 0
	currentEnd := 0
	farthest := 0

	for i := 0; i < len(nums)-1; i++ {
		if i+nums[i] > farthest {
			farthest = i + nums[i]
		}
		if i == currentEnd {
			jumps++
			currentEnd = farthest
		}
	}
	return jumps
}

// Pattern: gas tank surplus reset.
// Invariant: if tank drops below zero, no station in this failed segment can start.
// Complexity: O(n) time, O(1) space.
// Interview line: total surplus proves existence; local deficit chooses the next start.
func CanCompleteCircuit(gas []int, cost []int) int {
	totalSurplus := 0
	currentSurplus := 0
	startStation := 0

	for i := 0; i < len(gas); i++ {
		diff := gas[i] - cost[i]
		totalSurplus += diff
		currentSurplus += diff
		if currentSurplus < 0 {
			startStation = i + 1
			currentSurplus = 0
		}
	}

	if totalSurplus < 0 {
		return -1
	}
	return startStation
}

// Pattern: last occurrence partitioning.
// Invariant: current partition must extend to the farthest last occurrence of its letters.
// Complexity: O(n) time, O(1) space.
// Interview line: close a partition only when every character inside ends inside.
func PartitionLabels(s string) []int {
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	var res []int
	start := 0
	end := 0

	for i := 0; i < len(s); i++ {
		lastIdx := last[s[i]-'a']
		if lastIdx > end {
			end = lastIdx
		}
		if i == end {
			res = append(res, end-start+1)
			start = i + 1
		}
	}
	return res
}

// Pattern: greedy range of possible open counts.
// Invariant: low/high bound how many unmatched opens are possible.
// Complexity: O(n) time, O(1) space.
// Interview line: treat star as whichever choice keeps the open-count range viable.
func CheckValidString(s string) bool {
	low := 0
	high := 0

	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == '(' {
			low++
			high++
		} else if char == ')' {
			low--
			high--
		} else {
			low--
			high++
		}
		if low < 0 {
			low = 0
		}
		if high < 0 {
			return false
		}
	}
	return low == 0
}

// Pattern: greedy target coverage.
// Invariant: matched[i] is true once some triplet supplies target[i].
// Complexity: O(n) time, O(1) space.
// Interview line: discard triplets that overshoot target; useful ones can be merged coordinate-wise.
func MergeTriplets(triplets [][]int, target []int) bool {
	matched := [3]bool{}

	for _, t := range triplets {
		if t[0] <= target[0] && t[1] <= target[1] && t[2] <= target[2] {
			if t[0] == target[0] {
				matched[0] = true
			}
			if t[1] == target[1] {
				matched[1] = true
			}
			if t[2] == target[2] {
				matched[2] = true
			}
		}
	}
	return matched[0] && matched[1] && matched[2]
}

// Pattern: sorted counting with queue of group starts.
// Invariant: opened groups wait for the next consecutive card.
// Complexity: O(n log n) time, O(n) space.
// Interview line: always satisfy the smallest card first.
func IsNStraightHand(hand []int, groupSize int) bool {
	if len(hand)%groupSize != 0 {
		return false
	}

	counts := make(map[int]int)
	for _, card := range hand {
		counts[card]++
	}

	sort.Ints(hand)

	for _, card := range hand {
		if counts[card] == 0 {
			continue
		}
		// Try to form a group starting at `card`
		for i := 0; i < groupSize; i++ {
			if counts[card+i] == 0 {
				return false
			}
			counts[card+i]--
		}
	}
	return true
}
