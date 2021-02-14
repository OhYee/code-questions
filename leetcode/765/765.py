class Solution:
    def minSwapsCouples(self, row: List[int]) -> int:
        n = len(row)
        row = [row[i] // 2 for i in range(n)]
        pos = [[-1, -1] for i in range(n//2)]
        for idx, r in enumerate(row):
            if pos[r][0] == -1:
                pos[r][0] = idx
            else:
                pos[r][1] = idx
        
        res = 0
        for idx in range(n // 2):
            p1 = idx * 2
            p2 = p1 + 1

            if row[p1] != row[p2]:
                res += 1

                # 确定另一名情侣坐标，并交换
                target = row[p1]
                swapPos = pos[target][0]
                if swapPos <= p2:
                    swapPos = pos[target][1]
                row[swapPos], row[p2] = row[p2], row[swapPos]

                # 维护情侣位置信息
                pos[target] = [p1, p2]
                target = row[swapPos]
                if pos[target][0] == p2:
                    pos[target][0] = swapPos
                else:
                    pos[target][1] = swapPos
        # print(pos)
        return res
                                

            

