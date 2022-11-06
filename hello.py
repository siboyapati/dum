from typing import List
class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        result = [None] * len(nums)
        low = 0
        high = len(nums)-1
        arrLen = len(nums)-1
        while low <= high :
            lowSquare = nums[low]*nums[low]
            highSquare = nums[high]*nums[high]

            if highSquare > lowSquare :
                result[arrLen] = highSquare
                high-=1
            else:
                result[arrLen] = lowSquare
                low+=1

            arrLen-=1
        return result
    # Two pointer approach
    def sortedSquares(self, nums: List[int]) -> List[int]:
        result = [None]*len(nums)
        res_index = len(nums) - 1
        first = 0
        last = len(nums)-1

        while (first <= last):
            if((nums[first]*nums[first]) > (nums[last]*nums[last])):
                result[res_index] = (nums[first]*nums[first])
                first += 1
            else:
                result[res_index] = (nums[last]*nums[last])
                last -= 1
            res_index -= 1
        return (result)

    def hello(self, name: string):