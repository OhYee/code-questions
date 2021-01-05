class Solution:
    def largeGroupPositions(self, s: str) -> List[List[int]]:
        count = 0
        start = 0
        pre = ""
        res = []
        for i, c in enumerate(s):
            if c == pre:
                count += 1
            else:
                if count >= 3:
                    res.append([start, i-1])
                pre = c
                count = 1
                start = i
        if count >= 3:
            res.append([start, i])
        return res
