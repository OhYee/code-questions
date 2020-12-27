class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        m = {}
        n = len(s)
        for i in range(n):
            sc = s[i]
            tc = t[i]
            if sc not in m:
                m[sc] = tc
            else:
                if m[sc] != tc:
                    return False
        return len(m.values()) == len(set(m.values()))
