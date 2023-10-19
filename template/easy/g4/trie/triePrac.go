package trie

type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {

}

func (t *Trie) Insert(word string) {

}

func (t *Trie) Find(prefix string) []string {

}

func (t *Trie) collectWords(node *TrieNode, prefix string) []string {

}

func (t *Trie) Delete(word string) {

}

func (t *Trie) deleteHelper(node *TrieNode, word string, index int) bool {

}
