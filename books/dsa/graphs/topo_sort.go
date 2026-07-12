package graphs

// Pattern: Kahn topological sort.
// Invariant: queue contains nodes with zero remaining prerequisites.
// Complexity: O(V + E) time, O(V + E) space.
// Interview line: if we cannot consume every course, a dependency cycle remains.
func CanFinish(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		adj[prereq] = append(adj[prereq], course)
		indegree[course]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	visited := 0
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		visited++

		for _, neighbor := range adj[curr] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	return visited == numCourses
}

// Pattern: topological ordering with indegrees.
// Invariant: append a node only after all prerequisites have been consumed.
// Complexity: O(V + E) time, O(V + E) space.
// Interview line: the order exists only if the topo process visits every node.
func FindOrder(numCourses int, prerequisites [][]int) []int {
	adj := make(map[int][]int)
	indegree := make([]int, numCourses)

	for _, pre := range prerequisites {
		course, prereq := pre[0], pre[1]
		adj[prereq] = append(adj[prereq], course)
		indegree[course]++
	}

	var q []int
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	var order []int
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		order = append(order, curr)

		for _, neighbor := range adj[curr] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				q = append(q, neighbor)
			}
		}
	}

	if len(order) == numCourses {
		return order
	}
	return nil
}

// Pattern: DFS color marking for directed cycle detection.
// Invariant: color 1 means currently on recursion stack, color 2 means fully processed.
// Complexity: O(V + E) time, O(V + E) space.
// Interview line: seeing a gray node again means we found a back edge and therefore a cycle.
func HasCycleDFS(numCourses int, prerequisites [][]int) bool {
	adj := make(map[int][]int)
	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
	}

	// 0: unvisited, 1: visiting (gray), 2: visited (black)
	colors := make([]int, numCourses)

	var dfs func(node int) bool
	dfs = func(node int) bool {
		colors[node] = 1 // Mark visiting
		for _, neighbor := range adj[node] {
			if colors[neighbor] == 1 {
				return true // Cycle detected
			}
			if colors[neighbor] == 0 {
				if dfs(neighbor) {
					return true
				}
			}
		}
		colors[node] = 2 // Mark visited
		return false
	}

	for i := 0; i < numCourses; i++ {
		if colors[i] == 0 {
			if dfs(i) {
				return true
			}
		}
	}
	return false
}
