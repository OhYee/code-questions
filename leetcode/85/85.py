class Solution:
    def maximalRectangle(self, matrix: List[List[str]]) -> int:
        n = len(matrix)
        if n == 0:
            return 0
        m = len(matrix[0])
        
        res = 0
        h = [0] * m
        left  = [0] * m
        right = [0] * m
        
        for i in range(n):
            # 更新以当前行为底的每个小矩形的高
            for j in range(m):
                if matrix[i][j] == '1':
                    h[j] += 1
                else:
                    h[j] = 0
            # 将每个小矩形向两侧扩展成矩形
            for j in range(m):
                t = j - 1
                while t >= 0 and h[t] >= h[j]:
                    if h[t] == h[j]:
                        t = left[t]
                        break
                    else:
                        t -= 1
                left[j] = t 
            for j in range(m-1, -1, -1):
                t = j + 1
                while t <= m-1 and h[t] >= h[j]:
                    if h[t] == h[j]:
                        t = right[t]
                        break
                    else:
                        t += 1
                right[j] = t 
            # print(i, "h", h)
            # print(i, "l", left)
            # print(i, "r", right)
            for j in range(m):
                res = max(res, h[j] * (right[j] - left[j] - 1))
        return res
                

            
