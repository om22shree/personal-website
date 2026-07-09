from collections import deque, defaultdict
from typing import List


# Pattern: BFS over wildcard word patterns.
# Invariant: queue distance is the shortest transformation length to that word.
# Complexity: O(N * L^2) time, O(N * L) space for N words of length L.
# Interview line: wildcard buckets turn one-letter transformations into graph neighbors.
def ladder_length(begin_word: str, end_word: str, word_list: List[str]) -> int:
    words = set(word_list)
    if end_word not in words:
        return 0

    patterns = defaultdict(list)
    for word in words | {begin_word}:
        for i in range(len(word)):
            patterns[word[:i] + "*" + word[i + 1:]].append(word)

    q = deque([(begin_word, 1)])
    seen = {begin_word}

    while q:
        word, distance = q.popleft()
        if word == end_word:
            return distance
        for i in range(len(word)):
            pattern = word[:i] + "*" + word[i + 1:]
            for nxt in patterns[pattern]:
                if nxt not in seen:
                    seen.add(nxt)
                    q.append((nxt, distance + 1))
            patterns[pattern].clear()

    return 0
