import heapq

class Solution:
    def longestSubarray(self, nums: List[int], limit: int) -> int:
        n = len(nums)

        # 维护区间最大值最小值
        maxHeap = [-nums[0]]
        maxRemove = {} # 延迟删除
        minHeap = [nums[0]]
        minRemove = {} # 延迟删除

        def add(num):
            heapq.heappush(minHeap, num)
            heapq.heappush(maxHeap, -num)

        def remove(num):
            maxRemove[-num] = maxRemove.get(-num, 0) + 1
            minRemove[num] = minRemove.get(num, 0) + 1

        def getMax():
            while True:
                num = maxHeap[0]
                if maxRemove.get(num, 0) > 0:
                    heapq.heappop(maxHeap)
                    maxRemove[num] -= 1
                else:
                    return -num

        def getMin():
            while True:
                num = minHeap[0]
                if minRemove.get(num, 0) > 0:
                    heapq.heappop(minHeap)
                    minRemove[num] -= 1
                else:
                    return num

        # 滑动窗口
        i = 0
        res = 1
        while i + res < n:
            # 当前区间为 [i, i+res)
            mx = getMax()
            mn = getMin()
            # print(i, res, mx, mn)

            # 尝试向右扩充范围
            num = nums[i+res]
            if max(num, mx) - min(num, mn) > limit:
                # 新的区间不合法
                pass
            else:
                # 新区间合法
                res += 1
                add(num)
                continue

            # 右移
            remove(nums[i])
            add(num)
            i += 1

        return res
