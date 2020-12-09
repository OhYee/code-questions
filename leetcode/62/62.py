class Solution:
    def uniquePaths(self, m: int, n: int) -> int:
        def f(a):
            if (a==1 or a==0):
                return 1
            return f(a-1)*a
        def C(a,b):
            return int(f(a)/f(b)/f(a-b))
        return C(m+n-2, m-1)
