package graphs

import (
	"container/heap"
)

type DijkstraNode struct {
	node int
	dist int
}

type DijkstraHeap []DijkstraNode

func (h DijkstraHeap) Len() int           { return len(h) }
func (h DijkstraHeap) Less(i, j int) bool { return h[i].dist < h[j].dist }
func (h DijkstraHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *DijkstraHeap) Push(x interface{}) {
	*h = append(*h, x.(DijkstraNode))
}
func (h *DijkstraHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// Pattern: Dijkstra with adjacency list and min-heap.
// Invariant: the first time a node is popped, its shortest distance is finalized.
// Complexity: O((V + E) log V) time, O(V + E) space.
// Interview line: use Dijkstra when edge weights are non-negative and you need shortest paths.
func NetworkDelayTime(times [][]int, n int, k int) int {
	type edge struct {
		to   int
		cost int
	}
	adj := make(map[int][]edge)
	for _, t := range times {
		u, v, w := t[0], t[1], t[2]
		adj[u] = append(adj[u], edge{to: v, cost: w})
	}

	dist := make(map[int]int)
	for i := 1; i <= n; i++ {
		dist[i] = 1e9 // infinite sentinel
	}
	dist[k] = 0

	h := &DijkstraHeap{}
	heap.Init(h)
	heap.Push(h, DijkstraNode{node: k, dist: 0})

	for h.Len() > 0 {
		curr := heap.Pop(h).(DijkstraNode)
		u, d := curr.node, curr.dist

		if d > dist[u] {
			continue // stale entry
		}

		for _, edge := range adj[u] {
			v, cost := edge.to, edge.cost
			if dist[u]+cost < dist[v] {
				dist[v] = dist[u] + cost
				heap.Push(h, DijkstraNode{node: v, dist: dist[v]})
			}
		}
	}

	maxDist := 0
	for i := 1; i <= n; i++ {
		if dist[i] == 1e9 {
			return -1 // Unreachable node
		}
		if dist[i] > maxDist {
			maxDist = dist[i]
		}
	}
	return maxDist
}

// Pattern: reusable single-source shortest path.
// Invariant: dist[node] is the best known cost, and stale heap entries are skipped.
// Complexity: O((V + E) log V) time, O(V + E) space.
// Interview line: pushing improved distances is simpler than decrease-key in Python/Go heaps.
func ShortestPath(n int, adj map[int][]DijkstraNode, start int) []int {
	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = 1e9
	}
	dist[start] = 0

	h := &DijkstraHeap{}
	heap.Init(h)
	heap.Push(h, DijkstraNode{node: start, dist: 0})

	for h.Len() > 0 {
		curr := heap.Pop(h).(DijkstraNode)
		u, d := curr.node, curr.dist

		if d > dist[u] {
			continue
		}

		for _, edge := range adj[u] {
			v, cost := edge.node, edge.dist
			if dist[u]+cost < dist[v] {
				dist[v] = dist[u] + cost
				heap.Push(h, DijkstraNode{node: v, dist: dist[v]})
			}
		}
	}
	return dist
}
