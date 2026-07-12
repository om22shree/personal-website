package graphs

import (
	"container/heap"
	"sort"
)

// Graph Node container
type Node struct {
	Val       int
	Neighbors []*Node
}

// Pattern: DFS clone with visited map.
// Invariant: clones maps every visited original node to its copy.
// Complexity: O(V + E) time, O(V) space.
// Interview line: create the copy before cloning neighbors to break cycles.
func CloneGraph(node *Node) *Node {
	clones := make(map[*Node]*Node)

	var dfs func(cur *Node) *Node
	dfs = func(cur *Node) *Node {
		if cur == nil {
			return nil
		}
		if copyNode, ok := clones[cur]; ok {
			return copyNode
		}
		copyNode := &Node{Val: cur.Val}
		clones[cur] = copyNode
		for _, nei := range cur.Neighbors {
			if nei != nil {
				copyNode.Neighbors = append(copyNode.Neighbors, dfs(nei))
			}
		}
		return copyNode
	}

	return dfs(node)
}

// Pattern: topological sort over character precedence.
// Invariant: indegree zero characters have all prerequisites satisfied.
// Complexity: O(total chars + edges) time, O(unique chars) space.
// Interview line: the first differing character between adjacent words creates the ordering edge.
func AlienOrder(words []string) string {
	graph := make(map[byte]map[byte]bool)
	indegree := make(map[byte]int)

	// Initialize maps for all characters in all words
	for _, word := range words {
		for i := 0; i < len(word); i++ {
			ch := word[i]
			if _, ok := graph[ch]; !ok {
				graph[ch] = make(map[byte]bool)
				indegree[ch] = 0
			}
		}
	}

	for i := 0; i < len(words)-1; i++ {
		first := words[i]
		second := words[i+1]

		// Check prefix condition (e.g. "abc", "ab" is invalid)
		if len(first) > len(second) && first[:len(second)] == second {
			return ""
		}

		minLen := len(first)
		if len(second) < minLen {
			minLen = len(second)
		}

		for j := 0; j < minLen; j++ {
			a := first[j]
			b := second[j]
			if a != b {
				if !graph[a][b] {
					graph[a][b] = true
					indegree[b]++
				}
				break
			}
		}
	}

	var q []byte
	for ch, deg := range indegree {
		if deg == 0 {
			q = append(q, ch)
		}
	}

	var order []byte
	for len(q) > 0 {
		ch := q[0]
		q = q[1:]
		order = append(order, ch)

		for nei := range graph[ch] {
			indegree[nei]--
			if indegree[nei] == 0 {
				q = append(q, nei)
			}
		}
	}

	if len(order) == len(graph) {
		return string(order)
	}
	return ""
}

type MSTEdge struct {
	cost  int
	index int
}

type MSTHeap []MSTEdge

func (h MSTHeap) Len() int           { return len(h) }
func (h MSTHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h MSTHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MSTHeap) Push(x interface{}) {
	*h = append(*h, x.(MSTEdge))
}
func (h *MSTHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Pattern: Prim minimum spanning tree.
// Invariant: heap stores cheapest edges from visited points to unvisited points.
// Complexity: O(n^2 log n) time, O(n^2) heap space.
// Interview line: grow the connected set by always taking the cheapest crossing edge.
func MinCostConnectPoints(points [][]int) int {
	n := len(points)
	visited := make(map[int]bool)
	h := &MSTHeap{}
	heap.Init(h)
	heap.Push(h, MSTEdge{cost: 0, index: 0})
	total := 0

	for visitedCount := 0; visitedCount < n; {
		curr := heap.Pop(h).(MSTEdge)
		cost, i := curr.cost, curr.index
		if visited[i] {
			continue
		}
		visited[i] = true
		total += cost
		visitedCount++

		x1, y1 := points[i][0], points[i][1]
		for j := 0; j < n; j++ {
			if !visited[j] {
				x2, y2 := points[j][0], points[j][1]
				dist := abs(x1-x2) + abs(y1-y2)
				heap.Push(h, MSTEdge{cost: dist, index: j})
			}
		}
	}

	return total
}

type SwimNode struct {
	time int
	r    int
	c    int
}

type SwimHeap []SwimNode

func (h SwimHeap) Len() int           { return len(h) }
func (h SwimHeap) Less(i, j int) bool { return h[i].time < h[j].time }
func (h SwimHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *SwimHeap) Push(x interface{}) {
	*h = append(*h, x.(SwimNode))
}
func (h *SwimHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: Dijkstra over grid states.
// Invariant: first time bottom-right is popped, its effort is minimal.
// Complexity: O(mn log mn) time, O(mn) space.
// Interview line: path cost is the maximum cell height seen, so prioritize lower current effort.
func SwimInWater(grid [][]int) int {
	n := len(grid)
	h := &SwimHeap{}
	heap.Init(h)
	heap.Push(h, SwimNode{time: grid[0][0], r: 0, c: 0})

	seen := make(map[int]bool)
	seen[0] = true // 0 represents (0, 0) in grid

	dirs := []struct{ dr, dc int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	maxVal := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	for h.Len() > 0 {
		curr := heap.Pop(h).(SwimNode)
		time, r, c := curr.time, curr.r, curr.c

		if r == n-1 && c == n-1 {
			return time
		}

		for _, d := range dirs {
			nr, nc := r+d.dr, c+d.dc
			if nr >= 0 && nr < n && nc >= 0 && nc < n {
				flatIdx := nr*n + nc
				if !seen[flatIdx] {
					seen[flatIdx] = true
					heap.Push(h, SwimNode{time: maxVal(time, grid[nr][nc]), r: nr, c: nc})
				}
			}
		}
	}
	return -1
}

type oceanPoint struct {
	r, c int
}

// Pattern: multi-source boundary reachability.
// Invariant: a cell belongs to an ocean set if water can flow from it to that ocean.
// Complexity: O(mn) time, O(mn) space.
// Interview line: reverse the flow and start DFS from ocean borders.
func PacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return nil
	}
	rows := len(heights)
	cols := len(heights[0])

	flow := func(starts []oceanPoint) map[oceanPoint]bool {
		seen := make(map[oceanPoint]bool)
		var stack []oceanPoint
		for _, s := range starts {
			seen[s] = true
			stack = append(stack, s)
		}

		dirs := []struct{ dr, dc int }{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for len(stack) > 0 {
			curr := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			r, c := curr.r, curr.c

			for _, d := range dirs {
				nr, nc := r+d.dr, c+d.dc
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
					np := oceanPoint{r: nr, c: nc}
					if !seen[np] && heights[nr][nc] >= heights[r][c] {
						seen[np] = true
						stack = append(stack, np)
					}
				}
			}
		}
		return seen
	}

	var pacificStarts []oceanPoint
	var atlanticStarts []oceanPoint

	for c := 0; c < cols; c++ {
		pacificStarts = append(pacificStarts, oceanPoint{r: 0, c: c})
		atlanticStarts = append(atlanticStarts, oceanPoint{r: rows - 1, c: c})
	}
	for r := 0; r < rows; r++ {
		pacificStarts = append(pacificStarts, oceanPoint{r: r, c: 0})
		atlanticStarts = append(atlanticStarts, oceanPoint{r: r, c: cols - 1})
	}

	pacific := flow(pacificStarts)
	atlantic := flow(atlanticStarts)

	var res [][]int
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			p := oceanPoint{r: r, c: c}
			if pacific[p] && atlantic[p] {
				res = append(res, []int{r, c})
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] < res[j][1]
		}
		return res[i][0] < res[j][0]
	})

	return res
}

// Pattern: BFS coloring.
// Invariant: every colored edge connects opposite colors.
// Complexity: O(V + E) time, O(V) space.
// Interview line: a graph is bipartite if every component can be two-colored.
func IsBipartite(graph [][]int) bool {
	color := make(map[int]int)

	for start := 0; start < len(graph); start++ {
		if _, ok := color[start]; ok {
			continue
		}
		color[start] = 0
		q := []int{start}

		for len(q) > 0 {
			node := q[0]
			q = q[1:]

			for _, nei := range graph[node] {
				if c, ok := color[nei]; ok {
					if c == color[node] {
						return false
					}
				} else {
					color[nei] = 1 - color[node]
					q = append(q, nei)
				}
			}
		}
	}
	return true
}
