class Solution:
    def threeEqualParts(self, A: List[int]) -> List[int]:
        A = list(map(int, A))
        l = len(A)
        if l == 0:
            return [0, -1]
        n1 = len(list(filter(lambda x: x==1, A)))
        if n1 == 0:
            return [0, 4]
        if n1 % 3 != 0:
            return [-1, -1]
        n = int(n1 / 3)
        # 0001000100010010000101000   2 -> 6
        #         ^^^    ^^^^
        #         p0 p1  p2  p3

        p = [0,0,0,0,0]
        count = 0
        for idx, a in enumerate(A):
            if a == 1:
                count += 1
                if count == n:
                    p[0] = idx
                if count == n+1:
                    p[1] = idx
                if count == n*2:
                    p[2] = idx
                if count == n*2+1:
                    p[3] = idx
                p[4] = idx
        
        # 最后一个区域的后导零是确认的，共有 l-p3-1 个后导零
        # 故前两个区域也应该保证只有 l-p3-1 个后导零
        # 00000100100
        # 000100100
        # 00000100100
        zero = l-p[4]-1

        print(p)
        #      p0           p1      p2        p3
        # .... 1 0 0 0 0 0 0 1 .....1 0 0 0 0 1 ....
        #          i|i+1              j-1|j
        if p[0] + zero >= p[1] or p[2] + zero >= p[3]:
            return [-1, -1]

        i = p[0] + zero
        j = p[2] + zero + 1

        if i+1 >= j:
            return [-1, -1]

        s0 = int("".join(map(str, A[:i+1])), 2)
        s1 = int("".join(map(str, A[i+1:j])), 2)
        s2 = int("".join(map(str, A[j:])), 2)
        if s0 == s1 and s1 == s2:
            return [i, j]

        return [-1, -1]
