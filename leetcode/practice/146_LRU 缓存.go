package practice

// 思路:哈希表+双向链表,size、capacity、cache、head、tail

/**
1、设置两个虚拟头节点和未节点,后面所有的节点都放到虚拟节点中间
2、放入时,若不存在,则直接map复制,并且将节点放入到头部,若存在,则修改值,并将节点移动到头部,然后容量对比,超的进行尾部截取
3、查询时,若没有则直接返回-1,若有值,则从map中返回,并移动节点到头部

注意节点加入头部的指针移动:一共四步
先将node的next和pre 设置好
再将头部的下一个值得pre指向node
再将头部指向node
*/

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkNode
	head, tail *DLinkNode //头尾的节点对应的值为空(节点不为空),仅作为标识位
}

//双向链表

type DLinkNode struct {
	key, value int
	pre, next  *DLinkNode
}

//func Constructor(capacity int) LRUCache {
//	l := LRUCache{
//		size:     0,
//		capacity: capacity,
//		head:     &DLinkNode{0, 0, nil, nil},
//		tail:     &DLinkNode{0, 0, nil, nil},
//		cache:    make(map[int]*DLinkNode),
//	}
//	l.head.next = l.tail
//	l.tail.pre = l.head
//	return l
//}

func (this *LRUCache) Get(key int) int {
	//获取值,返回并将该节点移动到首位
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	//移动到首位
	this.moveHead(node)
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	//若缓存中不存在,则进行初始化node,放入map和双向链表
	//在判定链表长度是否大于容量,大于则进行尾部切除
	if _, ok := this.cache[key]; !ok {
		node := initNode(key, value)
		this.cache[key] = node
		//节点放到前面
		this.addHead(node)
		//长度+1
		this.size++
		//容量过载则进行切割,从尾部开始扔掉
		if this.size > this.capacity {
			deleted := this.deleteTail()
			delete(this.cache, deleted.key)
			this.size--
		}
	} else {
		//缓存存在则进行值变更,并进行位置移动
		node := this.cache[key]
		node.value = value
		this.moveHead(node)
	}
}

func (this *LRUCache) addHead(node *DLinkNode) {
	//head-n1-n2...-tail
	//将node加入到head和n1之间
	//则先将node的前后进行设置,然后在设置n1的pre指向node, head的next在也指向node
	node.pre = this.head
	node.next = this.head.next
	this.head.next.pre = node
	this.head.next = node
}

// 获取链尾的节点并删除它
func (this *LRUCache) deleteTail() *DLinkNode {
	node := this.tail.pre
	//b-node->tail
	b := node.pre
	b.next = this.tail
	this.tail.pre = b
	return node
}

// 将节点移动到头节点
func (this *LRUCache) moveHead(node *DLinkNode) {
	//head-1-node-3-tail
	//断开
	node.next.pre = node.pre
	node.pre.next = node.next

	//head-1-2
	//node
	//再连接
	node.next = this.head.next
	node.pre = this.head
	this.head.next.pre = node
	this.head.next = node
}

func initNode(key, value int) *DLinkNode {
	return &DLinkNode{
		key:   key,
		value: value,
	}
}
