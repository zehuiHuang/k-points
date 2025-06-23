package leetcode

// 前缀树
type trieNode struct {
	nexts   [26]*trieNode
	passCnt int
	end     bool
}

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: &trieNode{},
	}
}

func (t *Trie) Search(word string) bool {
	// 查找目标节点，使得从根节点开始抵达目标节点沿路字符形成的字符串恰好等于 word
	node := t.search(word)
	return node != nil && node.end
}

func (t *Trie) search(target string) *trieNode {
	// 移动指针从根节点出发
	move := t.root
	// 依次遍历 target 中的每个字符
	for _, ch := range target {
		// 倘若 nexts 中不存在对应于这个字符的节点，说明该单词没插入过，返回 nil
		if move.nexts[ch-'a'] == nil { //ch-'a' 相对a的下标
			return nil
		}
		// 指针向着子节点移动
		move = move.nexts[ch-'a']
	}
	// 来到末尾，说明已经完全匹配好单词，直接返回这个节点
	// 需要注意，找到目标节点不一定代表单词存在，因为该节点的 end 标识未必为 true
	// 比如我们之前往 trie 中插入了 apple 这个单词，但是查找 app 这个单词时，预期的返回结果应该是不存在，此时就需要使用到 end 标识 进行区分
	return move
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}
func (t *Trie) PassCnt(prefix string) int {
	node := t.search(prefix)
	if node == nil {
		return 0
	}
	return node.passCnt
}

func (t *Trie) Insert(word string) {
	if t.Search(word) {
		return
	}
	move := t.root
	for _, ch := range word {
		if move.nexts[ch-'a'] == nil {
			move.nexts[ch-'a'] = &trieNode{}
		}
		move.nexts[ch-'a'].passCnt++
		move = move.nexts[ch-'a']
	}
	move.end = true
}

func (t *Trie) Erase(word string) bool {
	if !t.Search(word) {
		return false
	}
	move := t.root
	for _, ch := range word {
		move.nexts[ch-'a'].passCnt--
		if move.nexts[ch-'a'].passCnt == 0 {
			move.nexts[ch-'a'] = nil
			return true
		}
		move = move.nexts[ch-'a']
	}
	move.end = false
	return true
}
