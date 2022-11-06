package main

import "fmt"

//Given a string s, return all the palindromic permutations (without duplicates) of it.
//
//You may return the answer in any order. If s has no palindromic permutation, return an empty list.
//
//
//
//Example 1:
//
//Input: s = "aabb"
//Output: ["abba","baab"]
//Example 2:
//
//Input: s = "abc"
//Output: []

func generatePalindromes(s string) []string {

	var result []string
	var tempArr []rune
	var tempStr string
	var tempMap map[rune]int
	var tempMap2 map[rune]int

	tempMap = make(map[rune]int)
	tempMap2 = make(map[rune]int)

	for _, v := range s {
		tempMap[v]++
	}

	for _, v := range s {
		if tempMap[v]%2 == 1 {
			tempArr = append(tempArr, v)
		}
	}

	if len(tempArr) == 0 {
		for _, v := range s {
			tempMap2[v]++
		}
		for _, v := range s {
			if tempMap2[v]%2 == 1 {
				tempStr += string(v)
			}
		}
		result = append(result, tempStr)
		return result
	}

	for i := 0; i < len(tempArr); i++ {
		for j := i + 1; j < len(tempArr); j++ {
			tempStr = string(tempArr[i]) + string(tempArr[j])
			tempMap2[tempArr[i]]++
			tempMap2[tempArr[j]]++
			for _, v := range s {
				if tempMap2[v]%2 == 1 {
					tempStr += string(v)
				}
			}
			result = append(result, tempStr)
			tempMap2[tempArr[i]]--
			tempMap2[tempArr[j]]--
		}
	}

	return result
}

func main() {
	s := "aabb"
	fmt.Println(generatePalindromes(s))
}
