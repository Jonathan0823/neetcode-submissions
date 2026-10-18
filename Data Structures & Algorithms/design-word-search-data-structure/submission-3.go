type TrieNode struct { 
	children [26]*TrieNode
	endOfWord bool
}
type WordDictionary struct {
    root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{ 
		root: &TrieNode{},
	}
}

func (this *WordDictionary) AddWord(word string)  {
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

func (this *WordDictionary) Search(word string) bool {
	return this.dfs(word, 0, this.root)
}

func (this *WordDictionary) dfs(word string, j int, root *TrieNode) bool { 
	cur := root
	for i := j; i < len(word); i++ { 
		c := word[i]
		if c == '.' { 
			for _, child := range cur.children { 
				if child != nil && this.dfs(word, i+1, child) { 
					return true
				}
			}
			return false
		} else { 
			 index := c - 'a'
			 if cur.children[index] == nil { 
				return false
			 }
			 cur = cur.children[index]
		}
	}

	return cur.endOfWord
}
