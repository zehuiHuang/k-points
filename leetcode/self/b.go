package self

import (
	"fmt"
	"math/rand"
)

const (
	MaxLevel = 16
)

type Node struct {
	key     int
	forward []*Node
}

type SkipList struct {
	header *Node
	level  int
}

func newNode(key int, level int) *Node {
	return &Node{
		key:     key,
		forward: make([]*Node, level+1),
	}
}

func newSkipList() *SkipList {
	header := newNode(0, MaxLevel)
	return &SkipList{
		header: header,
		level:  0,
	}
}

func (sl *SkipList) randomLevel() int {
	level := 0
	for rand.Float32() < 0.5 && level < MaxLevel {
		level++
	}
	return level
}

func (sl *SkipList) Insert(key int) {
	update := make([]*Node, MaxLevel+1)
	current := sl.header

	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
		update[i] = current
	}

	current = current.forward[0]

	if current == nil || current.key != key {
		level := sl.randomLevel()

		if level > sl.level {
			for i := sl.level + 1; i <= level; i++ {
				update[i] = sl.header
			}
			sl.level = level
		}

		newNode := newNode(key, level)

		for i := 0; i <= level; i++ {
			newNode.forward[i] = update[i].forward[i]
			update[i].forward[i] = newNode
		}

		fmt.Printf("Key %d inserted\n", key)
	} else {
		fmt.Printf("Key %d already exists\n", key)
	}
}

func (sl *SkipList) Search(key int) bool {
	current := sl.header

	for i := sl.level; i >= 0; i-- {
		for current.forward[i] != nil && current.forward[i].key < key {
			current = current.forward[i]
		}
	}

	current = current.forward[0]

	if current != nil && current.key == key {
		fmt.Printf("Key %d found\n", key)
		return true
	}

	fmt.Printf("Key %d not found\n", key)
	return false
}

func main() {
	skipList := newSkipList()

	keys := []int{3, 6, 7, 9, 12, 19, 17, 26, 21, 25}

	for _, key := range keys {
		skipList.Insert(key)
	}

	searchKeys := []int{7, 19, 26, 50, 3}

	for _, key := range searchKeys {
		skipList.Search(key)
	}
}
