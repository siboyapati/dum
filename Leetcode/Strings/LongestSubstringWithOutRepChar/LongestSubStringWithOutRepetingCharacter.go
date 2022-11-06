package LongestSubstringWithOutRepChar



/*Given a string s, find the length of the longest substring without repeating characters.
Example 1:

Input: s = "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: s = "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: s = "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.

*/


func LengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	var result int
	var prevIndex int
	var curIndex int
	for curIndex < len(s) {
		if _, ok := m[s[curIndex]]; ok {
			delete(m, s[prevIndex])
			prevIndex++
			continue
		}
		m[s[curIndex]] = curIndex
		curIndex++
		result = Max(result, len(m))
	}
	return result

}

func LengthOfLongestSubstring1(s string) int {
	m := make(map[byte]int)
	var result int
	var prevIndex int
	for i := 0; i < len(s); i++ {
		if v, ok := m[s[i]]; ok {
			if v > prevIndex {
				prevIndex = v + 1
			}
		}
		result = Max(result, i-prevIndex+1)
		m[s[i]] = i
	}
	return result

}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
