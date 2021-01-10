class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        if len(nums) == 0:
            return []
        nums.append(99999999999)
        res = []
        pre = 0
        start = 0
        for i, n in enumerate(nums):
            if i == 0:
                start = n
            else:
                if pre + 1 != n:
                    if start == pre:
                        res.append("%d"%pre)
                    else:
                        res.append("%d->%d"%(start, pre))
                    start = n
            pre = n
        
        return res
