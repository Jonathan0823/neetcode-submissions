type TrieNode struct { 
	children [26]*TrieNode
	isEnd bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
    return WordDictionary{ root: &TrieNode{} }
}

func (this *WordDictionary) AddWord(word string)  {
	cur := this.root
	for _, b := range word { 
		idx := b - 'a'
		if cur.children[idx] == nil { 
			cur.children[idx] = &TrieNode{}
		}
		cur = cur.children[idx]
	}
	cur.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	var dfs func(root *TrieNode, word string, index int) bool 
	dfs = func(root *TrieNode, word string, index int) bool { 
		if root == nil { 
			return false
		}

		if index == len(word) { 
			return root.isEnd
		}

		char := word[index]

		if char == '.' { 
			for _, child := range root.children { 
				if child != nil && dfs(child, word, index + 1) { 
					return true
				}
			}
			return false
		}
		
		return dfs(root.children[char - 'a'], word, index + 1)
	}

	return dfs(this.root, word, 0)
}
