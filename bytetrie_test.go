package bytetrie

import (
	"testing"

	"fmt"
	"github.com/akolb1/filetypes/bytetrie"
)

// TestEmpty checks properties of an empty tree
func TestEmpty(t *testing.T) {
	t.Parallel()
	trie := bytetrie.New()
	depth := trie.MaxDepth()
	if depth != 0 {
		t.Error("MaxDepth() is non-zero for an empty trie")
	}
	result, ok := trie.Get([]byte{0})
	if result != nil || ok {
		t.Error("Empty trie has a value")
	}
}

func TestInsert(t *testing.T) {
	// Various inputs that are inserted in a tree
	t.Parallel()
	trie := bytetrie.New()

	inputs := map[string]int{
		"hello":         1,
		"world":         2,
		"help":          3,
		"work":          4,
		"SomeLongWord":  5,
		"SomeOtherWord": 6,
	}

	// Something not in a trie
	badword := "What?"

	// Values in a tree that were visited during the walk
	visited := map[string]bool{}
	// Initialize visited to all false and insert all inputs into a tree
	longest := 0
	for k, v := range inputs {
		visited[k] = false
		trie.Insert(v, []byte(k))
		// Calculate the longest key length
		if longest < len(k) {
			longest = len(k)
		}
	}

	if longest != trie.MaxDepth() {
		t.Errorf("MaxDepth(): expected %d, got %d", longest,
			trie.MaxDepth())
	}
	// Walk the trie manually and make sure that we can get every value
	for k, v := range inputs {
		v1, ok := trie.Get([]byte(k))
		if !ok {
			t.Errorf("Key %s missing from trie", k)
		}
		if v1 != v {
			t.Errorf("Key %s: expected value %d, got %d", k, v, v1)
		}
	}

	// Verify Match()
	for k, v := range inputs {
		longerKey := k + "foo"
		v1, ok := trie.Match([]byte(longerKey))
		if !ok {
			t.Errorf("Value %s doesn't match", longerKey)
		}
		if v1 != v {
			t.Errorf("Match Key %s: expected value %d, got %d",
				k, v, v1)
		}
	}

	// Walk the trie and make sure everything is there
	walker := func(key []byte, val interface{}) {
		skey := string(key)
		visited[skey] = true
		if val != inputs[skey] {
			t.Errorf("Key %s: expected value %d, got %d",
				skey, inputs[skey], val)
		}
	}

	trie.Do(walker)

	// Make sure we visited every node
	for k, v := range visited {
		if !v {
			t.Errorf("Skipped key %s", k)
		}
	}

	// Make sure that a word not in a trie doesn't match
	{
		v, ok := trie.Get([]byte(badword))
		if v != nil || ok {
			t.Error("Get() returns true for a missing word")
		}
		v, ok = trie.Match([]byte(badword))
		if v != nil || ok {
			t.Error("Match() returns true for a missing word")
		}
	}
}


func Example() {
	trie := bytetrie.New()
	trie.Insert(true, []byte("hello"))
	trie.Insert(true, []byte("world"))
	trie.Insert(true, []byte("help"))
	trie.Insert(true, []byte("work"))
	trie.PrintKeys()
}

func ExampleTrie_Get() {
	trie := bytetrie.New()
	trie.Insert(1, []byte("hello"))
	trie.Insert(2, []byte("world"))
	trie.Insert(3, []byte("help"))
	trie.Insert(4, []byte("work"))
	value, ok := trie.Get([]byte("help"))
	fmt.Println(value, ok)
	// Output: 3 true
}

func ExampleTrie_Match() {
	trie := bytetrie.New()
	trie.Insert(1, []byte("hello"))
	trie.Insert(2, []byte("world"))
	trie.Insert(3, []byte("help"))
	trie.Insert(4, []byte("work"))
	value, ok := trie.Match([]byte("hello world"))
	fmt.Println(value, ok)
	// Output: 1 true
}

func ExampleTrie_Do() {
	trie := bytetrie.New()
	trie.Insert(1, []byte("hello"))
	trie.Insert(2, []byte("world"))
	trie.Insert(3, []byte("help"))
	trie.Insert(4, []byte("work"))
	// Count number of entries in a tree and total of all values
	count := 0
	total := 0
	counter := func(key []byte, val interface{}) {
		intval := val.(int)
		total += intval
		count++

	}
	trie.Do(counter)
	fmt.Println(count, total)
	// Output: 4 10
}