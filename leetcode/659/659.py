import sys


class Solution:
    def isPossible(self, nums: List[int]) -> bool:
        if len(nums) == 0:
            return True
        count = {}
        for n in nums:
            count[n] = count.get(n, 0) + 1

        pre = -1
        l = [0, 0, 0, 0]
        for i in range(min(nums), max(nums) + 1):
            if pre == -1:
                # 没有前一个数
                if count.get(i, 0) > 0:
                    # 当前数有值，则更新为当前数
                    l[1] = count[i]
                    l[2] = 0
                    l[3] = 0
                    pre = i
                else:
                    continue
            else:
                if count.get(i, 0) == 0:
                    # 存在前一个数，同时当前数不存在
                    if l[1] != 0 or l[2] != 0:
                        # 还有没凑够 3 个的数
                        return False
                    else:
                        # 全部都凑够 3 个，从新开始
                        l[3] = 0
                        pre = -1
                        continue
                else:
                    c = count[i]
                    temp = l[3]
                    # 优先凑 3 个
                    if c >= l[2]:
                        # 足够把 2 个的都凑成 3 个
                        c -= l[2]
                        l[3] = l[2]
                        if c >= l[1]:
                            # 足够把所有 1 个的凑成 2 个
                            c -= l[1]
                            l[2] = l[1]
                            if c > temp:
                                # 剩下的都补给原本就有 3 个的仍有剩余
                                l[1] = c - temp
                                l[3] += temp
                            else:
                                # 剩下的都补给原本就有 3 个的
                                l[1] = 0
                                l[3] += c
                        else:
                            # 只能把部分 1 个凑成 2 个
                            return False
                    else:
                        # 只能把部分 2 个凑成 3 个
                        return False
            print(i, l[1], l[2], l[3], file=sys.stderr)
        return l[1] == 0 and l[2] == 0
