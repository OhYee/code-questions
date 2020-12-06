class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        res = [[1] for i in range(numRows)] 
        for i in range(1, numRows):
            pre = res[i-1]
            for j in range(i-1):    
                res[i].append(pre[j] + pre[j+1])
            res[i].append(1)
        return res