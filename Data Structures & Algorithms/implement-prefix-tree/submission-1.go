type TrieNode struct { 
	child [26]*TrieNode
	isEnd bool
}
type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{
		root: &TrieNode{},
	}
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root 
	for _, c := range word { 
		i := c - 'a'
		if cur.child[i] == nil {
			cur.child[i] = &TrieNode{}
		}
		cur = cur.child[i]
	}
	cur.isEnd = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.searchPrefix(word)
	return cur != nil && cur.isEnd == true
}


func (this *PrefixTree) StartsWith(prefix string) bool {
	return this.searchPrefix(prefix) != nil
}

func (this *PrefixTree) searchPrefix(prefix string) *TrieNode { 
	cur := this.root
	for _, c := range prefix { 
		i := c - 'a'
		if cur.child[i] == nil { 
			return nil
		}
		cur = cur.child[i]
	}
	return cur
}
