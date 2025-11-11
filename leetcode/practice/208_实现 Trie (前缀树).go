package practice

//定义这种结构的目的是为了在每一层存储a~z的字符

type Trie struct {
	children [26]*Trie
	end      bool
}

//func Constructor() Trie {
//	return Trie{
//	}
//}

func (this *Trie) Insert(word string) {
	node := this
	for _, v := range word {
		//表示以word的某一个字符 在当前节点的子节点下不存在,那么就创建一个节点,并放到子集中
		if node.children[v-'a'] == nil {
			node.children[v-'a'] = &Trie{}
		}
		//指针下移,准备针对下一个字符串存储
		node = node.children[v-'a']
	}
	node.end = true
}

// 如果某字符存在与当前节点下的children,那么指针下移,继续对下一个字符进行判断,且遍历到最后node也存在且end标识为true

func (this *Trie) Search(word string) bool {
	node := this
	for _, v := range word {
		if node.children[v-'a'] == nil {
			return false
		}
		node = node.children[v-'a']
	}
	return node != nil && node.end
}

// 如果某字符存在与当前节点下的children,那么指针下移,继续对下一个字符进行判断,且遍历到最后只要node不为空即可

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, v := range prefix {
		if node.children[v-'a'] == nil {
			return false
		}
		node = node.children[v-'a']
	}
	return node != nil
}
