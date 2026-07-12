package graphs

// DSU struct definition
type DSU struct {
	parent []int
	rank   []int
	count  int
}

func NewDSU(size int) *DSU {
	parent := make([]int, size)
	rank := make([]int, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 1
	}
	return &DSU{
		parent: parent,
		rank:   rank,
		count:  size,
	}
}

func (dsu *DSU) Find(i int) int {
	if dsu.parent[i] == i {
		return i
	}
	dsu.parent[i] = dsu.Find(dsu.parent[i]) // Path compression
	return dsu.parent[i]
}

func (dsu *DSU) Union(i int, j int) bool {
	rootI := dsu.Find(i)
	rootJ := dsu.Find(j)

	if rootI == rootJ {
		return false
	}

	// Union by rank
	if dsu.rank[rootI] < dsu.rank[rootJ] {
		dsu.parent[rootI] = rootJ
	} else if dsu.rank[rootI] > dsu.rank[rootJ] {
		dsu.parent[rootJ] = rootI
	} else {
		dsu.parent[rootJ] = rootI
		dsu.rank[rootI]++
	}
	dsu.count--
	return true
}

// Pattern: DSU cycle detection in an undirected graph.
// Invariant: each successful union connects two previously separate components.
// Complexity: O(E alpha(V)) time, O(V) space.
// Interview line: the first edge whose endpoints are already connected is redundant.
func FindRedundantConnection(edges [][]int) []int {
	n := len(edges)
	dsu := NewDSU(n + 1) // 1-indexed nodes

	for _, edge := range edges {
		u, v := edge[0], edge[1]
		if !dsu.Union(u, v) {
			return edge
		}
	}
	return nil
}

// Pattern: DSU component counting.
// Invariant: count decreases only when a union merges two different components.
// Complexity: O(E alpha(V)) time, O(V) space.
// Interview line: union every edge, then the DSU count is the number of connected components.
func CountComponents(n int, edges [][]int) int {
	dsu := NewDSU(n)
	for _, edge := range edges {
		dsu.Union(edge[0], edge[1])
	}
	return dsu.count
}
