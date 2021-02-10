class Solution:
    def checkInclusion(self, s1: str, s2: str) -> bool:
        n = len(s1)
        m = len(s2)
        if n > m:
            return False

        count = [0 for i in range(26)]
        target = [0 for i in range(26)]
        def add(c, l=count):
            # print("add", c)
            l[ord(c) - ord('a')] += 1
        def remove(c, l=count):
            # print("remove", c)
            l[ord(c) - ord('a')] -= 1
        def check():
            # print(target)
            # print(count)
            # print("-"*10)
            for i in range(26):
                if count[i] != target[i]:
                    return False
            return True

        for i in range(n):
            add(s2[i])
            add(s1[i], l=target)

        if check():
            return True

        for i in range(1, m-n+1):
            remove(s2[i-1])
            add(s2[i+n-1])
            if check():
                return True
        return False
