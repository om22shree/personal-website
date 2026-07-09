from typing import List, Optional


class UnionFind:
    # Pattern: disjoint-set initialization.
    # Invariant: every node starts as its own parent/component.
    # Complexity: O(n) time, O(n) space.
    # Interview line: DSU gives near-constant connectivity checks after setup.
    def __init__(self, n: int):
        self.parent = list(range(n))
        self.rank = [0] * n
        self.count = n

    # Pattern: find with path compression.
    # Invariant: returned value is the representative root for x's component.
    # Complexity: near O(1) amortized time, O(1) extra space.
    # Interview line: path compression flattens future finds while preserving component identity.
    def find(self, x: int) -> int:
        while x != self.parent[x]:
            self.parent[x] = self.parent[self.parent[x]]
            x = self.parent[x]
        return x

    # Pattern: union by rank.
    # Invariant: return False only when a and b already share a component.
    # Complexity: near O(1) amortized time, O(1) extra space.
    # Interview line: failed union is the signal for an undirected cycle.
    def union(self, a: int, b: int) -> bool:
        root_a, root_b = self.find(a), self.find(b)
        if root_a == root_b:
            return False
        if self.rank[root_a] < self.rank[root_b]:
            root_a, root_b = root_b, root_a
        self.parent[root_b] = root_a
        if self.rank[root_a] == self.rank[root_b]:
            self.rank[root_a] += 1
        self.count -= 1
        return True


# Pattern: DSU cycle detection in an undirected graph.
# Invariant: each successful union connects two previously separate components.
# Complexity: O(E alpha(V)) time, O(V) space.
# Interview line: the first edge whose endpoints are already connected is redundant.
def redundant_connection(edges: List[List[int]]) -> Optional[List[int]]:
    uf = UnionFind(len(edges) + 1)
    for a, b in edges:
        if not uf.union(a, b):
            return [a, b]
    return None


# Pattern: DSU component counting.
# Invariant: count decreases only when a union merges two different components.
# Complexity: O(E alpha(V)) time, O(V) space.
# Interview line: union every edge, then the DSU count is the number of connected components.
def count_components(n: int, edges: List[List[int]]) -> int:
    uf = UnionFind(n)
    for a, b in edges:
        uf.union(a, b)
    return uf.count
