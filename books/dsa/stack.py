from typing import List


# Pattern: stack of expected closers.
# Invariant: stack top is the next closing bracket required.
# Complexity: O(n) time, O(n) space.
# Interview line: push what would make the current opener valid.
def valid_parentheses(s: str) -> bool:
    pairs = {"(": ")", "[": "]", "{": "}"}
    stack = []
    for ch in s:
        if ch in pairs:
            stack.append(pairs[ch])
        elif not stack or stack.pop() != ch:
            return False
    return not stack


class MinStack:
    # Pattern: value stack plus running-min stack.
    # Invariant: mins[-1] is the minimum of all values currently in stack.
    # Complexity: O(1) setup, O(n) space.
    # Interview line: duplicate the current minimum beside the normal stack.
    def __init__(self):
        self.stack = []
        self.mins = []

    # Pattern: push with current minimum snapshot.
    # Invariant: both stacks have the same length after every push.
    # Complexity: O(1) time, O(1) space.
    # Interview line: the min stack remembers history so pop can restore old minima.
    def push(self, val: int) -> None:
        self.stack.append(val)
        self.mins.append(val if not self.mins else min(val, self.mins[-1]))

    # Pattern: synchronized pop.
    # Invariant: removing a value removes its matching minimum snapshot.
    # Complexity: O(1) time, O(1) space.
    # Interview line: pop both stacks together.
    def pop(self) -> None:
        self.stack.pop()
        self.mins.pop()

    # Pattern: read stack top.
    # Invariant: stack[-1] is the most recent unpopped value.
    # Complexity: O(1) time, O(1) space.
    # Interview line: normal stack top still gives the latest value.
    def top(self) -> int:
        return self.stack[-1]

    # Pattern: read cached minimum.
    # Invariant: mins[-1] tracks the current stack minimum.
    # Complexity: O(1) time, O(1) space.
    # Interview line: the minimum is maintained during writes, not recomputed during reads.
    def get_min(self) -> int:
        return self.mins[-1]


# Pattern: stack evaluation of postfix expressions.
# Invariant: stack stores evaluated operands waiting for an operator.
# Complexity: O(n) time, O(n) space.
# Interview line: postfix operators consume the two most recent values.
def eval_rpn(tokens: List[str]) -> int:
    stack = []
    for token in tokens:
        if token not in {"+", "-", "*", "/"}:
            stack.append(int(token))
            continue
        b = stack.pop()
        a = stack.pop()
        if token == "+":
            stack.append(a + b)
        elif token == "-":
            stack.append(a - b)
        elif token == "*":
            stack.append(a * b)
        else:
            stack.append(int(a / b))
    return stack[-1]


# Pattern: backtracking with balance counters.
# Invariant: close_count never exceeds open_count, and open_count never exceeds n.
# Complexity: O(Catalan(n)) time, O(n) recursion space.
# Interview line: only add a parenthesis if it can still lead to a valid string.
def generate_parentheses(n: int) -> List[str]:
    ans = []

    def backtrack(path: List[str], opens: int, closes: int) -> None:
        if len(path) == 2 * n:
            ans.append("".join(path))
            return
        if opens < n:
            path.append("(")
            backtrack(path, opens + 1, closes)
            path.pop()
        if closes < opens:
            path.append(")")
            backtrack(path, opens, closes + 1)
            path.pop()

    backtrack([], 0, 0)
    return ans


# Pattern: monotonic decreasing stack of indexes.
# Invariant: temperatures at stack indexes wait for a warmer future day.
# Complexity: O(n) time, O(n) space.
# Interview line: when today is warmer, it resolves colder days on the stack.
def daily_temperatures(temperatures: List[int]) -> List[int]:
    ans = [0] * len(temperatures)
    stack = []
    for i, temp in enumerate(temperatures):
        while stack and temperatures[stack[-1]] < temp:
            j = stack.pop()
            ans[j] = i - j
        stack.append(i)
    return ans


# Pattern: sorted positions plus decreasing arrival times.
# Invariant: each fleet on stack has an arrival time not less than fleets behind it.
# Complexity: O(n log n) time, O(n) space.
# Interview line: a faster car behind merges if it reaches no later than the fleet ahead.
def car_fleet(target: int, position: List[int], speed: List[int]) -> int:
    fleets = []
    cars = sorted(zip(position, speed), reverse=True)
    for pos, spd in cars:
        time = (target - pos) / spd
        if not fleets or time > fleets[-1]:
            fleets.append(time)
    return len(fleets)


# Pattern: monotonic increasing stack of bar indexes.
# Invariant: stack heights are increasing, waiting for a shorter right boundary.
# Complexity: O(n) time, O(n) space.
# Interview line: a popped bar's rectangle is bounded by the previous smaller bar and current index.
def largest_rectangle_area(heights: List[int]) -> int:
    stack = []
    best = 0
    for i, height in enumerate(heights + [0]):
        while stack and heights[stack[-1]] > height:
            h = heights[stack.pop()]
            left = stack[-1] if stack else -1
            best = max(best, h * (i - left - 1))
        stack.append(i)
    return best
