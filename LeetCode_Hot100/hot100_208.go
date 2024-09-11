package main

type Trie struct {
	charSet [26]*Trie
	isLast  bool
}

func Constructor() Trie {
	return Trie{isLast: false}
}

func (this *Trie) Insert(word string) {
	node := this
	for i := 0; i < len(word); i++ {
		if node.charSet[int(word[i]-'a')] == nil {
			node.charSet[int(word[i]-'a')] = &Trie{}
			node = node.charSet[int(word[i]-'a')]
		} else {
			node = node.charSet[int(word[i]-'a')]
		}
	}
	node.isLast = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for i := 0; i < len(word); i++ {
		if node.charSet[int(word[i]-'a')] != nil {
			node = node.charSet[int(word[i]-'a')]
		} else {
			return false
		}
	}
	return node.isLast
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for i := 0; i < len(prefix); i++ {
		if node.charSet[int(prefix[i]-'a')] != nil {
			node = node.charSet[int(prefix[i]-'a')]
		} else {
			return false
		}
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
