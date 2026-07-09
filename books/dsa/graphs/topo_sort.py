from collections import deque, defaultdict
from typing import List, Tuple


# Pattern: Kahn topological sort.
# Invariant: queue contains nodes with zero remaining prerequisites.
# Complexity: O(V + E) time, O(V + E) space.
# Interview line: if we cannot consume every course, a dependency cycle remains.
def can_finish(num_courses: int, prerequisites: List[List[int]]) -> bool:
    graph = defaultdict(list)
    indegree = [0] * num_courses

    for course, prereq in prerequisites:
        graph[prereq].append(course)
        indegree[course] += 1

    q = deque(i for i, degree in enumerate(indegree) if degree == 0)
    taken = 0

    while q:
        node = q.popleft()
        taken += 1
        for nei in graph[node]:
            indegree[nei] -= 1
            if indegree[nei] == 0:
                q.append(nei)

    return taken == num_courses


# Pattern: topological ordering with indegrees.
# Invariant: append a node only after all prerequisites have been consumed.
# Complexity: O(V + E) time, O(V + E) space.
# Interview line: the order exists only if the topo process visits every node.
def find_course_order(num_courses: int, prerequisites: List[List[int]]) -> List[int]:
    graph = defaultdict(list)
    indegree = [0] * num_courses

    for course, prereq in prerequisites:
        graph[prereq].append(course)
        indegree[course] += 1

    q = deque(i for i, degree in enumerate(indegree) if degree == 0)
    order = []

    while q:
        node = q.popleft()
        order.append(node)
        for nei in graph[node]:
            indegree[nei] -= 1
            if indegree[nei] == 0:
                q.append(nei)

    return order if len(order) == num_courses else []


# Pattern: DFS color marking for directed cycle detection.
# Invariant: color 1 means currently on recursion stack, color 2 means fully processed.
# Complexity: O(V + E) time, O(V + E) space.
# Interview line: seeing a gray node again means we found a back edge and therefore a cycle.
def has_cycle_dfs(num_nodes: int, edges: List[Tuple[int, int]]) -> bool:
    graph = defaultdict(list)
    for a, b in edges:
        graph[a].append(b)

    color = [0] * num_nodes

    # Pattern: recursion-stack cycle helper.
    # Invariant: gray nodes are active ancestors in the current DFS path.
    # Complexity: O(outdegree(node)) local work, O(V) recursion stack overall.
    # Interview line: a gray neighbor means this path loops back to itself.
    def dfs(node: int) -> bool:
        color[node] = 1
        for nei in graph[node]:
            if color[nei] == 1:
                return True
            if color[nei] == 0 and dfs(nei):
                return True
        color[node] = 2
        return False

    return any(color[node] == 0 and dfs(node) for node in range(num_nodes))
