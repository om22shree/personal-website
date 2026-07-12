package dsa

// Pattern: XOR cancellation.
// Invariant: equal values cancel to zero, leaving the unpaired value.
// Complexity: O(n) time, O(1) space.
// Interview line: XOR is perfect when every duplicate appears exactly twice.
func SingleNumber(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

// Pattern: low-bit deletion.
// Invariant: each loop removes the lowest set bit from n.
// Complexity: O(number of set bits) time, O(1) space.
// Interview line: n & (n - 1) drops the rightmost one bit.
func HammingWeight(n uint32) int {
	count := 0
	for n > 0 {
		n &= (n - 1)
		count++
	}
	return count
}

// Pattern: DP over cleared lowest set bit.
// Invariant: bits[i] = bits[i without lowest one bit] + 1.
// Complexity: O(n) time, O(n) space.
// Interview line: reuse the count for a smaller number after removing one bit.
func CountBits(n int) []int {
	bits := make([]int, n+1)
	for i := 1; i <= n; i++ {
		bits[i] = bits[i&(i-1)] + 1
	}
	return bits
}

// Pattern: fixed-width bit reversal.
// Invariant: ans receives bits from n from low to high.
// Complexity: O(1) time for 32 bits, O(1) space.
// Interview line: shift answer left, append the current lowest bit, then shift input right.
func ReverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for i := 0; i < 32; i++ {
		ans = (ans << 1) | (num & 1)
		num >>= 1
	}
	return ans
}

// Pattern: XOR complete range with array values.
// Invariant: indexes and present values cancel, leaving the missing value.
// Complexity: O(n) time, O(1) space.
// Interview line: XOR every index and value when exactly one number is missing.
func MissingNumber(nums []int) int {
	res := len(nums)
	for i, num := range nums {
		res ^= i ^ num
	}
	return res
}

// Pattern: bitwise addition with carry.
// Invariant: a is partial sum without carries, b is pending carry.
// Complexity: O(1) time for fixed-width integers, O(1) space.
// Interview line: XOR adds bits, AND carries bits.
func GetSum(a int, b int) int {
	for b != 0 {
		carry := (a & b) << 1
		a = a ^ b
		b = carry
	}
	return a
}

// Pattern: common prefix of binary range.
// Invariant: right shifts remove changing suffix bits until left equals right.
// Complexity: O(1) time for fixed-width integers, O(1) space.
// Interview line: bitwise AND across a range preserves only the shared binary prefix.
func RangeBitwiseAnd(left int, right int) int {
	shifts := 0
	for left < right {
		left >>= 1
		right >>= 1
		shifts++
	}
	return left << shifts
}
