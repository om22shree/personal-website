package dsa

import (
	"sort"
	"strconv"
)

// Pattern: stack of expected closers.
// Invariant: stack top is the next closing bracket required.
// Complexity: O(n) time, O(n) space.
// Interview line: push what would make the current opener valid.
func IsValidParentheses(s string) bool {
	var stack []rune
	pairs := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	for _, r := range s {
		if closer, ok := pairs[r]; ok {
			stack = append(stack, closer)
		} else {
			if len(stack) == 0 || stack[len(stack)-1] != r {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// Pattern: stack evaluation of postfix expressions.
// Invariant: stack stores evaluated operands waiting for an operator.
// Complexity: O(n) time, O(n) space.
// Interview line: postfix operators consume the two most recent values.
func EvalRPN(tokens []string) int {
	var stack []int
	for _, tok := range tokens {
		switch tok {
		case "+", "-", "*", "/":
			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			var val int
			if tok == "+" {
				val = a + b
			} else if tok == "-" {
				val = a - b
			} else if tok == "*" {
				val = a * b
			} else {
				val = a / b
			}
			stack = append(stack, val)
		default:
			val, _ := strconv.Atoi(tok)
			stack = append(stack, val)
		}
	}
	return stack[0]
}

// Pattern: backtracking with balance counters.
// Invariant: close_count never exceeds open_count, and open_count never exceeds n.
// Complexity: O(Catalan(n)) time, O(n) recursion space.
// Interview line: only add a parenthesis if it can still lead to a valid string.
func GenerateParenthesis(n int) []string {
	var res []string
	var backtrack func(s string, open int, close int)
	backtrack = func(s string, open int, close int) {
		if len(s) == 2*n {
			res = append(res, s)
			return
		}
		if open < n {
			backtrack(s+"(", open+1, close)
		}
		if close < open {
			backtrack(s+")", open, close+1)
		}
	}
	backtrack("", 0, 0)
	return res
}

// Pattern: monotonic decreasing stack of indexes.
// Invariant: temperatures at stack indexes wait for a warmer future day.
// Complexity: O(n) time, O(n) space.
// Interview line: when today is warmer, it resolves colder days on the stack.
func DailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	res := make([]int, n)
	var stack []int // stores indexes

	for i, temp := range temperatures {
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			idx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[idx] = i - idx
		}
		stack = append(stack, i)
	}
	return res
}

type car struct {
	pos  int
	time float64
}

// Pattern: sorted positions plus decreasing arrival times.
// Invariant: each fleet on stack has an arrival time not less than fleets behind it.
// Complexity: O(n log n) time, O(n) space.
// Interview line: a faster car behind merges if it reaches no later than the fleet ahead.
func CarFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}
	cars := make([]car, n)
	for i := 0; i < n; i++ {
		cars[i] = car{
			pos:  position[i],
			time: float64(target-position[i]) / float64(speed[i]),
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	var lastTime float64
	for _, c := range cars {
		if c.time > lastTime {
			fleets++
			lastTime = c.time
		}
	}
	return fleets
}

// Pattern: monotonic increasing stack of bar indexes.
// Invariant: stack heights are increasing, waiting for a shorter right boundary.
// Complexity: O(n) time, O(n) space.
// Interview line: a popped bar's rectangle is bounded by the previous smaller bar and current index.
func LargestRectangleArea(heights []int) int {
	var stack []int // stores indexes
	maxArea := 0
	// Append a sentinel 0 to resolve all remaining bars
	hCopy := append(heights, 0)

	for i, h := range hCopy {
		for len(stack) > 0 && hCopy[stack[len(stack)-1]] > h {
			hIdx := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			height := hCopy[hIdx]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}
