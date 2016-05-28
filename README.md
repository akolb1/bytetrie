# bytetrie
Yet another GO implementation of trie with byte keys and Match() function

There are several trie implementations around. This one is very simple and only
supports trie where keys are byte sequences. Also, it has a Match() function
which is missing in other trie implementation. The intention is to use it to
match file type by initial bytes in the file.

The library is inspired by Trie code developed by Drew Noakes for the
https://github.com/drewnoakes/metadata-extractor library.

## Installation

Standard `go get`:

```
$ go get github.com/akolb1/bytetrie
```

## Usage & Example

For usage and examples see the
[![GoDoc](https://godoc.org/github.com/akolb1/bytetrie?status.svg)](https://godoc.org/github.com/akolb1/bytetrie)

