class Solution:
    def maxScore(self, cardPoints: List[int], k: int) -> int:
        n = len(cardPoints)
        left  = [0 for i in range(n+1)]
        right = [0 for i in range(n+1)]

        for i in range(1, n):
            left[i]  = left[i-1]  + cardPoints[i-1]
            right[i] = right[i-1] + cardPoints[n-i]

        res = 0
        for i in range(k+1):
            res = max(res, left[i] + right[k-i])
        return res
