package trees

// Pattern: DFS with lower/upper bounds.
// Invariant: every node must fit inside the range imposed by its ancestors.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: local child comparisons are not enough; carry ancestor bounds.
func IsValidBST(root *TreeNode) bool {
	var validate func(node *TreeNode, min *int, max *int) bool
	validate = func(node *TreeNode, min *int, max *int) bool {
		if node == nil {
			return true
		}
		if min != nil && node.Val <= *min {
			return false
		}
		if max != nil && node.Val >= *max {
			return false
		}
		return validate(node.Left, min, &node.Val) && validate(node.Right, &node.Val, max)
	}
	return validate(root, nil, nil)
}

// Pattern: iterative inorder traversal.
// Invariant: kth popped node in inorder is the kth smallest BST value.
// Complexity: O(h + k) time, O(h) space.
// Interview line: BST inorder order is sorted, so stop as soon as k reaches zero.
func KthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	curr := root
	count := k

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		count--
		if count == 0 {
			return curr.Val
		}

		curr = curr.Right
	}
	return -1
}

// Pattern: BST-guided descent.
// Invariant: if both targets are on one side, LCA must be on that side.
// Complexity: O(h) time, O(1) extra space.
// Interview line: the split point is the lowest node with p and q on different sides.
func LowestCommonAncestorBST(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	curr := root
	for curr != nil {
		if p.Val < curr.Val && q.Val < curr.Val {
			curr = curr.Left
		} else if p.Val > curr.Val && q.Val > curr.Val {
			curr = curr.Right
		} else {
			return curr
		}
	}
	return nil
}
