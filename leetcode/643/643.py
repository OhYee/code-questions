class Solution:
    def findMaxAverage(self, nums: List[int], k: int) -> float:
        n = len(nums)

        sums = sum(nums[:k])
        avg = -10000
        for l in range(n-k+1):
            avg = max(avg, sums / k)

            r = l + k
            if r < n:
                sums = sums - nums[l] + nums[r]
        return avg
