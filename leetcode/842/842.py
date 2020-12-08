class Solution:
    def splitIntoFibonacci(self, S: str) -> List[int]:
        n = len(S)
        maxn = (1<<31)-1
        def dfs(state):
            '''
            state: 当前已有的状态
            []
            [123]
            [123,456]
            '''
            l = len(state)
            finished = sum(map(len, state))
            if (finished == n):
                return state
            
            if (l == 0 or l == 1):
                if (l == 1 and state[0][0] == "0" and int(state[0]) != 0):
                    return []
                for i in range(int((n - finished) / 2)):
                    res = dfs(state[:] + [S[finished:finished + i + 1]])
                    if (len(res) != 0):
                        return res
                return []
            else:
                # l >= 2
                # 取后两项，计算和
                if (int(state[-1]) > maxn):
                    return []
                target = sum(map(int, state[-2:]))
                if (target > maxn):
                    return []
                ltarget = len(str(target))
                if ltarget > n - finished:
                    return []
                if int(S[finished:finished+ltarget]) == target:
                    return dfs(state[:] + [S[finished:finished+ltarget]])
                return []

        return dfs([])
