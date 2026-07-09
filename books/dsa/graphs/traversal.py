from collections import deque, defaultdict
from typing import Dict, List, Tuple


Graph = Dict[int, List[int]]


# Pattern: adjacency-list construction.
# Invariant: every edge is represented once for directed graphs, twice for undirected graphs.
# Complexity: O(E) time, O(V + E) space.
# Interview line: defaultdict(list) keeps graph construction concise and avoids missing-key checks.
def build_graph(edges: List[Tuple[int, int]], directed: bool = False) -> Graph:
    graph = defaultdict(list)
    for a, b in edges:
        graph[a].append(b)
        if not directed:
            graph[b].append(a)
    return graph


# Pattern: BFS with a queue.
# Invariant: seen contains every node already enqueued, so nodes are processed once.
# Complexity: O(V + E) time, O(V) space.
# Interview line: deque gives O(1) popleft for level-by-level or shortest-hop traversal.
def bfs_order(graph: Graph, start: int) -> List[int]:
    seen = {start}
    q = deque([start])
    order = []

    while q:
        node = q.popleft()
        order.append(node)
        for nei in graph[node]:
            if nei not in seen:
                seen.add(nei)
                q.append(nei)

    return order


# Pattern: recursive DFS.
# Invariant: seen prevents revisiting nodes and recursion explores one path fully.
# Complexity: O(V + E) time, O(V) space for recursion/seen.
# Interview line: DFS is natural when you need reachability, components, or backtracking.
def dfs_order(graph: Graph, start: int) -> List[int]:
    seen = set()
    order = []

    # Pattern: recursive visit helper.
    # Invariant: node is marked seen before any neighbor recursion.
    # Complexity: O(outdegree(node)) local work, O(V) recursion space overall.
    # Interview line: mark before recursing to avoid cycles re-entering the same node.
    def dfs(node: int) -> None:
        seen.add(node)
        order.append(node)
        for nei in graph[node]:
            if nei not in seen:
                dfs(nei)

    dfs(start)
    return order


# Pattern: grid BFS flood fill.
# Invariant: once land is queued, mark it water so it is not counted again.
# Complexity: O(rows * cols) time, O(rows * cols) worst-case space.
# Interview line: every time we discover unvisited land, that starts one island traversal.
def num_islands(grid: List[List[str]]) -> int:
    if not grid or not grid[0]:
        return 0

    rows, cols = len(grid), len(grid[0])
    islands = 0

    for r in range(rows):
        for c in range(cols):
            if grid[r][c] != "1":
                continue
            islands += 1
            grid[r][c] = "0"
            q = deque([(r, c)])

            while q:
                row, col = q.popleft()
                for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                    nr, nc = row + dr, col + dc
                    if 0 <= nr < rows and 0 <= nc < cols and grid[nr][nc] == "1":
                        grid[nr][nc] = "0"
                        q.append((nr, nc))

    return islands


# Pattern: recursive grid DFS flood fill.
# Invariant: each matching cell is consumed before exploring its four neighbors.
# Complexity: O(rows * cols) time, O(rows * cols) recursion space worst case.
# Interview line: mutate visited cells in-place when the prompt allows it.
def flood_fill_dfs(grid: List[List[str]], r: int, c: int, target: str = "1") -> None:
    if not grid or not grid[0]:
        return
    rows, cols = len(grid), len(grid[0])
    if not (0 <= r < rows and 0 <= c < cols) or grid[r][c] != target:
        return

    grid[r][c] = "0"
    flood_fill_dfs(grid, r + 1, c, target)
    flood_fill_dfs(grid, r - 1, c, target)
    flood_fill_dfs(grid, r, c + 1, target)
    flood_fill_dfs(grid, r, c - 1, target)
