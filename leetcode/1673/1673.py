class Solution:
    def mostCompetitive(self, nums: List[int], k: int) -> List[int]:
        unused = len(nums) - k
        top = 0
        stack = []

        for num in nums:
            while len(stack) > 0 and num < stack[-1] and unused > 0:
                unused -= 1
                stack.pop()
            if len(stack) < k:
                stack.append(num)
            else:
                unused -= 1

        return stack
