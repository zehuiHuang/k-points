package self

//前缀树练习

type Trie struct {
	children [26]*Trie
	end      bool
	count    int
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	node := t
	for _, c := range word {
		if node.children[c-'a'] == nil {
			node.children[c-'a'] = &Trie{}
			node.count++
		}
		node = node.children[c-'a']
	}
	node.end = true
}

func (t *Trie) Search(word string) bool {
	node := t
	for _, c := range word {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return node != nil && node.end
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t
	for _, c := range prefix {
		if node.children[c-'a'] == nil {
			return false
		}
		node = node.children[c-'a']
	}
	return node != nil
}
