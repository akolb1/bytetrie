package main

import (
	"fmt"

	"github.com/akolb1/bytetrie"
)

func main() {
	trie := bytetrie.New()
	trie.Init()
	trie.Insert(true, []byte("hello"))
	trie.Insert(true, []byte("world"))
	trie.Insert(true, []byte("help"))
	trie.Insert(true, []byte("work"))
	trie.PrintKeys()
	fmt.Println("max depth is", trie.MaxDepth())
}
