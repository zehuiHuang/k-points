package leetcode

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	var a = "abdf"
	for _, ch := range a {
		fmt.Println(ch - 'a')
	}
}

func TestInsert(t *testing.T) {
	//情况1
	var rie = &Trie{}
	rie.root = &trieNode{}
	rie.Insert("word")
	rie.Search("word")
	rie.Erase("word")
}

func TestName(t *testing.T) {
	a := "abc"
	fmt.Println(a[1:])
}
