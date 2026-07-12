package trees

// TrieNode definition
type TrieNode struct {
	Children map[rune]*TrieNode
	IsWord   bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		IsWord:   false,
	}
}

// Trie definition
type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

// Pattern: prefix-tree insertion.
// Invariant: after processing each character, node represents that prefix.
// Complexity: O(L) time, O(L) space for new nodes.
// Interview line: create child nodes lazily as the word is inserted.
func (this *Trie) Insert(word string) {
	node := this.Root
	for _, ch := range word {
		if _, ok := node.Children[ch]; !ok {
			node.Children[ch] = NewTrieNode()
		}
		node = node.Children[ch]
	}
	node.IsWord = true
}

// Pattern: shared trie traversal helper.
// Invariant: returns the node for the prefix or nil at the first missing edge.
// Complexity: O(L) time, O(1) extra space.
// Interview line: factor traversal so exact search and prefix search stay simple.
func (this *Trie) Find(prefix string) *TrieNode {
	node := this.Root
	for _, ch := range prefix {
		if _, ok := node.Children[ch]; !ok {
			return nil
		}
		node = node.Children[ch]
	}
	return node
}

// Pattern: exact trie lookup.
// Invariant: word is present only if traversal succeeds and final node is marked.
// Complexity: O(L) time, O(1) extra space.
// Interview line: prefix existence is not enough; check the terminal word flag.
func (this *Trie) Search(word string) bool {
	node := this.Find(word)
	return node != nil && node.IsWord
}

// Pattern: prefix trie lookup.
// Invariant: prefix exists if every character edge can be followed.
// Complexity: O(L) time, O(1) extra space.
// Interview line: service discovery and config keys often need prefix matching.
func (this *Trie) StartsWith(prefix string) bool {
	return this.Find(prefix) != nil
}

// Pattern: memoized DFS over string index.
// Invariant: can_break(i) means s[i:] can be segmented into dictionary words.
// Complexity: O(n * max_word_len) states/checks, O(n) memo space.
// Interview line: cache the start index so repeated suffix checks are computed once.
func TrieWordBreak(s string, wordDict []string) bool {
	words := make(map[string]bool)
	maxLen := 0
	for _, w := range wordDict {
		words[w] = true
		if len(w) > maxLen {
			maxLen = len(w)
		}
	}

	memo := make(map[int]bool)
	var canBreak func(i int) bool
	canBreak = func(i int) bool {
		if i == len(s) {
			return true
		}
		if val, ok := memo[i]; ok {
			return val
		}

		limit := i + maxLen
		if len(s) < limit {
			limit = len(s)
		}

		for j := i + 1; j <= limit; j++ {
			if words[s[i:j]] && canBreak(j) {
				memo[i] = true
				return true
			}
		}

		memo[i] = false
		return false
	}

	return canBreak(0)
}
