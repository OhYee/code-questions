class Solution:
    def findDisappearedNumbers(self, nums: List[int]) -> List[int]:
        n = len(nums)

        for i in range(n):
            t = nums[i] - 1
            while t >= 0:
                temp = nums[t] - 1
                nums[t] = -1
                t = temp
                
        return [i+1 for i in range(n) if nums[i] >= 0]
