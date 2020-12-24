class Solution:
    def candy(self, ratings: List[int]) -> int:
        n = len(ratings)
        left = [0] * n
        right = [0] * n
        res = 0
        for i in range(n):
            j = n-i-1
            left[i]  = left[i-1]  + 1 if i != 0   and ratings[i] > ratings[i-1] else 1
            right[j] = right[j+1] + 1 if j != n-1 and ratings[j] > ratings[j+1] else 1    
        for i in range(n):
            res += max(left[i], right[i])
        return res
