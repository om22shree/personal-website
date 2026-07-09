from collections import OrderedDict
from typing import Dict, Optional


class LRUCache:
    # Pattern: OrderedDict-backed LRU cache.
    # Invariant: oldest key is at the front, most recent key is at the end.
    # Complexity: O(1) setup, O(capacity) space.
    # Interview line: OrderedDict exposes move_to_end and popitem(last=False) for LRU behavior.
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.cache = OrderedDict()

    # Pattern: lookup plus recency update.
    # Invariant: a successful get moves the key to most-recent position.
    # Complexity: O(1) time, O(1) extra space.
    # Interview line: reads count as use, so they must update recency.
    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        self.cache.move_to_end(key)
        return self.cache[key]

    # Pattern: insert/update plus evict least recent.
    # Invariant: cache size never exceeds capacity after put returns.
    # Complexity: O(1) time, O(1) extra space.
    # Interview line: after insertion, pop the front item if capacity is exceeded.
    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            self.cache.move_to_end(key)
        self.cache[key] = value
        if len(self.cache) > self.capacity:
            self.cache.popitem(last=False)


class ManualLRUCache:
    class Node:
        # Pattern: doubly-linked cache node.
        # Invariant: each real node stores key, value, and prev/next links.
        # Complexity: O(1) time, O(1) space per node.
        # Interview line: the key is stored so eviction can delete the hashmap entry.
        def __init__(self, key: int = 0, value: int = 0):
            self.key = key
            self.value = value
            self.prev: Optional["ManualLRUCache.Node"] = None
            self.next: Optional["ManualLRUCache.Node"] = None

    # Pattern: hashmap plus sentinel doubly-linked list.
    # Invariant: left.next is least recent and right.prev is most recent.
    # Complexity: O(1) setup, O(capacity) space.
    # Interview line: sentinels remove edge cases for empty and one-node lists.
    def __init__(self, capacity: int):
        self.capacity = capacity
        self.cache: Dict[int, ManualLRUCache.Node] = {}
        self.left = self.Node()
        self.right = self.Node()
        self.left.next = self.right
        self.right.prev = self.left

    # Pattern: unlink a node from the doubly-linked list.
    # Invariant: neighboring nodes point to each other after removal.
    # Complexity: O(1) time, O(1) extra space.
    # Interview line: O(1) removal is why LRU uses a doubly-linked list.
    def _remove(self, node: "ManualLRUCache.Node") -> None:
        prev, nxt = node.prev, node.next
        prev.next = nxt
        nxt.prev = prev

    # Pattern: insert node at most-recent end.
    # Invariant: node becomes right.prev after insertion.
    # Complexity: O(1) time, O(1) extra space.
    # Interview line: every get or put refreshes a node by moving it to the MRU side.
    def _insert_mru(self, node: "ManualLRUCache.Node") -> None:
        prev, nxt = self.right.prev, self.right
        prev.next = nxt.prev = node
        node.prev, node.next = prev, nxt

    # Pattern: hashmap lookup plus list refresh.
    # Invariant: successful lookup moves the node to most-recent position.
    # Complexity: O(1) time, O(1) extra space.
    # Interview line: hashmap finds the node, linked list updates recency.
    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        node = self.cache[key]
        self._remove(node)
        self._insert_mru(node)
        return node.value

    # Pattern: upsert then evict LRU if needed.
    # Invariant: hashmap and linked list contain the same real nodes.
    # Complexity: O(1) time, O(1) extra space.
    # Interview line: evict left.next because it is the least recently used node.
    def put(self, key: int, value: int) -> None:
        if key in self.cache:
            self._remove(self.cache[key])
        node = self.Node(key, value)
        self.cache[key] = node
        self._insert_mru(node)

        if len(self.cache) > self.capacity:
            lru = self.left.next
            self._remove(lru)
            del self.cache[lru.key]
