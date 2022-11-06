package AddAndSearchWord_Tries

type WordDictionary struct {
	isWord bool
	child  map[byte]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{
		child: map[byte]*WordDictionary{},
	}
}

func (this *WordDictionary) AddWord(word string) {

	cur := this
	for i := 0; i < len(word); i++ {
		curChar := word[i]
		if _, ok := cur.child[curChar]; !ok {
			cur.child[curChar] = &WordDictionary{
				child: map[byte]*WordDictionary{},
			}
		}
		cur = cur.child[curChar]
	}
	cur.isWord = true
}

func (this *WordDictionary) Search(word string) bool {
	return match(this, word, 0)
}
func match(this *WordDictionary, word string, index int) bool {
	if index == len(word) {
		return true
	}
	if index == len(word)-1 && this.isWord {
		return this.isWord
	}
	if word[index] != '.' {
		chByte := word[index]
		if v, ok := this.child[chByte]; ok {
			return match(v, word, index+1)
		} else {
			return false
		}
	} else {
		for _, value := range this.child {
			if match(value, word, index+1) {
				return true
			}
		}
	}
	return false
}
