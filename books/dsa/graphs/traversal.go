package graphs

// Pattern: adjacency-list construction.
// Invariant: every edge is represented once for directed graphs, twice for undirected graphs.
// Complexity: O(E) time, O(V + E) space.
// Interview line: map[int][]int keeps graph construction concise and avoids missing-key checks.
func BuildGraph(edges [][]int, directed bool) map[int][]int {
	graph := make(map[int][]int)
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], v)
		if !directed {
			graph[v] = append(graph[v], u)
		}
	}
	return graph
}

// Pattern: BFS with a queue.
// Invariant: seen contains every node already enqueued, so nodes are processed once.
// Complexity: O(V + E) time, O(V) space.
// Interview line: slice acts as queue with O(1) popleft for level-by-level or shortest-hop traversal.
func BFSOrder(start int, graph map[int][]int) []int {
	var res []int
	seen := make(map[int]bool)
	seen[start] = true
	q := []int{start}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		res = append(res, curr)

		for _, neighbor := range graph[curr] {
			if !seen[neighbor] {
				seen[neighbor] = true
				q = append(q, neighbor)
			}
		}
	}
	return res
}

// Pattern: recursive DFS.
// Invariant: seen prevents revisiting nodes and recursion explores one path fully.
// Complexity: O(V + E) time, O(V) space for recursion/seen.
// Interview line: DFS is natural when you need reachability, components, or backtracking.
func DFSOrder(start int, graph map[int][]int) []int {
	var res []int
	seen := make(map[int]bool)

	var dfs func(node int)
	dfs = func(node int) {
		seen[node] = true
		res = append(res, node)
		for _, neighbor := range graph[node] {
			if !seen[neighbor] {
				dfs(neighbor)
			}
		}
	}

	dfs(start)
	return res
}

// Pattern: grid BFS flood fill.
// Invariant: once land is queued, mark it water so it is not counted again.
// Complexity: O(rows * cols) time, O(rows * cols) worst-case space.
// Interview line: every time we discover unvisited land, that starts one island traversal.
func NumIslands(grid [][]byte) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])
	islands := 0

	bfs := func(r int, c int) {
		type point struct{ r, c int }
		q := []point{{r, c}}
		grid[r][c] = '0' // mark visited

		dirs := []point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for len(q) > 0 {
			curr := q[0]
			q = q[1:]

			for _, d := range dirs {
				nr, nc := curr.r+d.r, curr.c+d.c
				if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == '1' {
					grid[nr][nc] = '0'
					q = append(q, point{nr, nc})
				}
			}
		}
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '1' {
				islands++
				bfs(r, c)
			}
		}
	}
	return islands
}

// Pattern: recursive grid DFS flood fill.
// Invariant: each matching cell is consumed before exploring its four neighbors.
// Complexity: O(rows * cols) time, O(rows * cols) recursion space worst case.
// Interview line: mutate visited cells in-place when the prompt allows it.
func FloodFill(image [][]int, sr int, sc int, color int) [][]int {
	startColor := image[sr][sc]
	if startColor == color {
		return image
	}
	rows := len(image)
	cols := len(image[0])

	var dfs func(r int, c int)
	dfs = func(r int, c int) {
		if r < 0 || r >= rows || c < 0 || c >= cols || image[r][c] != startColor {
			return
		}
		image[r][c] = color
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	dfs(sr, sc)
	return image
}
