// Package bytetrie provides implementation of a trie with sequence of bytes as
// a key
//
// There are several trie implementations around. This one is very simple and only
// supports trie where keys are byte sequences. Also, it has a Match() function
// which is missing in other trie implementation. The intention is to use it to
// match file type by initial bytes in the file.
// The library is inspired by Trie code developed by Drew Noakes for the
// https://github.com/drewnoakes/metadata-extractor library.

package bytetrie

import "fmt"

// Trie node
type node struct {
	value    interface{}
	hasValue bool
	children map[byte]*node
}

// Trie is the top-level description of a trie
type Trie struct {
	root     *node
	maxDepth int
}

// New returns a new instance of a trie
func New() *Trie {
	trie := &Trie{}
	trie.Init()
	return trie
}

// Init initializes a trie (in case it wasn't created with New() )
func (trie *Trie) Init() {
	if trie.root == nil {
		trie.root = &node{children: make(map[byte]*node)}
	}
}

// MaxDepth returns the length of the maximum tree path
func (trie *Trie) MaxDepth() int {
	return trie.maxDepth
}

// Insert given value using concatenation of all keys as one long key
func (trie *Trie) Insert(value interface{}, keys ...[]byte) {
	key := make([]byte, 0, 16)
	for _, k := range keys {
		key = append(key, k...)
	}
	trie.Init()
	currentNode := trie.root
	depth := 0
	for _, b := range key {
		child, ok := currentNode.children[b]
		if !ok {
			child = &node{children: make(map[byte]*node)}
			currentNode.children[b] = child
		}
		currentNode = child
		depth++
	}
	currentNode.value = value
	currentNode.hasValue = true
	if depth > trie.maxDepth {
		trie.maxDepth = depth
	}
}

// Recursive trie walk, apply cb function to each node that has a value.
func walk(n *node, currKey []byte, cb func(key []byte, value interface{})) {
	if n.hasValue {
		cb(currKey, n.value)
	}
	for b, child := range n.children {
		walk(child, append(currKey, b), cb)
	}
}

// Do walks the trie and calls callback for each element that has a value.
func (trie *Trie) Do(cb func(key []byte, value interface{})) {
	trie.Init()
	walk(trie.root, make([]byte, 0, trie.maxDepth), cb)
}

// PrintKeys prints all keys in the trie as strings
func (trie *Trie) PrintKeys() {
	if trie == nil || trie.root == nil {
		return
	}
	printer := func(key []byte, _ interface{}) {
		fmt.Println(string(key))
	}
	trie.Do(printer)
}

// Get returns the value for a key. Also returns true iff found
func (trie *Trie) Get(key []byte) (interface{}, bool) {
	if trie == nil || trie.root == nil {
		return nil, false
	}
	current := trie.root
	for _, b := range key {
		child, ok := current.children[b]
		if !ok {
			return nil, false
		}
		current = child
	}
	return current.value, current.hasValue
}

// Match scans input sequence trying to find the longest subsequent that is
// present in the input trie. If found, the matching value is returned.
// Also returns true iff the value is found.
func (trie *Trie) Match(sequence []byte) (interface{}, bool) {
	if trie == nil || trie.root == nil {
		return nil, false
	}
	current := trie.root
	for _, b := range sequence {
		child, ok := current.children[b]
		if !ok {
			break
		}
		current = child
	}
	return current.value, current.hasValue
}
