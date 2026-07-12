package trees

// Define TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Pattern: level-order tree construction.
// Invariant: queue stores parents waiting for left/right child assignment.
// Complexity: O(n) time, O(n) space.
// Interview line: consume the array in pairs while BFS assigns children.
func BuildTreeLevel(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: vals[0].(int)}
	q := []*TreeNode{root}
	i := 1

	for len(q) > 0 && i < len(vals) {
		curr := q[0]
		q = q[1:]

		if i < len(vals) && vals[i] != nil {
			curr.Left = &TreeNode{Val: vals[i].(int)}
			q = append(q, curr.Left)
		}
		i++

		if i < len(vals) && vals[i] != nil {
			curr.Right = &TreeNode{Val: vals[i].(int)}
			q = append(q, curr.Right)
		}
		i++
	}
	return root
}

// Pattern: recursive preorder traversal.
// Invariant: visit node before its left and right subtrees.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: preorder is useful when parent work happens before children.
func PreorderRecursive(root *TreeNode) []int {
	var res []int
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		walk(node.Left)
		walk(node.Right)
	}
	walk(root)
	return res
}

// Pattern: recursive inorder traversal.
// Invariant: visit left subtree, then node, then right subtree.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: inorder traversal of a BST yields sorted order.
func InorderRecursive(root *TreeNode) []int {
	var res []int
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		res = append(res, node.Val)
		walk(node.Right)
	}
	walk(root)
	return res
}

// Pattern: recursive postorder traversal.
// Invariant: visit both children before the node.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: postorder is natural when a node depends on child results.
func PostorderRecursive(root *TreeNode) []int {
	var res []int
	var walk func(node *TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		walk(node.Right)
		res = append(res, node.Val)
	}
	walk(root)
	return res
}

// Pattern: iterative inorder with explicit stack.
// Invariant: stack stores ancestors whose left side has been explored.
// Complexity: O(n) time, O(h) space.
// Interview line: simulate recursion by walking left, popping, then walking right.
func InorderIterative(root *TreeNode) []int {
	var res []int
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, curr.Val)
		curr = curr.Right
	}
	return res
}

// Pattern: BFS by levels.
// Invariant: each outer loop iteration consumes exactly one tree level.
// Complexity: O(n) time, O(width) space.
// Interview line: use len(queue) to separate levels cleanly.
func LevelOrder(root *TreeNode) [][]int {
	var res [][]int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		var level []int
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		res = append(res, level)
	}
	return res
}

// Pattern: recursive mirror transform.
// Invariant: after processing a node, its left and right subtrees are swapped and inverted.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: invert children first by swapping pointers, then recurse into both sides.
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	InvertTree(root.Left)
	InvertTree(root.Right)
	return root
}

// Pattern: paired tree DFS.
// Invariant: two subtrees are same only if roots match and both child pairs match.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: compare trees in lockstep: both null, one null, value mismatch, then children.
func SameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return SameTree(p.Left, q.Left) && SameTree(p.Right, q.Right)
}

// Pattern: level-order rightmost node capture.
// Invariant: the last node consumed at each level is visible from the right.
// Complexity: O(n) time, O(width) space.
// Interview line: BFS levels make right-side view just the final value per level.
func RightSideView(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		levelSize := len(q)
		for i := 0; i < levelSize; i++ {
			node := q[0]
			q = q[1:]
			if i == levelSize-1 {
				res = append(res, node.Val)
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
	}
	return res
}

// Pattern: preorder + inorder recursive construction.
// Invariant: preorder root splits the inorder interval into left and right subtrees.
// Complexity: O(n) time, O(n) space.
// Interview line: preorder gives roots; inorder tells how many nodes belong to each side.
func BuildTree(preorder []int, inorder []int) *TreeNode {
	inorderMap := make(map[int]int)
	for i, val := range inorder {
		inorderMap[val] = i
	}

	preIdx := 0
	var helper func(left int, right int) *TreeNode
	helper = func(left int, right int) *TreeNode {
		if left > right {
			return nil
		}

		rootVal := preorder[preIdx]
		preIdx++
		root := &TreeNode{Val: rootVal}

		mid := inorderMap[rootVal]
		root.Left = helper(left, mid-1)
		root.Right = helper(mid+1, right)

		return root
	}

	return helper(0, len(inorder)-1)
}
