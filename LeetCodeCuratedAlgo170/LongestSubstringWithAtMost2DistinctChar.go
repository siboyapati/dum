package main

import "fmt"

func lengthOfLongestSubstringTwoDistinct(s string) int {
	var i int
	var j int
	m := make(map[byte][]int)
	var result int
	for j < len(s) {
		curChar := s[j]
		m[curChar] = append(m[curChar], j)

		if len(m) <= 2 {
			result = Max1(result, j-i+1)
		}

		for len(m) > 2 {
			charAtFirst := s[i]
			arr := m[charAtFirst]
			if len(m[charAtFirst]) == 1 {
				delete(m, charAtFirst)
			} else {
				m[charAtFirst] = arr[1:]
			}
			i++
		}
		j++
	}
	return result
}

func Max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	s := "eceba"
	//s := "ccaabbb"
	//s := "abaccc"
	fmt.Println(lengthOfLongestSubstringTwoDistinct(s))
}
