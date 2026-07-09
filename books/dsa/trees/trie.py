from collections import defaultdict
from functools import lru_cache
from typing import Dict, List, Optional


class TrieNode:
    # Pattern: trie node with lazy children.
    # Invariant: children maps characters to child TrieNodes, is_word marks terminals.
    # Complexity: O(1) time, O(1) initial space.
    # Interview line: defaultdict lets insertion create missing child nodes on demand.
    def __init__(self):
        self.children: Dict[str, "TrieNode"] = defaultdict(TrieNode)
        self.is_word = False


class Trie:
    # Pattern: trie root initialization.
    # Invariant: root is an empty node shared by all prefixes.
    # Complexity: O(1) time, O(1) space.
    # Interview line: every word starts by walking from the same root.
    def __init__(self):
        self.root = TrieNode()

    # Pattern: prefix-tree insertion.
    # Invariant: after processing each character, node represents that prefix.
    # Complexity: O(L) time, O(L) space for new nodes.
    # Interview line: defaultdict creates child nodes lazily as the word is inserted.
    def insert(self, word: str) -> None:
        node = self.root
        for ch in word:
            node = node.children[ch]
        node.is_word = True

    # Pattern: exact trie lookup.
    # Invariant: word is present only if traversal succeeds and final node is marked.
    # Complexity: O(L) time, O(1) extra space.
    # Interview line: prefix existence is not enough; check the terminal word flag.
    def search(self, word: str) -> bool:
        node = self._find(word)
        return bool(node and node.is_word)

    # Pattern: prefix trie lookup.
    # Invariant: prefix exists if every character edge can be followed.
    # Complexity: O(L) time, O(1) extra space.
    # Interview line: service discovery and config keys often need prefix matching.
    def starts_with(self, prefix: str) -> bool:
        return self._find(prefix) is not None

    # Pattern: shared trie traversal helper.
    # Invariant: returns the node for the prefix or None at the first missing edge.
    # Complexity: O(L) time, O(1) extra space.
    # Interview line: factor traversal so exact search and prefix search stay simple.
    def _find(self, prefix: str) -> Optional[TrieNode]:
        node = self.root
        for ch in prefix:
            if ch not in node.children:
                return None
            node = node.children[ch]
        return node


# Pattern: memoized DFS over string index.
# Invariant: can_break(i) means s[i:] can be segmented into dictionary words.
# Complexity: O(n * max_word_len) states/checks, O(n) memo space.
# Interview line: cache the start index so repeated suffix checks are computed once.
def word_break(s: str, word_dict: List[str]) -> bool:
    words = set(word_dict)
    max_len = max(map(len, words), default=0)

    # Pattern: memoized suffix helper.
    # Invariant: can_break(i) answers whether the suffix starting at i is segmentable.
    # Complexity: O(max_word_len) transitions per index, O(n) memo states.
    # Interview line: cache by index because the same suffix appears through many split paths.
    @lru_cache(None)
    def can_break(i: int) -> bool:
        if i == len(s):
            return True
        for j in range(i + 1, min(len(s), i + max_len) + 1):
            if s[i:j] in words and can_break(j):
                return True
        return False

    return can_break(0)
