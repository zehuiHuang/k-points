package practice

//todo

type Skiplist struct {
	head *node
}

type node struct {
	list     []*node //节点的个数就是跳表最大的高度,list[level]即为当前节点的下一个节点，level为当前层数（注意：从0层开始）
	key, val int
}

//func Constructor() Skiplist {
//
//}

func (this *Skiplist) Search(target int) bool {
	return false
}

func (this *Skiplist) Add(num int) {

}

func (this *Skiplist) Erase(num int) bool {
	return false
}
