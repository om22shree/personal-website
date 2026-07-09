import heapq
from collections import defaultdict, deque
from typing import Dict, List, Optional, Set, Tuple


class Node:
    # Pattern: graph node container.
    # Invariant: neighbors stores adjacent Node references.
    # Complexity: O(1) time, O(1) space.
    # Interview line: graph cloning needs identity mapping, not just values.
    def __init__(self, val: int = 0, neighbors: Optional[List["Node"]] = None):
        self.val = val
        self.neighbors = neighbors if neighbors is not None else []


# Pattern: DFS clone with visited map.
# Invariant: clones maps every visited original node to its copy.
# Complexity: O(V + E) time, O(V) space.
# Interview line: create the copy before cloning neighbors to break cycles.
def clone_graph(node: Optional[Node]) -> Optional[Node]:
    clones: Dict[Node, Node] = {}

    def dfs(cur: Optional[Node]) -> Optional[Node]:
        if not cur:
            return None
        if cur in clones:
            return clones[cur]
        copy = Node(cur.val)
        clones[cur] = copy
        copy.neighbors = [dfs(nei) for nei in cur.neighbors if nei]
        return copy

    return dfs(node)


# Pattern: topological sort over character precedence.
# Invariant: indegree zero characters have all prerequisites satisfied.
# Complexity: O(total chars + edges) time, O(unique chars) space.
# Interview line: the first differing character between adjacent words creates the ordering edge.
def alien_order(words: List[str]) -> str:
    graph = {ch: set() for word in words for ch in word}
    indegree = {ch: 0 for ch in graph}

    for first, second in zip(words, words[1:]):
        if len(first) > len(second) and first.startswith(second):
            return ""
        for a, b in zip(first, second):
            if a != b:
                if b not in graph[a]:
                    graph[a].add(b)
                    indegree[b] += 1
                break

    q = deque([ch for ch, deg in indegree.items() if deg == 0])
    order = []
    while q:
        ch = q.popleft()
        order.append(ch)
        for nei in graph[ch]:
            indegree[nei] -= 1
            if indegree[nei] == 0:
                q.append(nei)

    return "".join(order) if len(order) == len(graph) else ""


# Pattern: Prim minimum spanning tree.
# Invariant: heap stores cheapest edges from visited points to unvisited points.
# Complexity: O(n^2 log n) time, O(n^2) heap space.
# Interview line: grow the connected set by always taking the cheapest crossing edge.
def min_cost_connect_points(points: List[List[int]]) -> int:
    n = len(points)
    visited: Set[int] = set()
    heap: List[Tuple[int, int]] = [(0, 0)]
    total = 0

    while len(visited) < n:
        cost, i = heapq.heappop(heap)
        if i in visited:
            continue
        visited.add(i)
        total += cost
        x1, y1 = points[i]
        for j in range(n):
            if j in visited:
                continue
            x2, y2 = points[j]
            heapq.heappush(heap, (abs(x1 - x2) + abs(y1 - y2), j))

    return total


# Pattern: Dijkstra over grid states.
# Invariant: first time bottom-right is popped, its effort is minimal.
# Complexity: O(mn log mn) time, O(mn) space.
# Interview line: path cost is the maximum cell height seen, so prioritize lower current effort.
def swim_in_water(grid: List[List[int]]) -> int:
    n = len(grid)
    heap = [(grid[0][0], 0, 0)]
    seen = {(0, 0)}
    dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]

    while heap:
        time, r, c = heapq.heappop(heap)
        if r == n - 1 and c == n - 1:
            return time
        for dr, dc in dirs:
            nr, nc = r + dr, c + dc
            if 0 <= nr < n and 0 <= nc < n and (nr, nc) not in seen:
                seen.add((nr, nc))
                heapq.heappush(heap, (max(time, grid[nr][nc]), nr, nc))

    return -1


# Pattern: multi-source boundary reachability.
# Invariant: a cell belongs to an ocean set if water can flow from it to that ocean.
# Complexity: O(mn) time, O(mn) space.
# Interview line: reverse the flow and start DFS from ocean borders.
def pacific_atlantic(heights: List[List[int]]) -> List[List[int]]:
    if not heights:
        return []
    rows, cols = len(heights), len(heights[0])

    def flow(starts: List[Tuple[int, int]]) -> Set[Tuple[int, int]]:
        seen = set(starts)
        stack = starts[:]
        while stack:
            r, c = stack.pop()
            for dr, dc in ((1, 0), (-1, 0), (0, 1), (0, -1)):
                nr, nc = r + dr, c + dc
                if (
                    0 <= nr < rows
                    and 0 <= nc < cols
                    and (nr, nc) not in seen
                    and heights[nr][nc] >= heights[r][c]
                ):
                    seen.add((nr, nc))
                    stack.append((nr, nc))
        return seen

    pacific = flow([(0, c) for c in range(cols)] + [(r, 0) for r in range(rows)])
    atlantic = flow(
        [(rows - 1, c) for c in range(cols)] + [(r, cols - 1) for r in range(rows)]
    )
    return [[r, c] for r, c in sorted(pacific & atlantic)]


# Pattern: BFS coloring.
# Invariant: every colored edge connects opposite colors.
# Complexity: O(V + E) time, O(V) space.
# Interview line: a graph is bipartite if every component can be two-colored.
def is_bipartite(graph: List[List[int]]) -> bool:
    color = {}
    for start in range(len(graph)):
        if start in color:
            continue
        color[start] = 0
        q = deque([start])
        while q:
            node = q.popleft()
            for nei in graph[node]:
                if nei in color:
                    if color[nei] == color[node]:
                        return False
                else:
                    color[nei] = 1 - color[node]
                    q.append(nei)
    return True
