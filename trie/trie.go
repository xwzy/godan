package trie

import "sync"

type trieNode struct {
	nodes map[rune]*trieNode
	isEnd bool
}

type Trie struct {
	root *trieNode
	mu   sync.Mutex
}

func DefaultTrie() *Trie {
	return NewTrie()
}

func NewTrie() *Trie {
	return &Trie{root: &trieNode{
		nodes: make(map[rune]*trieNode, 0),
		isEnd: false,
	}}
}

func (t *Trie) Insert(word string) {
	if t.root == nil {
		t.root = &trieNode{
			nodes: make(map[rune]*trieNode, 0),
			isEnd: false,
		}
	}

	charList := []rune(word)
	ptr := t.root
	for idx, item := range charList {
		if _, ok := ptr.nodes[item]; !ok {
			t.mu.Lock()
			ptr.nodes[item] = &trieNode{
				nodes: make(map[rune]*trieNode, 0),
				isEnd: false,
			}
			t.mu.Unlock()
		}
		ptr = ptr.nodes[item]
		if idx == len(charList)-1 {
			ptr.isEnd = true
		}
	}
}

func (t *Trie) Search(word string) bool {
	ptr := t.root
	charList := []rune(word)

	for idx, item := range charList {
		if ptr == nil {
			return false
		}
		if _, ok := ptr.nodes[item]; ok {
			ptr = ptr.nodes[item]
		} else {
			return false
		}
		if idx == len(charList)-1 {
			return ptr.isEnd
		}
	}
	return false
}

func (t *Trie) StartsWith(word string) bool {
	ptr := t.root
	charList := []rune(word)

	for _, item := range charList {
		if ptr == nil {
			return false
		}
		if _, ok := ptr.nodes[item]; ok {
			ptr = ptr.nodes[item]
		} else {
			return false
		}
	}
	return true
}
