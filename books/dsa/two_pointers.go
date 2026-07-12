package dsa

import (
	"sort"
)

// Pattern: sorted array + two pointers from both ends.
// Invariant: move left/right based on whether total is too small or too large.
// Complexity: O(n) time, O(1) extra space.
// Interview line: sorted order lets us adjust the sum directionally.
func TwoSumSorted(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1} // 1-indexed helper
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

// Pattern: sort + fixed anchor + two pointers.
// Invariant: move left/right based on whether total is too small or too large.
// Complexity: O(n^2) time, O(1) extra space excluding output.
// Interview line: sorting lets us skip duplicates and adjust the sum directionally.
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return res
}

// Pattern: two pointers maximizing width times limiting height.
// Invariant: moving the shorter side is the only move that can improve area.
// Complexity: O(n) time, O(1) extra space.
// Interview line: width only shrinks, so we discard the shorter wall.
func MaxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0

	for left < right {
		h := height[left]
		if height[right] < h {
			h = height[right]
		}
		area := h * (right - left)
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// Pattern: two pointers with left/right max boundaries.
// Invariant: water on the lower side is decided by that side's max.
// Complexity: O(n) time, O(1) extra space.
// Interview line: process the smaller boundary because the opposite side already bounds it.
func TrapRainWater(height []int) int {
	if len(height) == 0 {
		return 0
	}
	left := 0
	right := len(height) - 1
	leftMax := height[left]
	rightMax := height[right]
	water := 0

	for left < right {
		if leftMax < rightMax {
			left++
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
		} else {
			right--
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
		}
	}
	return water
}

// Pattern: read/write pointers on a sorted array.
// Invariant: nums[:write] contains the unique prefix.
// Complexity: O(n) time, O(1) extra space.
// Interview line: read scans everything, write only advances when a new value appears.
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	write := 1
	for read := 1; read < len(nums); read++ {
		if nums[read] != nums[read-1] {
			nums[write] = nums[read]
			write++
		}
	}
	return write
}

// Pattern: stable partition with read/write pointers.
// Invariant: nums[:write] contains all non-zero values seen so far.
// Complexity: O(n) time, O(1) extra space.
// Interview line: swap each non-zero into the next write slot while preserving order.
func MoveZeroes(nums []int) {
	write := 0
	for read := 0; read < len(nums); read++ {
		if nums[read] != 0 {
			nums[write], nums[read] = nums[read], nums[write]
			write++
		}
	}
}

// Pattern: two pointers over two strings.
// Invariant: i is the next character of s we still need to match.
// Complexity: O(len(t)) time, O(1) extra space.
// Interview line: scan the larger string once and advance the subsequence pointer on matches.
func IsSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == len(s)
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}
	return c
}

// Pattern: inward two pointers with filtering.
// Invariant: compare only alphanumeric characters after normalization.
// Complexity: O(n) time, O(1) extra space.
// Interview line: skip irrelevant characters, then compare the normalized ends.
func ValidPalindrome(s string) bool {
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !isAlphanumeric(s[left]) {
			left++
		}
		for left < right && !isAlphanumeric(s[right]) {
			right--
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

// Pattern: Dutch national flag with three regions.
// Invariant: left side is 0s, middle is 1s, right side is 2s.
// Complexity: O(n) time, O(1) extra space.
// Interview line: one pass partitions the array because there are only three values.
func SortColors(nums []int) {
	left := 0
	i := 0
	right := len(nums) - 1
	for i <= right {
		if nums[i] == 0 {
			nums[left], nums[i] = nums[i], nums[left]
			left++
			i++
		} else if nums[i] == 2 {
			nums[right], nums[i] = nums[i], nums[right]
			right--
		} else {
			i++
		}
	}
}
