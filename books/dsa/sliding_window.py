from collections import deque, defaultdict, Counter
from typing import Dict, List


# Pattern: fixed-size sliding window with rolling sum.
# Invariant: window always contains exactly k elements after the first k values.
# Complexity: O(n) time, O(1) extra space.
# Interview line: subtract the outgoing element and add the incoming one instead of recomputing.
def max_average_subarray(nums: List[int], k: int) -> float:
    window_sum = sum(nums[:k])
    best = window_sum

    for right in range(k, len(nums)):
        window_sum += nums[right] - nums[right - k]
        best = max(best, window_sum)

    return best / k


# Pattern: variable sliding window with last-seen index.
# Invariant: current window has no duplicate characters.
# Complexity: O(n) time, O(min(n, charset)) space.
# Interview line: when a duplicate appears inside the window, jump left past its last position.
def length_of_longest_substring(s: str) -> int:
    seen: Dict[str, int] = {}
    left = 0
    best = 0

    for right, ch in enumerate(s):
        if ch in seen and seen[ch] >= left:
            left = seen[ch] + 1
        seen[ch] = right
        best = max(best, right - left + 1)

    return best


# Pattern: variable sliding window with required character counts.
# Invariant: shrink only while the window satisfies every needed count.
# Complexity: O(len(s) + len(t)) time, O(len(t)) space.
# Interview line: expand to become valid, then shrink greedily to find the smallest valid window.
def min_window_substring(s: str, t: str) -> str:
    if not t:
        return ""

    need = Counter(t)
    window = defaultdict(int)
    have = 0
    required = len(need)
    left = 0
    best = (float("inf"), 0, 0)

    for right, ch in enumerate(s):
        window[ch] += 1
        if ch in need and window[ch] == need[ch]:
            have += 1

        while have == required:
            if right - left + 1 < best[0]:
                best = (right - left + 1, left, right)

            drop = s[left]
            window[drop] -= 1
            if drop in need and window[drop] < need[drop]:
                have -= 1
            left += 1

    return "" if best[0] == float("inf") else s[best[1] : best[2] + 1]


# Pattern: monotonic deque for fixed-size window maximum.
# Invariant: deque stores candidate indices in decreasing value order.
# Complexity: O(n) time, O(k) space.
# Interview line: smaller elements behind a larger incoming value can never become the max.
def sliding_window_maximum(nums: List[int], k: int) -> List[int]:
    q = deque()
    result = []

    for right, value in enumerate(nums):
        while q and nums[q[-1]] <= value:
            q.pop()
        q.append(right)

        if q[0] <= right - k:
            q.popleft()
        if right >= k - 1:
            result.append(nums[q[0]])

    return result


# Pattern: fixed-size sliding window with frequency maps.
# Invariant: window length never exceeds len(p), and matching counts mark an anagram.
# Complexity: O(n) time with bounded alphabet, O(len(p)) space.
# Interview line: keep the window exactly the pattern length, then compare character counts.
def find_anagrams(s: str, p: str) -> List[int]:
    need = Counter(p)
    window = Counter()
    result = []
    left = 0

    for right, ch in enumerate(s):
        window[ch] += 1

        if right - left + 1 > len(p):
            drop = s[left]
            window[drop] -= 1
            if window[drop] == 0:
                del window[drop]
            left += 1

        if window == need:
            result.append(left)

    return result


# Pattern: variable sliding window with most frequent character tracking.
# Invariant: window is valid when replacements needed <= k.
# Complexity: O(n) time, O(1) space for uppercase English letters.
# Interview line: the best target character is the most frequent one already in the window.
def longest_repeating_character_replacement(s: str, k: int) -> int:
    counts = defaultdict(int)
    left = 0
    max_freq = 0
    best = 0

    for right, ch in enumerate(s):
        counts[ch] += 1
        max_freq = max(max_freq, counts[ch])

        while right - left + 1 - max_freq > k:
            counts[s[left]] -= 1
            left += 1

        best = max(best, right - left + 1)

    return best


if __name__ == "__main__":
    assert length_of_longest_substring("abcabcbb") == 3
    assert min_window_substring("ADOBECODEBANC", "ABC") == "BANC"
    assert sliding_window_maximum([1, 3, -1, -3, 5, 3, 6, 7], 3) == [3, 3, 5, 5, 6, 7]
