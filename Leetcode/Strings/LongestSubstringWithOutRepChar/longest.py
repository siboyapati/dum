class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        start = 0
        result = 0
        m = {}
        for i in range(len(s)):
            if s[i] in m:
                if m[s[i]] > start:
                    start = m[s[i]] + 1
            result = max(result, i - start + 1)
            m[s[i]] = 1
        return result
