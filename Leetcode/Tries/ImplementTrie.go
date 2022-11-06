package Tries

type Trie struct {
	isWord      bool
	childrenMap map[byte]*Trie
}

func Constructor() Trie {
	return Trie{
		isWord:      false,
		childrenMap: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	cur := this
	for i := 0; i < len(word); i++ {
		curChar := word[i]
		if _, ok := cur.childrenMap[curChar]; !ok {
			cur.childrenMap[curChar] = &Trie{
				isWord:      false,
				childrenMap: make(map[byte]*Trie),
			}
		}
		cur = cur.childrenMap[curChar]
	}
	cur.isWord = true
}

func (this *Trie) Search(word string) bool {

	cur := this
	for i := 0; i < len(word); i++ {
		curChar := word[i]
		if v, ok := cur.childrenMap[curChar]; ok {
			cur = v
		} else {
			return false
		}
	}
	if cur.isWord == true {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for i := 0; i < len(prefix); i++ {
		curChar := prefix[i]
		if v, ok := cur.childrenMap[curChar]; ok {
			cur = v
		} else {
			return false
		}
	}
	return true
}
