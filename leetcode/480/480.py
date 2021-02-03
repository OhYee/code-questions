import bisect

class Solution:
    def medianSlidingWindow(self, nums: List[int], k: int) -> List[float]:
        n = len(nums)
        heap = sorted(nums[:k])

        def getMidOdd():
            return heap[k // 2]
        def getMidEven():
            mid = k//2
            return (heap[mid] + heap[mid-1]) /2 
        getMid = getMidEven if k % 2 == 0 else getMidOdd

        def pop(num):
            del heap[bisect.bisect_left(heap, num)]

        def push(num):
            bisect.insort(heap, num)            
        mids = [0 for i in range(n-k+1)]
        for l in range(n-k+1):
            mids[l] = getMid()
            r = l + k
            if r < n:
                pop(nums[l])
                push(nums[r])
        return mids

