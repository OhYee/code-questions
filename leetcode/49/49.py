class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        m = {}
        for s in strs:
            key = str(sorted(s))
            if key not in m:
                m[key] = []
            m[key].append(s)
        return list(m.values())
