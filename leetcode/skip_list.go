package leetcode

import (
	"math/rand"
)

//跳表

type Skiplist struct {
	head *node
}
type node struct {
	list     []*node //节点的个数就是跳表最大的高度,list[level]即为当前节点的下一个节点，level为当前层数（注意：从0层开始）
	key, val int
}

// Get 根据 key 读取 val，第二个 bool flag 反映 key 在 skiplist 中是否存在
func (s *Skiplist) Get(key int) (int, bool) {
	// 根据 key 尝试检索对应的 node，如果 node 存在，则返回对应的 val
	if _node := s.search(key); _node != nil {
		return _node.val, true
	}

	return -1, false
}

// 从跳表中检索 key 对应的 node
func (s *Skiplist) search(key int) *node {
	// 每次检索从头部出发
	move := s.head
	// 每次检索从最大高度出发，直到来到首层
	for level := len(s.head.list) - 1; level >= 0; level-- {
		// 在每一层中持续向右遍历，直到下一个节点不存在或者 key 值大于等于 key
		for move.list[level] != nil && key > move.list[level].key {
			move = move.list[level]
		}
		// 如果 key 值相等，则找到了目标直接返回
		if move.list[level] != nil && move.list[level].key == key {
			return move.list[level]
		}
		// 当前层没找到目标，则层数减 1，继续向下
	}
	// 遍历完所有层数，都没有找到目标，返回 nil
	return nil
}

// roll 骰子，决定一个待插入的新节点在 skiplist 中最高层对应的 index
func (s *Skiplist) roll() int {
	var level int
	// 每次投出 1，则层数加 1
	for rand.Int()%2 > 0 {
		level++
	}
	return level
}

// Put 将 key-val 对加入 skiplist
func (s *Skiplist) Put(key, val int) {
	// 假如 kv对已存在，则直接对值进行更新并返回
	if _node := s.search(key); _node != nil {
		_node.val = val
		return
	}

	// roll 出新节点的高度
	level := s.roll()

	// 新节点高度超出跳表最大高度，则需要对高度进行补齐
	for len(s.head.list)-1 < level {
		s.head.list = append(s.head.list, nil)
	}

	// 创建出新的节点
	newNode := node{
		key:  key,
		val:  val,
		list: make([]*node, level+1), //level+1 是因为有空的尾节点
	}

	// 从头节点的最高层出发
	move := s.head
	for level := level; level >= 0; level-- {
		// 向右遍历，直到右侧节点不存在或者 key 值大于 key
		for move.list[level] != nil && key > move.list[level].key {
			move = move.list[level]
		}
		// 调整指针关系，完成新节点的插入（链表插入指定节点常规操作）  newNode.list[level]为插入节点的下一个节点， move.list[level] 为move节点的下一个节点，
		newNode.list[level] = move.list[level]
		move.list[level] = &newNode
	}
}
