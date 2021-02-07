class Solution:
    def checkPossibility(self, nums: List[int]) -> bool:
        nums = [-9999999] + nums
        n = len(nums)
        flag = False
        for i in range(1, n):
            if nums[i-1] <= nums[i]:
                # 符合要求
                pass
            else:
                # 不合要求
                if nums[i-2] <= nums[i]:
                    # 当前比上上个大，变上一个
                    if flag:
                        return False
                    flag = True
                    nums[i-1] = nums[i-2]
                else:
                    # 当前比上上个小，变这个
                    if flag:
                        return False
                    flag = True
                    nums[i] = nums[i-1]
        return True

