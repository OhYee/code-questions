class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        pre = 0
        res = 0
        for f in flowerbed:
            if pre == 0:
                if f == 0:
                    # 0 0 
                    # 当前格子可以种
                    res += 1
                    pre = 1
                else:
                    # 0 1
                    # 当前格子不能种
                    pre = 1
            else:
                if f == 0:
                    # 1 0
                    # 当前格子不能种
                    pre = 0
                else:
                    # 1 1
                    # 当前格子不能种，上一个格子也不能种
                    res -= 1
                    pre = 1
        return res >= n
