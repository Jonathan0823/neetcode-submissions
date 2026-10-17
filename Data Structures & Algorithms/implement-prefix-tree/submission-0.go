type TrieNode struct {
	children [26]*TrieNode
	endOfWord bool
}

type PrefixTree struct {
	root *TrieNode
}

func Constructor() PrefixTree {
	return PrefixTree{root: &TrieNode{}}
    
}

func (this *PrefixTree) Insert(word string) {
	cur := this.root
	for _, c := range word { 
		i := c - 'a'
		if cur.children[i] == nil { 
			cur.children[i] = &TrieNode{}
		}
		cur = cur.children[i]
	}

	cur.endOfWord = true
}

func (this *PrefixTree) Search(word string) bool {
	cur := this.root
	for _, c := range word { 
		i := c - 'a'
		if cur.children[i] == nil { 
			return false
		}
		cur = cur.children[i]
	}

	return cur.endOfWord
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	cur := this.root
	for _, c := range prefix { 
		i := c -'a'
		if cur.children[i] == nil { 
			return false
		}
		cur = cur.children[i]
	}

	return true
}
