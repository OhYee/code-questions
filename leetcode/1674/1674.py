class Solution:
    def minMoves(self, nums: List[int], limit: int) -> int:
        n = len(nums)
        l = int(n / 2)

        sumab = {}  # a + b 的个数
        minab = {}  # min(a, b) 的个数
        maxab = {}  # max(a, b) 的个数
        for i in range(l):
            a = nums[i]
            b = nums[-i-1]
            a, b = min(a, b), max(a, b)

            sumab[a+b] = sumab.get(a+b, 0) + 1
            minab[a] = minab.get(a, 0) + 1
            maxab[b] = maxab.get(b, 0) + 1

        maxpossible = 2*limit+1

        leminab = [0] * maxpossible  # <= min(a, b) 的个数
        lemaxab = [0] * maxpossible  # <= max(a, b) 的个数
        for i in range(1, maxpossible):
            leminab[i] = leminab[i-1] + minab.get(i, 0)
            lemaxab[i] = lemaxab[i-1] + maxab.get(i, 0)

        # print(sumab)
        # print(minab)
        # print(maxab)
        # print(leminab)
        # print(lemaxab)

        res = n
        for target in range(2, maxpossible):
            op0 = sumab.get(target, 0)
            op1 = leminab[target - 1] - (lemaxab[target - limit - 1]
                                         if target - limit - 1 > 0 else 0) - op0
            op2 = l - op0 - op1
            # print(target, op0, op1, op2)
            res = min(res, op1+2*op2)

        return res
