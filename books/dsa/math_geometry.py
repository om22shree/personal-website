from typing import List


# Pattern: cycle detection by digit-square transform.
# Invariant: repeated values mean the process entered a cycle.
# Complexity: O(log n) per transform, O(1) eventual state space.
# Interview line: happy numbers either hit 1 or loop forever.
def happy_number(n: int) -> bool:
    seen = set()
    while n != 1 and n not in seen:
        seen.add(n)
        total = 0
        while n:
            n, digit = divmod(n, 10)
            total += digit * digit
        n = total
    return n == 1


# Pattern: exponentiation by squaring.
# Invariant: ans times x^power equals the original requested power.
# Complexity: O(log n) time, O(1) space.
# Interview line: square the base whenever you halve the exponent.
def pow_x_n(x: float, n: int) -> float:
    if n < 0:
        x = 1 / x
        n = -n
    ans = 1.0
    while n:
        if n & 1:
            ans *= x
        x *= x
        n >>= 1
    return ans


# Pattern: transpose then reverse rows.
# Invariant: after transpose and row reversal, each element moves to rotated position.
# Complexity: O(n^2) time, O(1) extra space.
# Interview line: 90-degree rotation is transpose plus horizontal reflection.
def rotate_matrix(matrix: List[List[int]]) -> None:
    n = len(matrix)
    for r in range(n):
        for c in range(r + 1, n):
            matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
    for row in matrix:
        row.reverse()


# Pattern: shrinking rectangle traversal.
# Invariant: top/bottom/left/right bound the remaining unvisited cells.
# Complexity: O(mn) time, O(1) extra space excluding output.
# Interview line: peel one border at a time and shrink the rectangle.
def spiral_order(matrix: List[List[int]]) -> List[int]:
    if not matrix:
        return []
    top, bottom = 0, len(matrix) - 1
    left, right = 0, len(matrix[0]) - 1
    ans = []
    while top <= bottom and left <= right:
        for c in range(left, right + 1):
            ans.append(matrix[top][c])
        top += 1
        for r in range(top, bottom + 1):
            ans.append(matrix[r][right])
        right -= 1
        if top <= bottom:
            for c in range(right, left - 1, -1):
                ans.append(matrix[bottom][c])
            bottom -= 1
        if left <= right:
            for r in range(bottom, top - 1, -1):
                ans.append(matrix[r][left])
            left += 1
    return ans


# Pattern: first-row/first-column marker matrix.
# Invariant: markers record which rows and columns must become zero.
# Complexity: O(mn) time, O(1) extra space.
# Interview line: reuse the matrix border as marker storage.
def set_zeroes(matrix: List[List[int]]) -> None:
    rows = len(matrix)
    cols = len(matrix[0]) if rows else 0
    first_col_zero = any(matrix[r][0] == 0 for r in range(rows))
    first_row_zero = any(matrix[0][c] == 0 for c in range(cols)) if rows else False

    for r in range(1, rows):
        for c in range(1, cols):
            if matrix[r][c] == 0:
                matrix[r][0] = 0
                matrix[0][c] = 0

    for r in range(1, rows):
        for c in range(1, cols):
            if matrix[r][0] == 0 or matrix[0][c] == 0:
                matrix[r][c] = 0

    if first_row_zero:
        for c in range(cols):
            matrix[0][c] = 0
    if first_col_zero:
        for r in range(rows):
            matrix[r][0] = 0


# Pattern: grade-school carry from the right.
# Invariant: suffix to the right of i has already been normalized.
# Complexity: O(n) time, O(1) extra space.
# Interview line: add one at the end, carry left only while digits become ten.
def plus_one(digits: List[int]) -> List[int]:
    for i in range(len(digits) - 1, -1, -1):
        if digits[i] < 9:
            digits[i] += 1
            return digits
        digits[i] = 0
    return [1] + digits


# Pattern: digit multiplication with positional buckets.
# Invariant: ans[i + j + 1] holds the ones place for digits i and j.
# Complexity: O(mn) time, O(m + n) space.
# Interview line: multiply each digit pair and push carry left.
def multiply_strings(num1: str, num2: str) -> str:
    if num1 == "0" or num2 == "0":
        return "0"
    ans = [0] * (len(num1) + len(num2))
    for i in range(len(num1) - 1, -1, -1):
        for j in range(len(num2) - 1, -1, -1):
            product = int(num1[i]) * int(num2[j]) + ans[i + j + 1]
            ans[i + j + 1] = product % 10
            ans[i + j] += product // 10
    return "".join(map(str, ans)).lstrip("0")
