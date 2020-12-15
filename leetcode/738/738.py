class Solution:
    def monotoneIncreasingDigits(self, N: int) -> int:        
        nums = list(map(int, list(str(N))))
        l = len(nums)
        fallback = -1
        for idx, n in enumerate(nums):
            if idx == 0:
                continue
            if nums[idx-1] > n:
                # 发现非递增数
                fallback = idx
                break

        if fallback == -1:
            return N
        else:
            # 后续的全部改为 9，上一个数减一，回溯
            for i in range(fallback, l):
                nums[i] = 9

            nums[fallback-1] -= 1

            for i in range(fallback-1, -1, -1):
                if i != 0 and nums[i-1] > nums[i]:
                    nums[i] = 9
                    nums[i-1] -= 1
            
            return int("".join(map(str, nums)))
