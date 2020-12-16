class Solution:
    def wordPattern(self, pattern: str, s: str) -> bool:
        m = {}
        m2 = {}
        words = s.split(" ")

        if len(words) != len(pattern):
            return False
        
        for idx, word in enumerate(words):
            p = pattern[idx]
            if p in m or word in m2:
                if m.get(p, None) != word:
                    return False
                if m2.get(word, None) != p:
                    return False
            else:
                m[p] = word
                m2[word] = p
        return True
