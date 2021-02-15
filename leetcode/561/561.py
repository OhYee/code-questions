class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        return sum(map(lambda x: x[1], filter(lambda x: x[0] % 2 == 0, enumerate(sorted(nums)))))
