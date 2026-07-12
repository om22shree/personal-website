package dsa

// Pattern: classic binary search on sorted array.
// Invariant: target, if present, remains inside [left, right].
// Complexity: O(log n) time, O(1) extra space.
// Interview line: each comparison discards half the remaining search space.
func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

// Pattern: lower-bound binary search.
// Invariant: returned index is the first position where target can be inserted.
// Complexity: O(log n) time, O(1) extra space.
// Interview line: lower_bound gives the first candidate index for target.
func LowerBound(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// Pattern: upper-bound binary search.
// Invariant: returned index is the first position after all target values.
// Complexity: O(log n) time, O(1) extra space.
// Interview line: upper_bound steps past duplicates, which is useful for counts.
func UpperBound(nums []int, target int) int {
	left := 0
	right := len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

// Pattern: duplicate count from lower and upper bounds.
// Invariant: all target values live in [lower_bound, upper_bound).
// Complexity: O(log n) time, O(1) extra space.
// Interview line: two boundary searches avoid scanning repeated values.
func CountOccurrences(nums []int, target int) int {
	return UpperBound(nums, target) - LowerBound(nums, target)
}

// Pattern: binary search with one sorted half.
// Invariant: at least one side of mid is sorted each iteration.
// Complexity: O(log n) time, O(1) extra space.
// Interview line: identify the sorted side, then decide if target lies inside it.
func SearchRotated(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		// Left side is sorted
		if nums[left] <= nums[mid] {
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Right side is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

// Helper to determine if speed works
func feasibleSpeed(piles []int, speed int, h int) bool {
	hours := 0
	for _, p := range piles {
		hours += (p + speed - 1) / speed // Ceil division
	}
	return hours <= h
}

// Pattern: binary search on answer space.
// Invariant: feasible speeds form a monotonic true suffix.
// Complexity: O(n log max(piles)) time, O(1) extra space.
// Interview line: if a speed works, every larger speed also works.
func MinEatingSpeed(piles []int, h int) int {
	maxPile := 0
	for _, p := range piles {
		if p > maxPile {
			maxPile = p
		}
	}

	left := 1
	right := maxPile
	ans := right

	for left <= right {
		mid := left + (right-left)/2
		if feasibleSpeed(piles, mid, h) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}

// Helper to check weight shipping feasibility
func feasibleCapacity(weights []int, cap int, days int) bool {
	d := 1
	curWeight := 0
	for _, w := range weights {
		if curWeight+w > cap {
			d++
			curWeight = 0
		}
		curWeight += w
	}
	return d <= days
}

// Pattern: binary search on minimum capacity.
// Invariant: feasible capacities form a monotonic true suffix.
// Complexity: O(n log sum(weights)) time, O(1) extra space.
// Interview line: if a capacity ships in time, any larger capacity also ships in time.
func ShipWithinDays(weights []int, days int) int {
	maxW := 0
	sumW := 0
	for _, w := range weights {
		if w > maxW {
			maxW = w
		}
		sumW += w
	}

	left := maxW
	right := sumW
	ans := right

	for left <= right {
		mid := left + (right-left)/2
		if feasibleCapacity(weights, mid, days) {
			ans = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
