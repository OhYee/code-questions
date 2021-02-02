class Solution:
    def characterReplacement(self, s: str, k: int) -> int:
        n = len(s)
        l = 0
        r = 0
        count = [0 for i in range(27)]
        res = 0
        for r in range(n):
            rc = ord(s[r]) - ord('A')
            count[rc] += 1
            count[26] += 1
            while count[26] - max(count[:26]) > k:
                lc = ord(s[l]) - ord('A')
                count[lc] -= 1
                count[26] -= 1
                l += 1
            res = max(res, r-l+1)
        return res
