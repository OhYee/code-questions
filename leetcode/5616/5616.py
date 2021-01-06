import queue

class Solution:
    def minimumDeviation(self, nums: List[int]) -> int:

        nums = [num << 1 if num & 1 else num for num in nums]
        nums.sort()

        q = queue.PriorityQueue()
        for num in nums:
            q.put(-num)
        
        min_number, max_number = min(nums), max(nums)
        res = max_number - min_number

        while 1:
            max_number = -q.get()
            res = min(res, max_number - min_number)
            
            if max_number & 1:
                break
            
            max_number = int(max_number / 2)
            q.put(-max_number)
            min_number = min(max_number, min_number)
            
        return res
