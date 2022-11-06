package ReplaceWords

import "strings"

//https://leetcode.com/explore/learn/card/trie/148/practical-application-i/1053/

type Tries struct {
	IsWord bool
	child  map[byte]*Tries
}

func Constructor() Tries {
	return Tries{
		child: map[byte]*Tries{},
	}
}

func replaceWords(dictionary []string, sentence string) string {

	node := Constructor()
	for i := 0; i < len(dictionary); i++ {
		cur := &node
		curWord := dictionary[i]
		for j := 0; j < len(curWord); j++ {
			if cur.child[curWord[j]] == nil {
				cur.child[curWord[j]] = &Tries{child: map[byte]*Tries{}}
			}
			cur = cur.child[curWord[j]]
		}
		cur.IsWord = true
	}
	sentArray := strings.Split(sentence, " ")
	for i := 0; i < len(sentArray); i++ {
		cur := &node
		curWord := sentArray[i]
		var tempByteArray []byte
		for j := 0; j < len(curWord); j++ {
			if cur.child[curWord[j]] == nil {
				break
			}
			tempByteArray = append(tempByteArray, curWord[j])
			cur = cur.child[curWord[j]]
			if cur.IsWord {
				sentArray[i] = string(tempByteArray)
				break
			}
		}
	}
	return strings.Join(sentArray, " ")
}
