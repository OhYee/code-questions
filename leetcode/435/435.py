class Solution:
    def eraseOverlapIntervals(self, intervals: List[List[int]]) -> int:
        intervals.sort(key=lambda x: x[1])
        pre_end = 0
        res = 0
        for idx, interval in enumerate(intervals):
            if idx == 0 or interval[0] >= pre_end:
                pre_end = interval[1]
            else:
                res += 1
        return res
