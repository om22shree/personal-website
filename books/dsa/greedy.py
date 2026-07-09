from collections import Counter, deque
from typing import List


# Pattern: one-pass minimum price.
# Invariant: min_price is the lowest price before the current day.
# Complexity: O(n) time, O(1) space.
# Interview line: sell today against the cheapest earlier buy.
def max_profit(prices: List[int]) -> int:
    min_price = float("inf")
    best = 0
    for price in prices:
        min_price = min(min_price, price)
        best = max(best, price - min_price)
    return best


# Pattern: collect every upward edge.
# Invariant: every positive day-to-day gain can be taken independently.
# Complexity: O(n) time, O(1) space.
# Interview line: unlimited transactions reduce to summing all positive slopes.
def max_profit_many(prices: List[int]) -> int:
    return sum(max(prices[i] - prices[i - 1], 0) for i in range(1, len(prices)))


# Pattern: greedy farthest reach.
# Invariant: reach is the farthest index reachable from processed positions.
# Complexity: O(n) time, O(1) space.
# Interview line: if you ever stand beyond reach, the end is impossible.
def can_jump(nums: List[int]) -> bool:
    reach = 0
    for i, jump in enumerate(nums):
        if i > reach:
            return False
        reach = max(reach, i + jump)
    return True


# Pattern: BFS levels without queue.
# Invariant: current_end bounds indexes reachable with jumps jumps.
# Complexity: O(n) time, O(1) space.
# Interview line: when you exhaust the current range, commit one jump to the farthest next range.
def jump_game_ii(nums: List[int]) -> int:
    jumps = 0
    current_end = 0
    farthest = 0
    for i in range(len(nums) - 1):
        farthest = max(farthest, i + nums[i])
        if i == current_end:
            jumps += 1
            current_end = farthest
    return jumps


# Pattern: gas tank surplus reset.
# Invariant: if tank drops below zero, no station in this failed segment can start.
# Complexity: O(n) time, O(1) space.
# Interview line: total surplus proves existence; local deficit chooses the next start.
def can_complete_circuit(gas: List[int], cost: List[int]) -> int:
    if sum(gas) < sum(cost):
        return -1
    start = 0
    tank = 0
    for i, (g, c) in enumerate(zip(gas, cost)):
        tank += g - c
        if tank < 0:
            start = i + 1
            tank = 0
    return start


# Pattern: last occurrence partitioning.
# Invariant: current partition must extend to the farthest last occurrence of its letters.
# Complexity: O(n) time, O(1) space.
# Interview line: close a partition only when every character inside ends inside.
def partition_labels(s: str) -> List[int]:
    last = {ch: i for i, ch in enumerate(s)}
    ans = []
    start = 0
    end = 0
    for i, ch in enumerate(s):
        end = max(end, last[ch])
        if i == end:
            ans.append(end - start + 1)
            start = i + 1
    return ans


# Pattern: greedy range of possible open counts.
# Invariant: low/high bound how many unmatched opens are possible.
# Complexity: O(n) time, O(1) space.
# Interview line: treat star as whichever choice keeps the open-count range viable.
def check_valid_string(s: str) -> bool:
    low = high = 0
    for ch in s:
        if ch == "(":
            low += 1
            high += 1
        elif ch == ")":
            low = max(low - 1, 0)
            high -= 1
        else:
            low = max(low - 1, 0)
            high += 1
        if high < 0:
            return False
    return low == 0


# Pattern: greedy target coverage.
# Invariant: matched[i] is true once some triplet supplies target[i].
# Complexity: O(n) time, O(1) space.
# Interview line: discard triplets that overshoot target; useful ones can be merged coordinate-wise.
def merge_triplets(triplets: List[List[int]], target: List[int]) -> bool:
    matched = [False, False, False]
    for triplet in triplets:
        if any(triplet[i] > target[i] for i in range(3)):
            continue
        for i in range(3):
            if triplet[i] == target[i]:
                matched[i] = True
    return all(matched)


# Pattern: sorted counting with queue of group starts.
# Invariant: opened groups wait for the next consecutive card.
# Complexity: O(n log n) time, O(n) space.
# Interview line: always satisfy the smallest card first.
def hand_of_straights(hand: List[int], group_size: int) -> bool:
    if len(hand) % group_size:
        return False
    count = Counter(hand)
    starts = deque()
    opened = 0
    previous = None

    for card in sorted(count):
        if opened and card != previous + 1:
            return False
        starts.append(count[card] - opened)
        if starts[-1] < 0:
            return False
        opened = count[card]
        if len(starts) == group_size:
            opened -= starts.popleft()
        previous = card

    return opened == 0
