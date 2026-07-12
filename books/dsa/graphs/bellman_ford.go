package graphs

// Pattern: bounded Bellman-Ford over edge count.
// Invariant: after i rounds, prices uses at most i edges.
// Complexity: O(k * E) time, O(V) space.
// Interview line: copy the previous price array so one round only adds one extra flight.
func FindCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	prices := make([]int, n)
	for i := range prices {
		prices[i] = 1e9 // infinite sentinel
	}
	prices[src] = 0

	for i := 0; i <= k; i++ {
		// Create a copy of the prices to isolate changes to current round
		temp := make([]int, n)
		copy(temp, prices)

		for _, flight := range flights {
			u, v, w := flight[0], flight[1], flight[2]
			if prices[u] == 1e9 {
				continue
			}
			if prices[u]+w < temp[v] {
				temp[v] = prices[u] + w
			}
		}
		prices = temp
	}

	if prices[dst] == 1e9 {
		return -1
	}
	return prices[dst]
}
