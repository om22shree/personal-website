package dsa

// Pattern: cycle detection by digit-square transform.
// Invariant: repeated values mean the process entered a cycle.
// Complexity: O(log n) per transform, O(1) eventual state space.
// Interview line: happy numbers either hit 1 or loop forever.
func IsHappy(n int) bool {
	sumSquares := func(num int) int {
		sum := 0
		for num > 0 {
			digit := num % 10
			sum += digit * digit
			num /= 10
		}
		return sum
	}

	slow := n
	fast := sumSquares(n)

	for fast != 1 && slow != fast {
		slow = sumSquares(slow)
		fast = sumSquares(sumSquares(fast))
	}
	return fast == 1
}

// Pattern: exponentiation by squaring.
// Invariant: ans times x^power equals the original requested power.
// Complexity: O(log n) time, O(1) space.
// Interview line: square the base whenever you halve the exponent.
func MyPow(x float64, n int) float64 {
	p := n
	if p < 0 {
		x = 1.0 / x
		p = -p
	}

	ans := 1.0
	currentProduct := x

	for p > 0 {
		if p%2 == 1 {
			ans *= currentProduct
		}
		currentProduct *= currentProduct
		p /= 2
	}
	return ans
}

// Pattern: transpose then reverse rows.
// Invariant: after transpose and row reversal, each element moves to rotated position.
// Complexity: O(n^2) time, O(1) extra space.
// Interview line: 90-degree rotation is transpose plus horizontal reflection.
func Rotate(matrix [][]int) {
	n := len(matrix)
	if n == 0 {
		return
	}

	// Step 1: Transpose
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// Step 2: Reverse each row
	for i := 0; i < n; i++ {
		left := 0
		right := n - 1
		for left < right {
			matrix[i][left], matrix[i][right] = matrix[i][right], matrix[i][left]
			left++
			right--
		}
	}
}

// Pattern: shrinking rectangle traversal.
// Invariant: top/bottom/left/right bound the remaining unvisited cells.
// Complexity: O(mn) time, O(1) extra space excluding output.
// Interview line: peel one border at a time and shrink the rectangle.
func SpiralOrder(matrix [][]int) []int {
	var res []int
	if len(matrix) == 0 {
		return res
	}

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for top <= bottom && left <= right {
		// Traverse right
		for c := left; c <= right; c++ {
			res = append(res, matrix[top][c])
		}
		top++

		// Traverse down
		for r := top; r <= bottom; r++ {
			res = append(res, matrix[r][right])
		}
		right--

		if top <= bottom {
			// Traverse left
			for c := right; c >= left; c-- {
				res = append(res, matrix[bottom][c])
			}
			bottom--
		}

		if left <= right {
			// Traverse up
			for r := bottom; r >= top; r-- {
				res = append(res, matrix[r][left])
			}
			left++
		}
	}
	return res
}

// Pattern: first-row/first-column marker matrix.
// Invariant: markers record which rows and columns must become zero.
// Complexity: O(mn) time, O(1) extra space.
// Interview line: reuse the matrix border as marker storage.
func SetZeroes(matrix [][]int) {
	rows := len(matrix)
	if rows == 0 {
		return
	}
	cols := len(matrix[0])

	rowZero := false

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				matrix[0][c] = 0
				if r > 0 {
					matrix[r][0] = 0
				} else {
					rowZero = true
				}
			}
		}
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			if matrix[r][0] == 0 || matrix[0][c] == 0 {
				matrix[r][c] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for r := 0; r < rows; r++ {
			matrix[r][0] = 0
		}
	}

	if rowZero {
		for c := 0; c < cols; c++ {
			matrix[0][c] = 0
		}
	}
}

// Pattern: grade-school carry from the right.
// Invariant: suffix to the right of i has already been normalized.
// Complexity: O(n) time, O(1) extra space.
// Interview line: add one at the end, carry left only while digits become ten.
func PlusOne(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	// If we are here, all digits were 9
	newDigits := make([]int, n+1)
	newDigits[0] = 1
	return newDigits
}

// Pattern: digit multiplication with positional buckets.
// Invariant: ans[i + j + 1] holds the ones place for digits i and j.
// Complexity: O(mn) time, O(m + n) space.
// Interview line: multiply each digit pair and push carry left.
func MultiplyStrings(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	m := len(num1)
	n := len(num2)
	pos := make([]int, m+n)

	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1 := i + j
			p2 := i + j + 1
			sum := mul + pos[p2]

			pos[p2] = sum % 10
			pos[p1] += sum / 10
		}
	}

	var res []byte
	for _, p := range pos {
		if !(len(res) == 0 && p == 0) {
			res = append(res, byte(p)+'0')
		}
	}
	return string(res)
}
