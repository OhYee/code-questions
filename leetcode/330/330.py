class Solution:
    def minPatches(self, nums: List[int], n: int) -> int:
        l = len(nums)
        m = 1
        i = 0
        res = 0
        while m <= n:
            if i < l and nums[i] <= m:
                m += nums[i]
                i += 1
            else:
                m += m 
                res += 1
        return res
