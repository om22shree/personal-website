import heapq
from typing import List, Optional

from .basics import ListNode, reverse_list


# Pattern: two sorted lists with dummy tail.
# Invariant: tail points to the end of the merged sorted prefix.
# Complexity: O(m + n) time, O(1) extra space.
# Interview line: repeatedly attach the smaller head and advance that list.
def merge_two_lists(
    list1: Optional[ListNode], list2: Optional[ListNode]
) -> Optional[ListNode]:
    dummy = ListNode()
    tail = dummy

    while list1 and list2:
        if list1.val <= list2.val:
            tail.next = list1
            list1 = list1.next
        else:
            tail.next = list2
            list2 = list2.next
        tail = tail.next

    tail.next = list1 or list2
    return dummy.next


# Pattern: k-way merge with min-heap.
# Invariant: heap stores the current smallest candidate from each list.
# Complexity: O(n log k) time, O(k) space.
# Interview line: use index as a tie-breaker so equal values do not compare nodes.
def merge_k_lists(lists: List[Optional[ListNode]]) -> Optional[ListNode]:
    heap = []
    for idx, node in enumerate(lists):
        if node:
            heapq.heappush(heap, (node.val, idx, node))

    dummy = ListNode()
    tail = dummy

    while heap:
        _, idx, node = heapq.heappop(heap)
        tail.next = node
        tail = tail.next
        if node.next:
            heapq.heappush(heap, (node.next.val, idx, node.next))

    return dummy.next


# Pattern: two pointers with an n-node gap.
# Invariant: fast is n nodes ahead of slow before both move together.
# Complexity: O(n) time, O(1) extra space.
# Interview line: a dummy node makes deleting the head node clean.
def remove_nth_from_end(head: Optional[ListNode], n: int) -> Optional[ListNode]:
    dummy = ListNode(0, head)
    fast = slow = dummy

    for _ in range(n):
        fast = fast.next

    while fast.next:
        fast = fast.next
        slow = slow.next

    slow.next = slow.next.next
    return dummy.next


# Pattern: split, reverse second half, then weave.
# Invariant: first and reversed second halves are alternately connected.
# Complexity: O(n) time, O(1) extra space.
# Interview line: find middle, reverse the back half, then merge one node at a time.
def reorder_list(head: Optional[ListNode]) -> None:
    if not head or not head.next:
        return

    slow = fast = head
    while fast.next and fast.next.next:
        slow = slow.next
        fast = fast.next.next

    second = reverse_list(slow.next)
    slow.next = None
    first = head

    while second:
        first_next, second_next = first.next, second.next
        first.next = second
        second.next = first_next
        first = first_next
        second = second_next
