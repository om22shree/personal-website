from collections import deque
from typing import Dict, List, Optional


class TreeNode:
    # Pattern: binary-tree node container.
    # Invariant: each node stores a value plus optional left/right children.
    # Complexity: O(1) time, O(1) space.
    # Interview line: tree problems are mostly about what each node returns to its parent.
    def __init__(
        self,
        val: int = 0,
        left: Optional["TreeNode"] = None,
        right: Optional["TreeNode"] = None,
    ):
        self.val = val
        self.left = left
        self.right = right

    # Pattern: debug representation.
    # Invariant: repr shows only this node value, not the full subtree.
    # Complexity: O(1) time, O(1) space.
    # Interview line: avoid recursive repr on trees because it can hide expensive traversal.
    def __repr__(self) -> str:
        return f"TreeNode({self.val})"


# Pattern: level-order tree construction.
# Invariant: queue stores parents waiting for left/right child assignment.
# Complexity: O(n) time, O(n) space.
# Interview line: consume the array in pairs while BFS assigns children.
def build_tree_level(values: List[Optional[int]]) -> Optional[TreeNode]:
    if not values or values[0] is None:
        return None

    root = TreeNode(values[0])
    q = deque([root])
    idx = 1

    while q and idx < len(values):
        node = q.popleft()
        if idx < len(values) and values[idx] is not None:
            node.left = TreeNode(values[idx])
            q.append(node.left)
        idx += 1
        if idx < len(values) and values[idx] is not None:
            node.right = TreeNode(values[idx])
            q.append(node.right)
        idx += 1

    return root


# Pattern: recursive preorder traversal.
# Invariant: visit node before its left and right subtrees.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: preorder is useful when parent work happens before children.
def preorder_recursive(root: Optional[TreeNode]) -> List[int]:
    if not root:
        return []
    return [root.val] + preorder_recursive(root.left) + preorder_recursive(root.right)


# Pattern: recursive inorder traversal.
# Invariant: visit left subtree, then node, then right subtree.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: inorder traversal of a BST yields sorted order.
def inorder_recursive(root: Optional[TreeNode]) -> List[int]:
    if not root:
        return []
    return inorder_recursive(root.left) + [root.val] + inorder_recursive(root.right)


# Pattern: recursive postorder traversal.
# Invariant: visit both children before the node.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: postorder is natural when a node depends on child results.
def postorder_recursive(root: Optional[TreeNode]) -> List[int]:
    if not root:
        return []
    return postorder_recursive(root.left) + postorder_recursive(root.right) + [root.val]


# Pattern: iterative inorder with explicit stack.
# Invariant: stack stores ancestors whose left side has been explored.
# Complexity: O(n) time, O(h) space.
# Interview line: simulate recursion by walking left, popping, then walking right.
def inorder_iterative(root: Optional[TreeNode]) -> List[int]:
    stack = []
    order = []
    cur = root

    while cur or stack:
        while cur:
            stack.append(cur)
            cur = cur.left
        cur = stack.pop()
        order.append(cur.val)
        cur = cur.right

    return order


# Pattern: BFS by levels.
# Invariant: each outer loop iteration consumes exactly one tree level.
# Complexity: O(n) time, O(width) space.
# Interview line: use len(queue) to separate levels cleanly.
def level_order(root: Optional[TreeNode]) -> List[List[int]]:
    if not root:
        return []

    levels = []
    q = deque([root])

    while q:
        level = []
        for _ in range(len(q)):
            node = q.popleft()
            level.append(node.val)
            if node.left:
                q.append(node.left)
            if node.right:
                q.append(node.right)
        levels.append(level)

    return levels


# Pattern: recursive mirror transform.
# Invariant: after processing a node, its left and right subtrees are swapped and inverted.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: invert children first by swapping pointers, then recurse into both sides.
def invert_tree(root: Optional[TreeNode]) -> Optional[TreeNode]:
    if not root:
        return None
    root.left, root.right = invert_tree(root.right), invert_tree(root.left)
    return root


# Pattern: paired tree DFS.
# Invariant: two subtrees are same only if roots match and both child pairs match.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: compare trees in lockstep: both null, one null, value mismatch, then children.
def same_tree(p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
    if not p and not q:
        return True
    if not p or not q or p.val != q.val:
        return False
    return same_tree(p.left, q.left) and same_tree(p.right, q.right)


# Pattern: level-order rightmost node capture.
# Invariant: the last node consumed at each level is visible from the right.
# Complexity: O(n) time, O(width) space.
# Interview line: BFS levels make right-side view just the final value per level.
def right_side_view(root: Optional[TreeNode]) -> List[int]:
    if not root:
        return []
    ans = []
    q = deque([root])
    while q:
        rightmost = None
        for _ in range(len(q)):
            node = q.popleft()
            rightmost = node.val
            if node.left:
                q.append(node.left)
            if node.right:
                q.append(node.right)
        ans.append(rightmost)
    return ans


# Pattern: preorder + inorder recursive construction.
# Invariant: preorder root splits the inorder interval into left and right subtrees.
# Complexity: O(n) time, O(n) space.
# Interview line: preorder gives roots; inorder tells how many nodes belong to each side.
def build_tree_pre_in(preorder: List[int], inorder: List[int]) -> Optional[TreeNode]:
    index: Dict[int, int] = {value: i for i, value in enumerate(inorder)}
    pre_i = 0

    def build(left: int, right: int) -> Optional[TreeNode]:
        nonlocal pre_i
        if left > right:
            return None
        value = preorder[pre_i]
        pre_i += 1
        root = TreeNode(value)
        mid = index[value]
        root.left = build(left, mid - 1)
        root.right = build(mid + 1, right)
        return root

    return build(0, len(inorder) - 1)
