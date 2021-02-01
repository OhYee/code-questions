class Solution:
    def fairCandySwap(self, A: List[int], B: List[int]) -> List[int]:
        sa = sum(A)
        sb = sum(B)
        total = sa + sb
        half = total // 2
        m = {}
        for a in A:
            m[half - (sa - a)] = a
        for b in B:
            if b in m:
                return [m[b], b]
        return [-1, -1]
