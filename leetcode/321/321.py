class Solution:   
    def maxNumber(self, nums1: List[int], nums2: List[int], k: int) -> List[int]:
        res = [0 for i in range(k)]
        
        m, n = len(nums1), len(nums2)
        start = max(0, k - n)
        end = min(k, m)
        
        for i in range(start, end + 1):
            s1 = self.getMax(nums1, i)
            s2 = self.getMax(nums2, k - i)
            s = self.merge(s1, s2)
            if self.compare(s, res) > 0:
                res = s

        return res

    def getMax(self, nums: List[int], k: int) -> List[int]:
        '''
        获取从 nums 选取 k 个数能拼成的最大序列
        '''
        stack = [0] * k
        top = -1
        unuse = len(nums) - k # 可以不使用的数目个数

        for num in nums: 
            while top >= 0 and stack[top] < num and unuse > 0:
                # 把栈顶比当前数小的都删掉
                # 尽可能保持高位最大（贪心）
                top -= 1
                unuse -= 1
            if top < k - 1:
                # 插入当前数
                top += 1
                stack[top] = num
            else:
                # 当前数未被使用
                unuse -= 1
        
        return stack


    def merge(self, s1: List[int], s2: List[int]) -> List[int]:
        '''
        将两个序列拼接
        '''
        l1 = len(s1)
        l2 = len(s2)
        i1 = 0
        i2 = 0
        pos = 0

        res = [0 for i in range(l1+l2)]
        while pos < l1+l2:
            if self.use1(s1,s2,i1,i2):
                res[pos] = s1[i1]
                i1 += 1
            else:
                res[pos] = s2[i2]
                i2 += 1
            pos +=1 
        return res
    
    def use1(self, s1: List[int], s2:List[int], i1: int, i2:int)->bool:
        '''
        在拼接序列时，是否应该从 s1 选取下一个值
        '''
        l1 = len(s1)
        l2 = len(s2)

        if i1>=l1 and i2>=l2:
            # s1 s2 都已经用完了，随便用谁返回下
            return True
        elif i1<l1 and i2>=l2:
            # s1 没用完，s2 用完，所以用 s1
            return True
        elif i1>=l1 and i2<l2:
            # s1 用完了，s2 没用完，所以用 s2
            return False

        if s1[i1] == s2[i2]:
            # 两个值一样，比较下一个
            return self.use1(s1, s2, i1+1, i2+1)
        elif s1[i1] > s2[i2]:
            # s1 的值更大，用 s1
            return True
        else:
            # s2 的值更大，用 s2
            return False

    def compare(self, s1: List[int], s2: List[int]) -> bool:
        '''
        比较序列大小
        '''
        for i in range(len(s1)):
            c1 = s1[i]
            c2 = s2[i]
            if c1 != c2:
                return c1 - c2
        return 0
