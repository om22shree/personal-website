from collections import deque
from typing import List, Optional


class ListNode:
    # Pattern: linked-list node container.
    # Invariant: each node stores a value and pointer to the next node.
    # Complexity: O(1) time, O(1) space.
    # Interview line: all pointer problems reduce to carefully rewiring next links.
    def __init__(self, val: int = 0, next: Optional["ListNode"] = None):
        self.val = val
        self.next = next

    # Pattern: debug representation.
    # Invariant: repr shows this node value without walking the list.
    # Complexity: O(1) time, O(1) space.
    # Interview line: keep debug helpers simple so they cannot accidentally traverse cycles.
    def __repr__(self) -> str:
        return f"ListNode({self.val})"


# Pattern: dummy-head linked-list construction.
# Invariant: tail always points to the last node in the built list.
# Complexity: O(n) time, O(n) space.
# Interview line: a dummy head avoids special-casing the first node.
def build_linked_list(values: List[int]) -> Optional[ListNode]:
    dummy = ListNode()
    tail = dummy
    for value in values:
        tail.next = ListNode(value)
        tail = tail.next
    return dummy.next


# Pattern: linear traversal to array.
# Invariant: values contains every node value already visited in order.
# Complexity: O(n) time, O(n) space.
# Interview line: conversion is handy for tests, but avoid it if the prompt asks O(1) space.
def to_list(head: Optional[ListNode]) -> List[int]:
    values = []
    while head:
        values.append(head.val)
        head = head.next
    return values


# Pattern: iterative pointer reversal.
# Invariant: prev is the reversed prefix, cur is the next node to move.
# Complexity: O(n) time, O(1) extra space.
# Interview line: save next before rewiring cur.next.
def reverse_list(head: Optional[ListNode]) -> Optional[ListNode]:
    prev = None
    cur = head
    while cur:
        nxt = cur.next
        cur.next = prev
        prev = cur
        cur = nxt
    return prev


# Pattern: fast/slow pointer cycle detection.
# Invariant: fast moves twice as quickly, so it meets slow iff a cycle exists.
# Complexity: O(n) time, O(1) extra space.
# Interview line: Floyd's algorithm detects a cycle without extra memory.
def has_cycle(head: Optional[ListNode]) -> bool:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            return True
    return False


# Pattern: Floyd cycle entry detection.
# Invariant: after meeting, moving head and meeting pointer together finds cycle start.
# Complexity: O(n) time, O(1) extra space.
# Interview line: the distance math makes the second phase land at the entry node.
def detect_cycle_start(head: Optional[ListNode]) -> Optional[ListNode]:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        if slow is fast:
            break
    else:
        return None

    slow = head
    while slow is not fast:
        slow = slow.next
        fast = fast.next
    return slow


# Pattern: fast/slow pointer midpoint.
# Invariant: slow advances once for every two fast steps.
# Complexity: O(n) time, O(1) extra space.
# Interview line: when fast reaches the end, slow is at the middle.
def middle_node(head: Optional[ListNode]) -> Optional[ListNode]:
    slow = fast = head
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
    return slow


# Pattern: deque comparison from both ends.
# Invariant: remaining deque values must still form a palindrome.
# Complexity: O(n) time, O(n) space.
# Interview line: this is the simple version; reverse second half for O(1) space.
def is_palindrome(head: Optional[ListNode]) -> bool:
    values = deque()
    while head:
        values.append(head.val)
        head = head.next

    while len(values) > 1:
        if values.popleft() != values.pop():
            return False
    return True
