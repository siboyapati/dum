package PermutationInString_567

//Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
//
//In other words, return true if one of s1's permutations is the substring of s2.

//Input: s1 = "ab", s2 = "eidbaooo"
//Output: true
//Explanation: s2 contains one permutation of s1 ("ba").
//
//Input: s1 = "ab", s2 = "eidboaoo"
//Output: false

func checkInclusion(s1 string, s2 string) bool {

	var m map[byte]int
	for i := 0; i < len(s2); i++ {
		m[s2[i]]++
	}

	i, j := 0, 1024

	for i < len(s1){

	}

}
