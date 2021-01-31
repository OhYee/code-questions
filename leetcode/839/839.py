class Solution:
    def numSimilarGroups(self, strs: List[str]) -> int:
        n = len(strs)

        f = [i for i in range(n)]
        def getF(x):
            f[x] = x if x == f[x] else getF(f[x])
            return f[x]
        def connect(x, y):
            f[getF(x)] = getF(y)

        def isSimilar(s1, s2):
            if s1 == s2:
                # 如果两个字符串是相等的，那么他们也是相似的
                return True

            l = len(s1)
            checkOK = False
            p = -1
            for i in range(l):
                if s1[i] != s2[i]:
                    if checkOK:
                        # 只允许有一组不同
                        return False
                    if p == -1:
                        # 记录第一次不同的位置
                        p = i
                    else:
                        # 判断是否为交换关系
                        if s1[i] == s2[p] and s1[p] == s2[i]:
                            checkOK = True
                        else:
                            return False
            return True

        for i in range(n):
            for j in range(i+1, n):
                if isSimilar(strs[i], strs[j]):
                    connect(i, j)
        
        res = 0
        for i in range(n):
            if getF(i) == i:
                res += 1

        return res

