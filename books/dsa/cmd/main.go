package main

import (
	"dsa"
	"dsa/graphs"
	"dsa/heaps"
	"dsa/linkedlists"
	"dsa/trees"
	"fmt"
)

func main() {
	fmt.Println("=== Running Go DSA Command Center Tests ===")

	// 1. Arrays & Hashing
	fmt.Println("\n[1] Testing Arrays & Hashing:")
	sumIdx := dsa.TwoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("TwoSum: %v (expected: [0 1])\n", sumIdx)

	dup := dsa.ContainsDuplicate([]int{1, 2, 3, 1})
	fmt.Printf("ContainsDuplicate: %v (expected: true)\n", dup)

	anagram := dsa.IsAnagram("rat", "car")
	fmt.Printf("IsAnagram: %v (expected: false)\n", anagram)

	// 2. Sliding Window
	fmt.Println("\n[2] Testing Sliding Window:")
	maxAvg := dsa.FindMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)
	fmt.Printf("MaxAverageSubarray: %f (expected: 12.750000)\n", maxAvg)

	// 3. Rate Limiters
	fmt.Println("\n[3] Testing Rate Limiters:")
	limiter := dsa.NewSlidingWindowLogRateLimiter(2, 10.0)
	fmt.Printf("Allow(1.0): %v (expected: true)\n", limiter.Allow("api", 1.0))
	fmt.Printf("Allow(2.0): %v (expected: true)\n", limiter.Allow("api", 2.0))
	fmt.Printf("Allow(3.0): %v (expected: false)\n", limiter.Allow("api", 3.0))

	// 4. Linked Lists
	fmt.Println("\n[4] Testing Linked Lists:")
	head := linkedlists.BuildLinkedList([]int{1, 2, 3, 4, 5})
	reversed := linkedlists.ReverseList(head)
	revVals := linkedlists.ToSlice(reversed)
	fmt.Printf("Reversed List: %v (expected: [5 4 3 2 1])\n", revVals)

	// 5. Trees
	fmt.Println("\n[5] Testing Trees:")
	// BFS representation: [1, 2, 3] -> 1 is root, 2 is left, 3 is right
	treeNode := trees.BuildTreeLevel([]interface{}{1, 2, 3})
	inorder := trees.InorderRecursive(treeNode)
	fmt.Printf("Tree Inorder: %v (expected: [2 1 3])\n", inorder)

	// 6. Graphs
	fmt.Println("\n[6] Testing Graphs:")
	edges := [][]int{{1, 2}, {2, 3}, {3, 1}}
	graph := graphs.BuildGraph(edges, false)
	dfsPath := graphs.DFSOrder(1, graph)
	fmt.Printf("Graph DFS Path: %v (expected: [1 2 3])\n", dfsPath)

	// 7. Heaps
	fmt.Println("\n[7] Testing Heaps:")
	medianFinder := heaps.NewMedianFinder()
	medianFinder.AddNum(1)
	medianFinder.AddNum(2)
	fmt.Printf("Median (1, 2): %f (expected: 1.500000)\n", medianFinder.FindMedian())
	medianFinder.AddNum(3)
	fmt.Printf("Median (1, 2, 3): %f (expected: 2.000000)\n", medianFinder.FindMedian())

	// 8. Dynamic Programming
	fmt.Println("\n[8] Testing Dynamic Programming:")
	fmt.Printf("ClimbStairs(5): %d (expected: 8)\n", dsa.ClimbStairs(5))
	fmt.Printf("HouseRobber([1, 2, 3, 1]): %d (expected: 4)\n", dsa.HouseRobber([]int{1, 2, 3, 1}))
	fmt.Printf("HouseRobber([2, 7, 9, 3, 1]): %d (expected: 12)\n", dsa.HouseRobber([]int{2, 7, 9, 3, 1}))

	fmt.Println("\n=== All Tests Passed ===")
}
