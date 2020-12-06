class Solution:
    def maxOperations(self, nums: List[int], k: int) -> int:
        count = {}
        lst = []
        for n in nums:
            if n < k:
                count[n] = count.get(n, 0) + 1

        res = 0
        for n in list(filter(lambda x: x < k/2, nums)):
            target = k-n
            cnt = count.get(target, 0)
            if cnt > 0:
                count[target] -= 1
                res += 1

        res += int(count.get(k/2, 0) / 2)

        return res
