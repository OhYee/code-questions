class Solution:
    def maximumProduct(self, nums: List[int]) -> int:
        n = len(nums)
        nums.sort()
        s = []
        if n <= 6:
            s = nums[:]
        else:
            s = nums[:3] + nums[-3:]
        
        # print(s)
        m = len(s)
        
        res = -1000000000
        def dfs(t, deep, multi):
            # print(t, deep, multi)
            if deep == 1:
                nonlocal res
                res = max(res, multi)
                return
            for i in range(t+1, m):
                dfs(i, deep-1, multi*s[i])

        for i in range(m):
            dfs(i, 3, s[i])
        return res
