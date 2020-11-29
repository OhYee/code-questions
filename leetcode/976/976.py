class Solution:
    def largestPerimeter(self, A: List[int]) -> int:
        A.sort(reverse=True)
        l = len(A)
        for i in range(2, l):
            if A[i-2] < A[i-1] + A[i]:
                return A[i-2] + A[i-1] + A[i]
        return 0