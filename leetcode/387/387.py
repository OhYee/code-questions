class Solution:
    def firstUniqChar(self, s: str) -> int:
        for idx,c in enumerate(s):
            if s.count(c) == 1:
                return idx
        return -1
