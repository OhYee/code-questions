class Solution:
    def subarraysWithKDistinct(self, A: List[int], K: int) -> int:
        n = len(A)

        count = [0 for i in range(n+1)]
        diff = 0

        def add(num):
            '''
            向目前区间添加一个数字
            '''
            nonlocal diff
            count[num] += 1
            if count[num] == 1:
                diff += 1
            # print("add ",num, "diff: ", diff)
        def remove(num):
            '''
            从目前区间移除一个数字
            '''
            nonlocal diff
            count[num] -= 1
            if count[num] == 0:
                diff -= 1
            # print("remove ",num, "diff: ", diff)

        i = 0
        j = 0
        res = 0
        # 从 [0,0] 区间开始
        add(A[0])
        if diff == K:
            res += 1
            print(A[i:j+1])

        while j + 1 < n:
            j += 1
            add(A[j])
            while diff > K:
                remove(A[i])
                i += 1
            if diff == K:
                res += 1
                # print(A[i:j+1])

                start = i
                while diff == K:
                    remove(A[i])
                    i += 1
                    if diff == K:
                        res += 1
                        # print(A[i:j+1])
                    else:
                        break
                while i > start:
                    i -= 1
                    add(A[i])
            
        return res                

