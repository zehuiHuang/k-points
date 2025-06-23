package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestPut(t *testing.T) {
	var skiplist = &Skiplist{
		head: &node{},
	}
	skiplist.Put(1, 100)
	skiplist.Put(2, 200)
	skiplist.Put(3, 300)
	fmt.Println("-----")
}

func a() int {
	var level int
	// 每次投出 1，则层数加 1
	r := rand.Int() % 2
	for r > 0 {
		level++
	}
	return level
}
func TestName22(t *testing.T) {
	var m = make(map[int]int)
	for i := 0; i < 100; i++ {
		v := a()
		m[v] = m[v] + 1
	}
	fmt.Println("dddddddddddddddddd")
}
