class Solution:
    def wiggleMaxLength(self, nums: List[int]) -> int:
        if len(nums) < 2:
            return len(nums)

        pre = nums[0]
        flag = 0
        res = 1
        for n in nums:
            if n < pre and flag >= 0:
                res += 1
                flag = -1
            elif n > pre and flag <= 0:
                res += 1
                flag = 1
            pre = n
            
        return res
