class Solution:
    def findShortestSubArray(self, nums: List[int]) -> int:
        n = len(nums)
        numObj = {}
        for i in range(n):
            num = nums[i]
            count, first, last = numObj.get(num, [0, i, i])
            numObj[num] = [count + 1, first, i]
        
        maxValue = 0
        res = 0
        for (count, first, last) in numObj.values():
            if count > maxValue:
                maxValue = count
                res = last - first
            elif count == maxValue:
                res = min(res, last - first)
        return res + 1

