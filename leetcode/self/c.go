package self

import (
	"math/rand"
)

//跳表练习

type SkipList2 struct {
	head *node
}

type node struct {
	key, val int
	list     []*node
}

func (sl SkipList2) get(key int) (int, bool) {
	nd := sl.search(key)
	if nd == nil {
		return 0, false
	}
	return nd.val, true
}
func (sl SkipList2) search(key int) *node {
	move := sl.head
	level := len(move.list) - 1
	for l := level; l >= 0; l-- {
		for move.list[level] != nil && key > move.list[level].key {
			move = move.list[level]
		}
		if move.list[level].key == key {
			return move.list[level]
		}
	}
	return nil
}

func roll() int {
	var level int
	if rand.Int()%2 > 0 {
		level++
	}
	return level
}
func (sl SkipList2) put(key, val int) {
	if ret := sl.search(key); ret != nil {
		ret.val = val
		return
	}
	//获取随机层数
	level := roll()
	move := sl.head
	//超过最高层，则head也需要跟着增加
	if len(move.list)-1 < level {
		move.list = append(move.list, nil)
	}
	//创建新节点
	newNodeData := &node{
		key:  key,
		val:  val,
		list: make([]*node, level+1),
	}
	for l := level; l >= 0; l-- {
		for move.list[level] != nil && key > move.list[level].key {
			move = move.list[level]
		}
		//指针变更
		newNodeData.list[level] = move.list[level]
		move.list[level] = newNodeData
	}
}

func (sl SkipList2) delete(key int) {
	if _node := sl.search(key); _node == nil {
		return
	}
	move := sl.head
	level := len(move.list) - 1
	for l := level; l >= 0; l-- {
		for move.list[level] != nil && key > move.list[level].key {
			move = move.list[level]
		}
		//if move.list[level] == nil || key > move.list[level].key {
		//	continue
		//}
		if move.list[level] != nil && move.list[level].key == key {
			move.list[level] = move.list[level].list[level]
		}
	}
	//更新调表最大高度
	var dif int
	for l := level; l > 0 && sl.head.list[l] == nil; l-- {
		dif++
	}
	sl.head.list = sl.head.list[:len(sl.head.list)-dif]
}
