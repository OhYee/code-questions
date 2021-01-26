class Solution:
    def numEquivDominoPairs(self, dominoes: List[List[int]]) -> int:
        dominoes = [tuple(sorted(d)) for d in dominoes]
        # print(dominoes)

        count = {}
        for d in dominoes:
            count[d] = count.get(d, 0) + 1

        res = 0
        for c in count.values():
            if c >= 2:
                res += self.C(c, 2)

        return int(res)

    def C(self, n, m):
        a = b = res = 1
        m = min(n, n-m)
        for j in range(0, m):
            a = a * (n-j)
            b = b * (m-j)
            res = a // b 
        return res


