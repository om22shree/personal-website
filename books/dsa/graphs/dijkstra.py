from collections import defaultdict
import heapq
from typing import List, Tuple


# Pattern: Dijkstra with adjacency list and min-heap.
# Invariant: the first time a node is popped, its shortest distance is finalized.
# Complexity: O((V + E) log V) time, O(V + E) space.
# Interview line: use Dijkstra when edge weights are non-negative and you need shortest paths.
def network_delay_time(times: List[List[int]], n: int, k: int) -> int:
    graph = defaultdict(list)
    for src, dst, weight in times:
        graph[src].append((dst, weight))

    dist = {}
    heap = [(0, k)]

    while heap:
        cost, node = heapq.heappop(heap)
        if node in dist:
            continue
        dist[node] = cost
        for nei, weight in graph[node]:
            if nei not in dist:
                heapq.heappush(heap, (cost + weight, nei))

    return max(dist.values()) if len(dist) == n else -1


# Pattern: reusable single-source shortest path.
# Invariant: dist[node] is the best known cost, and stale heap entries are skipped.
# Complexity: O((V + E) log V) time, O(V + E) space.
# Interview line: pushing improved distances is simpler than decrease-key in Python heaps.
def shortest_path(n: int, edges: List[Tuple[int, int, int]], start: int) -> List[float]:
    graph = defaultdict(list)
    for a, b, weight in edges:
        graph[a].append((b, weight))

    dist = [float("inf")] * n
    dist[start] = 0
    heap = [(0, start)]

    while heap:
        cost, node = heapq.heappop(heap)
        if cost > dist[node]:
            continue
        for nei, weight in graph[node]:
            new_cost = cost + weight
            if new_cost < dist[nei]:
                dist[nei] = new_cost
                heapq.heappush(heap, (new_cost, nei))

    return dist
