class Solution:
    def lemonadeChange(self, bills: List[int]) -> bool:
        m5 = 0
        m10 = 0
        m20 = 0
        for b in bills:
            if b == 5:
                m5 += 1
            elif b == 10:
                m10 += 1
                if m5 >= 1:
                    m5 -= 1
                else:
                    return False
            else:
                m20 += 1
                if m10 >= 1 and m5 >= 1:
                    m10 -= 1
                    m5 -= 1
                elif m5 >= 3:
                    m5 -= 3
                else:
                    return False
        return True
