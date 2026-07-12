package dsa

import (
	"sort"
	"strconv"
	"strings"
)

// Pattern: hashmap complement lookup.
// Invariant: seen maps every previous value to its index.
// Complexity: O(n) time, O(n) space.
// Interview line: for pair sums, store what you have seen and ask whether the complement exists.
func TwoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		need := target - num
		if idx, ok := seen[need]; ok {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

// Pattern: set membership duplicate check.
// Invariant: seen contains exactly the distinct values processed so far.
// Complexity: O(n) time, O(n) space.
// Interview line: a duplicate is the first value already present in the set.
func ContainsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := seen[num]; ok {
			return true
		}
		seen[num] = struct{}{}
	}
	return false
}

// Pattern: frequency equality.
// Invariant: anagrams have identical character counts.
// Complexity: O(n) time, O(1) space for fixed alphabet.
// Interview line: order does not matter, counts do.
func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}
	for _, char := range t {
		counts[char]--
		if counts[char] < 0 {
			return false
		}
	}
	return true
}

// Helper to sort a string
func sortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// Pattern: canonical sorted key grouping.
// Invariant: words with the same sorted letters share one bucket.
// Complexity: O(n k log k) time, O(n k) space.
// Interview line: choose a key that erases the irrelevant order.
func GroupAnagrams(strs []string) [][]string {
	groups := make(map[string][]string)
	for _, word := range strs {
		key := sortString(word)
		groups[key] = append(groups[key], word)
	}
	res := make([][]string, 0, len(groups))
	for _, group := range groups {
		res = append(res, group)
	}
	return res
}

// Pattern: prefix product plus suffix product.
// Invariant: answer[i] accumulates product of all values left and right of i.
// Complexity: O(n) time, O(1) extra space excluding output.
// Interview line: multiply prefix on the way forward, suffix on the way back.
func ProductExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := range ans {
		ans[i] = 1
	}

	prefix := 1
	for i, num := range nums {
		ans[i] = prefix
		prefix *= num
	}

	suffix := 1
	for i := n - 1; i >= 0; i-- {
		ans[i] *= suffix
		suffix *= nums[i]
	}

	return ans
}

// Pattern: length-prefixed string encoding.
// Invariant: each encoded string is length, delimiter, then exact payload.
// Complexity: O(total chars) time, O(total chars) space.
// Interview line: length prefixes avoid delimiter ambiguity inside strings.
func EncodeStrings(strs []string) string {
	var sb strings.Builder
	for _, s := range strs {
		sb.WriteString(strconv.Itoa(len(s)))
		sb.WriteByte('#')
		sb.WriteString(s)
	}
	return sb.String()
}

// Pattern: parse length-prefixed strings.
// Invariant: i always points to the start of the next length field.
// Complexity: O(total chars) time, O(total chars) space.
// Interview line: read the length first, then jump exactly that many characters.
func DecodeStrings(data string) []string {
	var ans []string
	i := 0
	for i < len(data) {
		j := strings.Index(data[i:], "#")
		if j == -1 {
			break
		}
		j = i + j
		length, _ := strconv.Atoi(data[i:j])
		start := j + 1
		ans = append(ans, data[start:start+length])
		i = start + length
	}
	return ans
}

// Pattern: row/column/box set validation.
// Invariant: no non-dot value repeats in its row, column, or 3x3 box.
// Complexity: O(1) time, O(1) space for fixed 9x9 board.
// Interview line: each filled cell belongs to three uniqueness constraints.
func ValidSudoku(board [][]byte) bool {
	var rows [9]map[byte]bool
	var cols [9]map[byte]bool
	var boxes [9]map[byte]bool

	for i := 0; i < 9; i++ {
		rows[i] = make(map[byte]bool)
		cols[i] = make(map[byte]bool)
		boxes[i] = make(map[byte]bool)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			val := board[r][c]
			if val == '.' {
				continue
			}
			box := (r/3)*3 + c/3
			if rows[r][val] || cols[c][val] || boxes[box][val] {
				return false
			}
			rows[r][val] = true
			cols[c][val] = true
			boxes[box][val] = true
		}
	}
	return true
}

// Pattern: hash-set sequence starts.
// Invariant: only numbers without a predecessor start counted sequences.
// Complexity: O(n) time, O(n) space.
// Interview line: count each consecutive run once, from its left edge.
func LongestConsecutive(nums []int) int {
	values := make(map[int]bool)
	for _, num := range nums {
		values[num] = true
	}
	best := 0
	for num := range values {
		if values[num-1] {
			continue
		}
		cur := num
		for values[cur] {
			cur++
		}
		if cur-num > best {
			best = cur - num
		}
	}
	return best
}

// Pattern: prefix sum frequency.
// Invariant: counts[prefix] is how many earlier prefixes had that sum.
// Complexity: O(n) time, O(n) space.
// Interview line: subarray sum equals k when current_prefix - old_prefix equals k.
func SubarraySumEqualsK(nums []int, k int) int {
	counts := make(map[int]int)
	counts[0] = 1
	prefix := 0
	ans := 0
	for _, num := range nums {
		prefix += num
		ans += counts[prefix-k]
		counts[prefix]++
	}
	return ans
}

// Pattern: Boyer-Moore majority vote.
// Invariant: candidate survives pair cancellations against other values.
// Complexity: O(n) time, O(1) space.
// Interview line: the true majority cannot be fully canceled by all non-majority values.
func MajorityElement(nums []int) int {
	var candidate int
	votes := 0
	for _, num := range nums {
		if votes == 0 {
			candidate = num
		}
		if num == candidate {
			votes++
		} else {
			votes--
		}
	}
	return candidate
}

// Pattern: 2D prefix sums.
// Invariant: prefix[r][c] stores sum of rectangle [0:r) x [0:c).
// Complexity: O(mn) setup, O(mn) space.
// Interview line: inclusion-exclusion turns rectangle queries into four prefix reads.
type NumMatrix struct {
	prefix [][]int
}

func NewNumMatrix(matrix [][]int) NumMatrix {
	rows := len(matrix)
	if rows == 0 {
		return NumMatrix{prefix: [][]int{}}
	}
	cols := len(matrix[0])
	prefix := make([][]int, rows+1)
	for r := 0; r <= rows; r++ {
		prefix[r] = make([]int, cols+1)
	}

	for r := 0; r < rows; r++ {
		rowSum := 0
		for c := 0; c < cols; c++ {
			rowSum += matrix[r][c]
			prefix[r+1][c+1] = prefix[r][c+1] + rowSum
		}
	}
	return NumMatrix{prefix: prefix}
}

// Pattern: 2D prefix sum query.
// Invariant: bottom/right indexes are converted to exclusive prefix coordinates.
// Complexity: O(1) time, O(1) space.
// Interview line: add big rectangle, subtract strips, add back overlap.
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	p := this.prefix
	return p[row2+1][col2+1] - p[row1][col2+1] - p[row2+1][col1] + p[row1][col1]
}
