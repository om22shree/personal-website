package trees

// Pattern: recursive tree height.
// Invariant: depth is one plus the larger child depth.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: each subtree returns its height to its parent.
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return 1 + left
	}
	return 1 + right
}

// Pattern: postorder height with global diameter.
// Invariant: best diameter through a node is left_height + right_height.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: return height upward, update diameter locally at every node.
func DiameterOfBinaryTree(root *TreeNode) int {
	diameter := 0

	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := height(node.Left)
		right := height(node.Right)
		if left+right > diameter {
			diameter = left + right
		}
		if left > right {
			return 1 + left
		}
		return 1 + right
	}

	height(root)
	return diameter
}

func absInt(val int) int {
	if val < 0 {
		return -val
	}
	return val
}

// Pattern: postorder height with early failure.
// Invariant: -1 means the subtree is already unbalanced.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: compute balance and height in one DFS instead of recomputing heights.
func IsBalanced(root *TreeNode) bool {
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := height(node.Left)
		if left == -1 {
			return -1
		}
		right := height(node.Right)
		if right == -1 {
			return -1
		}
		if absInt(left-right) > 1 {
			return -1
		}
		if left > right {
			return 1 + left
		}
		return 1 + right
	}
	return height(root) != -1
}

// Pattern: DFS backtracking from root to leaves.
// Invariant: path contains the current root-to-node path and remaining is updated.
// Complexity: O(n) time excluding output, O(h) recursion space.
// Interview line: append before recursion and pop after so sibling paths stay clean.
func PathSum(root *TreeNode, targetSum int) [][]int {
	var paths [][]int
	var path []int

	var dfs func(node *TreeNode, remaining int)
	dfs = func(node *TreeNode, remaining int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		remaining -= node.Val
		if node.Left == nil && node.Right == nil && remaining == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			paths = append(paths, temp)
		}
		dfs(node.Left, remaining)
		dfs(node.Right, remaining)
		path = path[:len(path)-1]
	}

	dfs(root, targetSum)
	return paths
}

// Pattern: postorder max-gain DP on a tree.
// Invariant: returned gain is the best single-branch path extendable to the parent.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: update global best with both branches, but return only one branch upward.
func MaxPathSum(root *TreeNode) int {
	best := -1000000000 // sentinel minus infinity

	var gain func(node *TreeNode) int
	gain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		leftVal := gain(node.Left)
		rightVal := gain(node.Right)

		left := 0
		if leftVal > 0 {
			left = leftVal
		}
		right := 0
		if rightVal > 0 {
			right = rightVal
		}

		if node.Val+left+right > best {
			best = node.Val + left + right
		}

		maxChild := left
		if right > maxChild {
			maxChild = right
		}
		return node.Val + maxChild
	}

	gain(root)
	return best
}

func sameSubtree(a *TreeNode, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil || a.Val != b.Val {
		return false
	}
	return sameSubtree(a.Left, b.Left) && sameSubtree(a.Right, b.Right)
}

// Pattern: recursive subtree matching.
// Invariant: subroot matches if some node in root starts an identical tree.
// Complexity: O(n m) worst-case time, O(h) recursion space.
// Interview line: at each node, either match from here or keep searching both children.
func IsSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil {
		return false
	}
	if sameSubtree(root, subRoot) {
		return true
	}
	return IsSubtree(root.Left, subRoot) || IsSubtree(root.Right, subRoot)
}

// Pattern: DFS with path maximum.
// Invariant: max_so_far is the largest value on the root-to-parent path.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: a node is good if no ancestor on the path is greater.
func GoodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, maxSoFar int) int
	dfs = func(node *TreeNode, maxSoFar int) int {
		if node == nil {
			return 0
		}
		good := 0
		if node.Val >= maxSoFar {
			good = 1
		}
		nextMax := maxSoFar
		if node.Val > nextMax {
			nextMax = node.Val
		}
		return good + dfs(node.Left, nextMax) + dfs(node.Right, nextMax)
	}
	return dfs(root, -1000000000)
}

// Pattern: binary-tree LCA postorder.
// Invariant: a subtree returns p/q if found, or LCA if both sides contain targets.
// Complexity: O(n) time, O(h) recursion space.
// Interview line: if p and q split across left and right, the current node is their LCA.
func LowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}
