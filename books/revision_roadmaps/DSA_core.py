"""
NeetCode 150 - High-Signal Python Reference
Single file containing all core patterns, imports, and data structures
"""

# ============================================================================
# IMPORTS (The Essentials)
# ============================================================================
from collections import deque, defaultdict, Counter, OrderedDict
import heapq
import bisect
from typing import List, Optional, Dict, Set, Tuple
import math

# ============================================================================
# DATA STRUCTURE DEFINITIONS (Object Shapes)
# ============================================================================

# Definition for singly-linked list
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

# Definition for a binary tree node
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

# Definition for an N-ary tree node
class Node:
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

# Definition for a graph node (adjacency list representation)
class GraphNode:
    def __init__(self, val=0, neighbors=None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []

# ============================================================================
# CORE PATTERNS - TREES
# ============================================================================

# DFS Traversals (Recursive)
def preorder(root: Optional[TreeNode]) -> List[int]:
    """Root -> Left -> Right"""
    if not root:
        return []
    return [root.val] + preorder(root.left) + preorder(root.right)

def inorder(root: Optional[TreeNode]) -> List[int]:
    """Left -> Root -> Right (gives sorted order for BST)"""
    if not root:
        return []
    return inorder(root.left) + [root.val] + inorder(root.right)

def postorder(root: Optional[TreeNode]) -> List[int]:
    """Left -> Right -> Root"""
    if not root:
        return []
    return postorder(root.left) + postorder(root.right) + [root.val]

# BFS Level-Order Traversal
def levelOrder(root: Optional[TreeNode]) -> List[List[int]]:
    """Process tree level by level"""
    if not root:
        return []
    
    result = []
    queue = deque([root])
    
    while queue:
        level_size = len(queue)
        current_level = []
        
        for _ in range(level_size):
            node = queue.popleft()
            current_level.append(node.val)
            
            if node.left:
                queue.append(node.left)
            if node.right:
                queue.append(node.right)
        
        result.append(current_level)
    
    return result

# ============================================================================
# CORE PATTERNS - GRAPHS
# ============================================================================

# Graph DFS (with cycle detection)
def dfs_graph(graph: Dict[int, List[int]], start: int) -> bool:
    """Returns True if cycle detected, False otherwise"""
    visited = set()
    rec_stack = set()
    
    def dfs(node):
        visited.add(node)
        rec_stack.add(node)
        
        for neighbor in graph[node]:
            if neighbor not in visited:
                if dfs(neighbor):
                    return True
            elif neighbor in rec_stack:
                return True
        
        rec_stack.remove(node)
        return False
    
    return dfs(start)

# Graph BFS (shortest path in unweighted graph)
def bfs_shortest_path(graph: Dict[int, List[int]], start: int, end: int) -> int:
    """Returns shortest path length, -1 if no path"""
    queue = deque([(start, 0)])  # (node, distance)
    visited = {start}
    
    while queue:
        node, dist = queue.popleft()
        
        if node == end:
            return dist
        
        for neighbor in graph[node]:
            if neighbor not in visited:
                visited.add(neighbor)
                queue.append((neighbor, dist + 1))
    
    return -1

# Topological Sort (Kahn's Algorithm - BFS)
def topological_sort(num_nodes: int, edges: List[List[int]]) -> List[int]:
    """Returns topological order, empty list if cycle exists"""
    in_degree = [0] * num_nodes
    adj = [[] for _ in range(num_nodes)]
    
    for src, dst in edges:
        adj[src].append(dst)
        in_degree[dst] += 1
    
    queue = deque([i for i in range(num_nodes) if in_degree[i] == 0])
    result = []
    
    while queue:
        node = queue.popleft()
        result.append(node)
        
        for neighbor in adj[node]:
            in_degree[neighbor] -= 1
            if in_degree[neighbor] == 0:
                queue.append(neighbor)
    
    return result if len(result) == num_nodes else []

# Union-Find (Disjoint Set Union)
class UnionFind:
    def __init__(self, size: int):
        self.parent = list(range(size))
        self.rank = [0] * size
        self.components = size
    
    def find(self, x: int) -> int:
        """Find root with path compression"""
        if self.parent[x] != x:
            self.parent[x] = self.find(self.parent[x])
        return self.parent[x]
    
    def union(self, x: int, y: int) -> bool:
        """Union by rank, returns True if merge happened"""
        root_x = self.find(x)
        root_y = self.find(y)
        
        if root_x == root_y:
            return False
        
        if self.rank[root_x] < self.rank[root_y]:
            self.parent[root_x] = root_y
        elif self.rank[root_x] > self.rank[root_y]:
            self.parent[root_y] = root_x
        else:
            self.parent[root_y] = root_x
            self.rank[root_x] += 1
        
        self.components -= 1
        return True

# ============================================================================
# CORE PATTERNS - SLIDING WINDOW
# ============================================================================

# Fixed-size sliding window
def max_sum_subarray(arr: List[int], k: int) -> int:
    """Maximum sum of subarray of size k"""
    if len(arr) < k:
        return -1
    
    window_sum = sum(arr[:k])
    max_sum = window_sum
    
    for i in range(k, len(arr)):
        window_sum = window_sum - arr[i - k] + arr[i]
        max_sum = max(max_sum, window_sum)
    
    return max_sum

# Dynamic sliding window (longest substring without repeating)
def length_of_longest_substring(s: str) -> int:
    """Classic sliding window with hash set"""
    char_set = set()
    left = 0
    max_length = 0
    
    for right in range(len(s)):
        while s[right] in char_set:
            char_set.remove(s[left])
            left += 1
        
        char_set.add(s[right])
        max_length = max(max_length, right - left + 1)
    
    return max_length

# ============================================================================
# CORE PATTERNS - TWO POINTERS
# ============================================================================

# Two pointers from ends
def two_sum_sorted(nums: List[int], target: int) -> List[int]:
    """For sorted array, find two numbers that sum to target"""
    left, right = 0, len(nums) - 1
    
    while left < right:
        current_sum = nums[left] + nums[right]
        
        if current_sum == target:
            return [left, right]
        elif current_sum < target:
            left += 1
        else:
            right -= 1
    
    return []

# Fast and slow pointers (cycle detection)
def has_cycle(head: Optional[ListNode]) -> bool:
    """Floyd's cycle detection algorithm"""
    slow = fast = head
    
    while fast and fast.next:
        slow = slow.next
        fast = fast.next.next
        
        if slow == fast:
            return True
    
    return False

# ============================================================================
# CORE PATTERNS - BINARY SEARCH
# ============================================================================

# Standard binary search
def binary_search(arr: List[int], target: int) -> int:
    """Returns index of target, -1 if not found"""
    left, right = 0, len(arr) - 1
    
    while left <= right:
        mid = (left + right) // 2
        
        if arr[mid] == target:
            return mid
        elif arr[mid] < target:
            left = mid + 1
        else:
            right = mid - 1
    
    return -1

# Binary search with bisect module
def binary_search_bisect(arr: List[int], target: int) -> int:
    """Using bisect for cleaner code"""
    idx = bisect.bisect_left(arr, target)
    
    if idx < len(arr) and arr[idx] == target:
        return idx
    return -1

# Search in rotated sorted array
def search_rotated(arr: List[int], target: int) -> int:
    """Binary search on rotated array"""
    left, right = 0, len(arr) - 1
    
    while left <= right:
        mid = (left + right) // 2
        
        if arr[mid] == target:
            return mid
        
        # Determine which side is sorted
        if arr[left] <= arr[mid]:
            if arr[left] <= target < arr[mid]:
                right = mid - 1
            else:
                left = mid + 1
        else:
            if arr[mid] < target <= arr[right]:
                left = mid + 1
            else:
                right = mid - 1
    
    return -1

# ============================================================================
# CORE PATTERNS - HEAPS / PRIORITY QUEUE
# ============================================================================

# Min-heap operations
def kth_largest(nums: List[int], k: int) -> int:
    """Find kth largest element using min-heap"""
    min_heap = nums[:k]
    heapq.heapify(min_heap)
    
    for num in nums[k:]:
        if num > min_heap[0]:
            heapq.heappushpop(min_heap, num)
    
    return min_heap[0]

# Merge K sorted lists
def merge_k_lists(lists: List[Optional[ListNode]]) -> Optional[ListNode]:
    """Using min-heap to merge K sorted linked lists"""
    min_heap = []
    
    # Initialize heap with first node from each list
    for i, lst in enumerate(lists):
        if lst:
            heapq.heappush(min_heap, (lst.val, i, lst))
    
    dummy = ListNode(0)
    current = dummy
    
    while min_heap:
        val, idx, node = heapq.heappop(min_heap)
        current.next = node
        current = current.next
        
        if node.next:
            heapq.heappush(min_heap, (node.next.val, idx, node.next))
    
    return dummy.next

# ============================================================================
# CORE PATTERNS - DYNAMIC PROGRAMMING
# ============================================================================

# 1D DP - Climbing stairs
def climb_stairs(n: int) -> int:
    """Classic 1D DP"""
    if n <= 2:
        return n
    
    dp = [0] * (n + 1)
    dp[1], dp[2] = 1, 2
    
    for i in range(3, n + 1):
        dp[i] = dp[i - 1] + dp[i - 2]
    
    return dp[n]

# 1D DP with space optimization
def climb_stairs_optimized(n: int) -> int:
    """Space-optimized version"""
    if n <= 2:
        return n
    
    prev1, prev2 = 1, 2
    
    for _ in range(3, n + 1):
        curr = prev1 + prev2
        prev1 = prev2
        prev2 = curr
    
    return prev2

# 2D DP - Unique paths
def unique_paths(m: int, n: int) -> int:
    """Grid DP"""
    dp = [[1] * n for _ in range(m)]
    
    for i in range(1, m):
        for j in range(1, n):
            dp[i][j] = dp[i - 1][j] + dp[i][j - 1]
    
    return dp[m - 1][n - 1]

# ============================================================================
# CORE PATTERNS - INTERVALS
# ============================================================================

# Merge intervals
def merge_intervals(intervals: List[List[int]]) -> List[List[int]]:
    """Merge overlapping intervals"""
    if not intervals:
        return []
    
    intervals.sort(key=lambda x: x[0])
    merged = [intervals[0]]
    
    for start, end in intervals[1:]:
        last_end = merged[-1][1]
        
        if start <= last_end:
            merged[-1][1] = max(last_end, end)
        else:
            merged.append([start, end])
    
    return merged

# ============================================================================
# CORE PATTERNS - BACKTRACKING
# ============================================================================

# Permutations
def permute(nums: List[int]) -> List[List[int]]:
    """Generate all permutations"""
    result = []
    
    def backtrack(path):
        if len(path) == len(nums):
            result.append(path[:])
            return
        
        for num in nums:
            if num in path:
                continue
            path.append(num)
            backtrack(path)
            path.pop()
    
    backtrack([])
    return result

# Subsets
def subsets(nums: List[int]) -> List[List[int]]:
    """Generate all subsets"""
    result = []
    
    def backtrack(start, path):
        result.append(path[:])
        
        for i in range(start, len(nums)):
            path.append(nums[i])
            backtrack(i + 1, path)
            path.pop()
    
    backtrack(0, [])
    return result

# ============================================================================
# UTILITY FUNCTIONS
# ============================================================================

# Convert array to linked list
def array_to_list(arr: List[int]) -> Optional[ListNode]:
    """Helper for linked list problems"""
    if not arr:
        return None
    
    dummy = ListNode(0)
    current = dummy
    
    for val in arr:
        current.next = ListNode(val)
        current = current.next
    
    return dummy.next

# Convert linked list to array
def list_to_array(head: Optional[ListNode]) -> List[int]:
    """Helper for linked list problems"""
    result = []
    current = head
    
    while current:
        result.append(current.val)
        current = current.next
    
    return result

# Convert array to binary tree (level-order)
def array_to_tree(arr: List[Optional[int]]) -> Optional[TreeNode]:
    """Helper for tree problems"""
    if not arr or arr[0] is None:
        return None
    
    root = TreeNode(arr[0])
    queue = deque([root])
    i = 1
    
    while queue and i < len(arr):
        node = queue.popleft()
        
        if i < len(arr) and arr[i] is not None:
            node.left = TreeNode(arr[i])
            queue.append(node.left)
        i += 1
        
        if i < len(arr) and arr[i] is not None:
            node.right = TreeNode(arr[i])
            queue.append(node.right)
        i += 1
    
    return root

# ============================================================================
# COMMON PATTERNS SUMMARY
# ============================================================================
"""
SLIDING WINDOW: Use when dealing with contiguous subarrays/substrings
- Fixed size: sum/max/min over window of size k
- Dynamic size: longest/shortest substring with condition

TWO POINTERS: Use for sorted arrays or linked lists
- From ends: two sum, palindrome check
- Fast/slow: cycle detection, middle finding

BINARY SEARCH: Use for sorted arrays or monotonic conditions
- Standard: find target in sorted array
- Rotated: handle array rotation
- Answer space: binary search on answer (koko eating bananas)

HEAP/PRIORITY QUEUE: Use for kth largest/smallest, merge sorted, scheduling
- Min-heap: kth largest, merge k lists
- Max-heap: kth smallest (negate values)

DFS/BFS: Use for trees and graphs
- DFS: path finding, cycle detection, topological sort
- BFS: shortest path, level-order traversal

DYNAMIC PROGRAMMING: Use for optimization with overlapping subproblems
- 1D: climbing stairs, house robber
- 2D: unique paths, longest common subsequence
- Memoization vs tabulation

BACKTRACKING: Use for generating all combinations/permutations/subsets
- Permutations: all arrangements
- Subsets: all combinations
- Constraints: prune invalid paths early

INTERVALS: Use for overlapping ranges
- Merge: combine overlapping intervals
- Insert: add interval and merge
- Non-overlapping: remove minimum to make non-overlapping
"""