class Solution:
    def findLengthOfLCIS(self, nums: List[int]) -> int:
        n = len(nums)
        res = 0
        count = 0
        for i in range(n):
            if i == 0:
                count = 1
            else:
                if nums[i] > nums[i-1]:
                    count += 1
                else:
                    res = max(res, count)
                    count = 1
        res = max(res, count)
        return res
        
