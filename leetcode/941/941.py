class Solution:
    def validMountainArray(self, A: List[int]) -> bool:
        if len(A) < 3:
            return False
        up = True
        lst = A[1]
        if A[0] >= A[1]:
            return False
        for a in A[2:]:
            if up:
                if a > lst:
                    lst = a
                elif a == lst:
                    return False
                else:
                    # a < lst
                    lst = a
                    up = False
            else:
                if a >= lst:
                    return False
                else:
                    lst = a
        return up == False
