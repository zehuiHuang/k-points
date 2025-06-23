package self

import "math/rand"

type skipTable struct {
	head *node2
}
type node2 struct {
	next     []*node2
	key, val int
}

func (s *skipTable) roll() int {
	var level int
	for rand.Int()%2 > 0 {
		level++
	}
	return level
}

func (s *skipTable) search(key int) *node2 {
	move := s.head
	if move == nil {
		return nil
	}
	for level := len(s.head.next) - 1; level >= 0; level-- {
		for move.next[level] != nil && key > move.next[level].key {
			move = move.next[level]
		}
		if move.next[level] != nil && move.next[level].key == key {
			return move.next[level]
		}
	}
	return nil
}

func (s *skipTable) Put(key, val int) {
	n := s.search(key)
	if n != nil {
		n.val = val
		return
	}
	level := s.roll()
	newNode := node2{
		key:  key,
		val:  val,
		next: make([]*node2, level+1),
	}
	move := s.head
	for len(move.next)-1 < level {
		move.next = append(move.next, nil)
	}
	for level := level; level >= 0; level-- {
		if move.next[level] != nil && key > move.next[level].key {
			move = move.next[level]
		}
		newNode.next[level] = move.next[level]
		move.next[level] = &newNode
	}
}
