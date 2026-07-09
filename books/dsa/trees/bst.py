from typing import Optional

from .traversal import TreeNode


# Pattern: DFS with lower/upper bounds.
# Invariant: every node must fit inside the range imposed by its ancestors.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: local child comparisons are not enough; carry ancestor bounds.
def is_valid_bst(root: Optional[TreeNode]) -> bool:
    # Pattern: recursive bounds helper.
    # Invariant: node must remain strictly between low and high.
    # Complexity: O(size of subtree) time, O(h) recursion space overall.
    # Interview line: pass narrowed bounds downward from ancestors.
    def valid(node: Optional[TreeNode], low: float, high: float) -> bool:
        if not node:
            return True
        if not low < node.val < high:
            return False
        return valid(node.left, low, node.val) and valid(node.right, node.val, high)

    return valid(root, float("-inf"), float("inf"))


# Pattern: iterative inorder traversal.
# Invariant: kth popped node in inorder is the kth smallest BST value.
# Complexity: O(h + k) time, O(h) space.
# Interview line: BST inorder order is sorted, so stop as soon as k reaches zero.
def kth_smallest(root: Optional[TreeNode], k: int) -> int:
    stack = []
    cur = root

    while cur or stack:
        while cur:
            stack.append(cur)
            cur = cur.left
        cur = stack.pop()
        k -= 1
        if k == 0:
            return cur.val
        cur = cur.right

    raise ValueError("k is larger than the number of nodes")


# Pattern: BST-guided descent.
# Invariant: if both targets are on one side, LCA must be on that side.
# Complexity: O(h) time, O(1) extra space.
# Interview line: the split point is the lowest node with p and q on different sides.
def lowest_common_ancestor_bst(
    root: Optional[TreeNode], p: TreeNode, q: TreeNode
) -> Optional[TreeNode]:
    cur = root
    while cur:
        if p.val < cur.val and q.val < cur.val:
            cur = cur.left
        elif p.val > cur.val and q.val > cur.val:
            cur = cur.right
        else:
            return cur
    return None
