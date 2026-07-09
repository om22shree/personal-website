# DSA Interview Pattern Command Center

Use this as the first stop before practice. Pick a file, solve one must-solve problem, then say the recall line out loud before coding.

| File | Pattern | Must-solve problems | Recall line |
|---|---|---|---|
| `common-imports.py` | Python interview toolkit | Rehearse `deque`, `defaultdict`, `Counter`, `OrderedDict`, `heapq`, `bisect`, `lru_cache` | Know which library matches the smell of the problem. |
| `arrays_hashing.py` | Hashmaps, sets, prefix sums, encoding | Two Sum, Group Anagrams, Product Except Self, Valid Sudoku, Longest Consecutive, Subarray Sum Equals K | Choose state that makes membership, frequency, or prefix math O(1). |
| `sliding_window.py` | Fixed and variable windows | Longest Substring Without Repeat, Minimum Window Substring, Sliding Window Maximum, Find All Anagrams | Expand until useful, shrink until valid or minimal. |
| `two_pointers.py` | Directional pointer movement | Two Sum II, 3Sum, Container With Most Water, Trapping Rain Water, Sort Colors | Sorted/order structure tells which pointer can move. |
| `stack.py` | Stack and monotonic stack | Valid Parentheses, Min Stack, RPN, Daily Temperatures, Car Fleet, Largest Rectangle | A stack remembers unresolved work in last-in-first-out order. |
| `binary_search.py` | Classic search and search-on-answer | Binary Search, Search Rotated Array, Koko Eating Bananas, Ship Packages Within D Days | Define monotonic feasibility, then search the boundary. |
| `intervals.py` | Sorted ranges and overlap logic | Merge Intervals, Insert Interval, Meeting Rooms II, Interval List Intersections | Sort first; only neighboring or active intervals can conflict. |
| `heaps.py` | Priority queue patterns | Kth Largest, Top K Frequent, Merge K Sorted Lists, Task Scheduler, Median from Data Stream | Use a heap when the next best item matters repeatedly. |
| `backtracking.py` | Choice trees and undo | Subsets, Combination Sum, Permutations, Word Search, Palindrome Partitioning | Make a choice, recurse, then undo before the next choice. |
| `dp_lite.py` | Small memoization and 1D DP | Climbing Stairs, House Robber, Coin Change, Word Break, LIS | Name the state, cache it, and only keep the history you need. |
| `dp_2d.py` | Table DP and grid/string state | LCS, Edit Distance, Unique Paths II, Min Path Sum, Palindromic Substrings | Define what each cell means before writing transitions. |
| `greedy.py` | Local choices with carried state | Jump Game, Gas Station, Partition Labels, Valid Parenthesis String, Hand of Straights | Track the smallest state that proves the local choice stays valid. |
| `math_geometry.py` | Matrix geometry and numeric simulation | Rotate Image, Spiral Matrix, Set Matrix Zeroes, Pow(x,n), Happy Number | Translate the geometry or digits into a stable invariant. |
| `bit_manipulation.py` | Bit masks, XOR, shifts | Single Number, Number of 1 Bits, Counting Bits, Reverse Bits, Sum of Two Integers | XOR cancels, shifts move place value, masks isolate facts. |
| `rate_limiter.py` | Infra-flavored rolling controls | Sliding-window limiter, Token Bucket, Fixed Window, Health-check Window | Decide what state survives between requests and expire stale state. |
| `graphs/traversal.py` | BFS/DFS reachability | Number of Islands, Clone Graph, Flood Fill | Mark visited when enqueuing or entering, not after wandering around. |
| `graphs/topo_sort.py` | Dependencies and cycles | Course Schedule, Course Schedule II, Detect Cycle in Directed Graph | Zero indegree means ready; gray node means cycle. |
| `graphs/union_find.py` | Connectivity and undirected cycles | Redundant Connection, Number of Components, Accounts Merge | Failed union means the two nodes were already connected. |
| `graphs/dijkstra.py` | Weighted shortest path | Network Delay Time, Path With Minimum Effort | First heap pop finalizes distance when weights are non-negative. |
| `graphs/advanced.py` | Clone, ordering, MST, grid Dijkstra | Clone Graph, Alien Dictionary, Min Cost Connect Points, Swim in Water, Pacific Atlantic | Model the hidden graph, then pick BFS, topo, heap, or MST by edge meaning. |
| `graphs/matrix_bfs.py` | Multi-source grid BFS | 01 Matrix, Rotting Oranges, Walls and Gates | Start from all sources at once and let distance/time spread outward. |
| `graphs/word_ladder.py` | BFS over implicit graph | Word Ladder | Wildcard buckets make one-letter transitions cheap to find. |
| `graphs/bellman_ford.py` | Bounded edge relaxation | Cheapest Flights Within K Stops | Copy the previous round so each pass adds only one edge. |
| `heaps/top_k.py` | Top-k selection | Kth Largest, Top K Frequent, K Closest Points | Keep only the k candidates that can still matter. |
| `heaps/merge_k.py` | K-way sorted merge | Merge K Sorted Lists, Merge Sorted Arrays, Merge Sorted Logs | Emit the smallest current head, then advance that source. |
| `heaps/scheduling.py` | Scheduling and streaming median | Task Scheduler, Median from Data Stream | Priority handles urgency; two heaps balance the middle. |
| `linkedlists/basics.py` | Pointer fundamentals | Reverse Linked List, Linked List Cycle, Middle of Linked List, Palindrome Linked List | Save `next` before rewiring anything. |
| `linkedlists/merge.py` | Dummy node and list weaving | Merge Two Lists, Merge K Lists, Remove Nth From End, Reorder List | Dummy heads and fast/slow pointers remove edge-case pain. |
| `linkedlists/cache.py` | LRU cache design | LRU Cache with `OrderedDict`, LRU Cache with DLL + hashmap | Hashmap finds nodes; order structure tracks recency. |
| `trees/traversal.py` | Tree DFS/BFS basics | Invert Tree, Same Tree, Level Order, Right Side View, Build Tree | Decide whether node work happens before, between, or after children. |
| `trees/bst.py` | Ordered tree constraints | Validate BST, Kth Smallest, LCA in BST | Carry bounds; inorder of BST is sorted. |
| `trees/dfs_properties.py` | Postorder tree properties | Max Depth, Diameter, Balanced Tree, Subtree, Good Nodes, Binary Tree LCA, Max Path Sum | Return one value upward, update global/local answer at the node. |
| `trees/serialization.py` | Preserve tree shape | Serialize/Deserialize Binary Tree | Null markers preserve structure, not just values. |
| `trees/trie.py` | Prefix tree and memoized word split | Implement Trie, Word Break, Prefix Search | Walk characters as edges; terminal flags distinguish words from prefixes. |

## Practice Loop

1. Pick one row.
2. Read the recall line.
3. Re-type the core function from memory.
4. Run a tiny test.
5. Say pattern, invariant, complexity, and interview line out loud.
