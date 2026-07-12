package trees

import (
	"strconv"
	"strings"
)

// Pattern: BFS serialization with null sentinels.
// Invariant: output preserves enough structure to rebuild left/right child positions.
// Complexity: O(n) time, O(n) space.
// Interview line: include null markers during traversal, then trim trailing nulls for compactness.
func Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	var values []string
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node != nil {
			values = append(values, strconv.Itoa(node.Val))
			q = append(q, node.Left)
			q = append(q, node.Right)
		} else {
			values = append(values, "#")
		}
	}

	// Trim trailing nulls
	for len(values) > 0 && values[len(values)-1] == "#" {
		values = values[:len(values)-1]
	}

	return strings.Join(values, ",")
}

// Pattern: BFS deserialization from level-order tokens.
// Invariant: queue stores parents waiting for child tokens.
// Complexity: O(n) time, O(n) space.
// Interview line: consume two tokens per queued parent to restore the original shape.
func Deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	values := strings.Split(data, ",")
	val, _ := strconv.Atoi(values[0])
	root := &TreeNode{Val: val}
	q := []*TreeNode{root}
	idx := 1

	for len(q) > 0 && idx < len(values) {
		node := q[0]
		q = q[1:]

		if idx < len(values) && values[idx] != "#" {
			v, _ := strconv.Atoi(values[idx])
			node.Left = &TreeNode{Val: v}
			q = append(q, node.Left)
		}
		idx++

		if idx < len(values) && values[idx] != "#" {
			v, _ := strconv.Atoi(values[idx])
			node.Right = &TreeNode{Val: v}
			q = append(q, node.Right)
		}
		idx++
	}

	return root
}
