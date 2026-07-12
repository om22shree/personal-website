package graphs

// Pattern: BFS over wildcard word patterns.
// Invariant: queue distance is the shortest transformation length to that word.
// Complexity: O(N * L^2) time, O(N * L) space for N words of length L.
// Interview line: wildcard buckets turn one-letter transformations into graph neighbors.
func LadderLength(beginWord string, endWord string, wordList []string) int {
	wordSet := make(map[string]bool)
	for _, w := range wordList {
		wordSet[w] = true
	}

	if !wordSet[endWord] {
		return 0
	}

	// Group words by their wildcard patterns
	// e.g. "hot" -> "*ot", "h*t", "ho*"
	wildcards := make(map[string][]string)
	L := len(beginWord)

	for _, w := range wordList {
		for i := 0; i < L; i++ {
			pattern := w[:i] + "*" + w[i+1:]
			wildcards[pattern] = append(wildcards[pattern], w)
		}
	}

	type queueNode struct {
		word  string
		level int
	}

	q := []queueNode{{word: beginWord, level: 1}}
	visited := make(map[string]bool)
	visited[beginWord] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		word, level := curr.word, curr.level
		if word == endWord {
			return level
		}

		for i := 0; i < L; i++ {
			pattern := word[:i] + "*" + word[i+1:]
			for _, neighbor := range wildcards[pattern] {
				if !visited[neighbor] {
					visited[neighbor] = true
					q = append(q, queueNode{word: neighbor, level: level + 1})
				}
			}
		}
	}

	return 0
}
