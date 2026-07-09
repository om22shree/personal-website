from typing import List, Optional

from .traversal import TreeNode


# Pattern: recursive tree height.
# Invariant: depth is one plus the larger child depth.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: each subtree returns its height to its parent.
def max_depth(root: Optional[TreeNode]) -> int:
    if not root:
        return 0
    return 1 + max(max_depth(root.left), max_depth(root.right))


# Pattern: postorder height with global diameter.
# Invariant: best diameter through a node is left_height + right_height.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: return height upward, update diameter locally at every node.
def diameter_of_binary_tree(root: Optional[TreeNode]) -> int:
    diameter = 0

    # Pattern: height-returning postorder helper.
    # Invariant: update diameter before returning this subtree's height.
    # Complexity: O(size of subtree) local total, O(h) recursion space overall.
    # Interview line: the longest path through a node uses both child heights.
    def height(node: Optional[TreeNode]) -> int:
        nonlocal diameter
        if not node:
            return 0
        left = height(node.left)
        right = height(node.right)
        diameter = max(diameter, left + right)
        return 1 + max(left, right)

    height(root)
    return diameter


# Pattern: postorder height with early failure.
# Invariant: -1 means the subtree is already unbalanced.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: compute balance and height in one DFS instead of recomputing heights.
def is_balanced(root: Optional[TreeNode]) -> bool:
    # Pattern: height helper with sentinel failure.
    # Invariant: -1 propagates an unbalanced subtree upward immediately.
    # Complexity: O(size of subtree) total, O(h) recursion space overall.
    # Interview line: combine height calculation and balance checking in one pass.
    def height(node: Optional[TreeNode]) -> int:
        if not node:
            return 0
        left = height(node.left)
        if left == -1:
            return -1
        right = height(node.right)
        if right == -1 or abs(left - right) > 1:
            return -1
        return 1 + max(left, right)

    return height(root) != -1


# Pattern: DFS backtracking from root to leaves.
# Invariant: path contains the current root-to-node path and remaining is updated.
# Complexity: O(n) time excluding output, O(h) recursion space.
# Interview line: append before recursion and pop after so sibling paths stay clean.
def path_sum(root: Optional[TreeNode], target_sum: int) -> List[List[int]]:
    paths = []

    # Pattern: root-to-leaf backtracking helper.
    # Invariant: path exactly matches the current recursion path.
    # Complexity: O(size of subtree) excluding output, O(h) recursion space overall.
    # Interview line: append before child calls and pop after child calls.
    def dfs(node: Optional[TreeNode], remaining: int, path: List[int]) -> None:
        if not node:
            return
        path.append(node.val)
        remaining -= node.val
        if not node.left and not node.right and remaining == 0:
            paths.append(path[:])
        dfs(node.left, remaining, path)
        dfs(node.right, remaining, path)
        path.pop()

    dfs(root, target_sum, [])
    return paths


# Pattern: postorder max-gain DP on a tree.
# Invariant: returned gain is the best single-branch path extendable to the parent.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: update global best with both branches, but return only one branch upward.
def max_path_sum(root: Optional[TreeNode]) -> int:
    best = float("-inf")

    # Pattern: max extendable gain helper.
    # Invariant: return one branch upward, but update best with both branches.
    # Complexity: O(size of subtree) total, O(h) recursion space overall.
    # Interview line: paths can split at the current node only for the global answer.
    def gain(node: Optional[TreeNode]) -> int:
        nonlocal best
        if not node:
            return 0
        left = max(gain(node.left), 0)
        right = max(gain(node.right), 0)
        best = max(best, node.val + left + right)
        return node.val + max(left, right)

    gain(root)
    return int(best)


# Pattern: recursive subtree matching.
# Invariant: subroot matches if some node in root starts an identical tree.
# Complexity: O(n m) worst-case time, O(h) recursion space.
# Interview line: at each node, either match from here or keep searching both children.
def is_subtree(root: Optional[TreeNode], subroot: Optional[TreeNode]) -> bool:
    if not subroot:
        return True
    if not root:
        return False

    def same(a: Optional[TreeNode], b: Optional[TreeNode]) -> bool:
        if not a and not b:
            return True
        if not a or not b or a.val != b.val:
            return False
        return same(a.left, b.left) and same(a.right, b.right)

    return same(root, subroot) or is_subtree(root.left, subroot) or is_subtree(root.right, subroot)


# Pattern: DFS with path maximum.
# Invariant: max_so_far is the largest value on the root-to-parent path.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: a node is good if no ancestor on the path is greater.
def good_nodes(root: Optional[TreeNode]) -> int:
    def dfs(node: Optional[TreeNode], max_so_far: int) -> int:
        if not node:
            return 0
        good = 1 if node.val >= max_so_far else 0
        next_max = max(max_so_far, node.val)
        return good + dfs(node.left, next_max) + dfs(node.right, next_max)

    return dfs(root, float("-inf"))


# Pattern: binary-tree LCA postorder.
# Invariant: a subtree returns p/q if found, or LCA if both sides contain targets.
# Complexity: O(n) time, O(h) recursion space.
# Interview line: if p and q split across left and right, the current node is their LCA.
def lowest_common_ancestor(
    root: Optional[TreeNode], p: TreeNode, q: TreeNode
) -> Optional[TreeNode]:
    if not root or root is p or root is q:
        return root
    left = lowest_common_ancestor(root.left, p, q)
    right = lowest_common_ancestor(root.right, p, q)
    if left and right:
        return root
    return left or right
