from typing import List

class Solution:
    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        result = 0
        dict = {}
        for n1 in nums1:
            for n2 in nums2:
                if n1+n2 in dict:
                    dict[n1+n2]= dict[n1+n2]+1
                else:
                    dict[n1+n2] = 1

        for n3 in nums3:
            for n4 in nums4:
                target = -(n3 + n4)
                if target in dict:
                    result += dict[target]
        return result
