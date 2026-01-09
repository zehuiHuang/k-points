package practice

import "math/rand"

type Skiplist struct {
	head *node
}

type node struct {
	list     []*node //节点的个数就是跳表最大的高度,list[level]即为当前节点的下一个节点，level为当前层数（注意：从0层开始）
	key, val int
}

//func Constructor() Skiplist {
//	return Skiplist{&node{list: make([]*node, 1)}}
//}

func (this *Skiplist) Search(target int) bool {
	nodePod := this.search(target)
	if nodePod != nil {
		return true
	}
	return false
}

func (s *Skiplist) Add(key int) {
	val := key
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

func (s *Skiplist) Erase(num int) bool {
	// 查找要删除的节点
	targetNode := s.search(num)
	if targetNode == nil {
		return false // 节点不存在，返回false
	}

	// 记录需要更新的前驱节点
	prevNodes := make([]*node, len(s.head.list))

	// 从最高层开始查找每一层中目标节点的前驱节点
	current := s.head
	for level := len(s.head.list) - 1; level >= 0; level-- {
		// 在当前层中向右移动，直到下一个节点是要删除的节点
		for current.list[level] != nil && current.list[level].key < num {
			current = current.list[level]
		}

		// 记录当前层中目标节点的前驱节点
		prevNodes[level] = current
	}

	// 更新每一层的指针，跳过目标节点
	for level := 0; level < len(targetNode.list); level++ {
		// 将前驱节点的指针指向目标节点的下一个节点
		prevNodes[level].list[level] = targetNode.list[level]
	}

	// 检查是否需要减少跳表的最大高度（如果最高层没有节点了）
	for i := len(s.head.list) - 1; i >= 0 && s.head.list[i] == nil; i-- {
		s.head.list = s.head.list[:i]
	}
	return true
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
