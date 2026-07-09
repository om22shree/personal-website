from typing import List


# Pattern: XOR cancellation.
# Invariant: equal values cancel to zero, leaving the unpaired value.
# Complexity: O(n) time, O(1) space.
# Interview line: XOR is perfect when every duplicate appears exactly twice.
def single_number(nums: List[int]) -> int:
    ans = 0
    for num in nums:
        ans ^= num
    return ans


# Pattern: low-bit deletion.
# Invariant: each loop removes the lowest set bit from n.
# Complexity: O(number of set bits) time, O(1) space.
# Interview line: n & (n - 1) drops the rightmost one bit.
def hamming_weight(n: int) -> int:
    count = 0
    while n:
        n &= n - 1
        count += 1
    return count


# Pattern: DP over cleared lowest set bit.
# Invariant: bits[i] = bits[i without lowest one bit] + 1.
# Complexity: O(n) time, O(n) space.
# Interview line: reuse the count for a smaller number after removing one bit.
def count_bits(n: int) -> List[int]:
    bits = [0] * (n + 1)
    for i in range(1, n + 1):
        bits[i] = bits[i & (i - 1)] + 1
    return bits


# Pattern: fixed-width bit reversal.
# Invariant: ans receives bits from n from low to high.
# Complexity: O(1) time for 32 bits, O(1) space.
# Interview line: shift answer left, append the current lowest bit, then shift input right.
def reverse_bits(n: int) -> int:
    ans = 0
    for _ in range(32):
        ans = (ans << 1) | (n & 1)
        n >>= 1
    return ans


# Pattern: XOR complete range with array values.
# Invariant: indexes and present values cancel, leaving the missing value.
# Complexity: O(n) time, O(1) space.
# Interview line: XOR every index and value when exactly one number is missing.
def missing_number(nums: List[int]) -> int:
    ans = len(nums)
    for i, num in enumerate(nums):
        ans ^= i ^ num
    return ans


# Pattern: bitwise addition with carry.
# Invariant: a is partial sum without carries, b is pending carry.
# Complexity: O(1) time for fixed-width integers, O(1) space.
# Interview line: XOR adds bits, AND carries bits.
def get_sum(a: int, b: int) -> int:
    mask = 0xFFFFFFFF
    max_int = 0x7FFFFFFF
    while b:
        a, b = (a ^ b) & mask, ((a & b) << 1) & mask
    return a if a <= max_int else ~(a ^ mask)


# Pattern: common prefix of binary range.
# Invariant: right shifts remove changing suffix bits until left equals right.
# Complexity: O(1) time for fixed-width integers, O(1) space.
# Interview line: bitwise AND across a range preserves only the shared binary prefix.
def range_bitwise_and(left: int, right: int) -> int:
    shifts = 0
    while left < right:
        left >>= 1
        right >>= 1
        shifts += 1
    return left << shifts
