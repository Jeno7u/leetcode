class Solution:
    def longestMonotonicSubarray(self, nums: list[int]) -> int:
        res = 1
        curr = 1
        increasing = 0
        for i in range(1, len(nums)):
            if nums[i - 1] < nums[i]:
                if increasing > 0:
                    curr += 1
                else:
                    curr = 2
                    increasing = 1
            elif nums[i - 1] > nums[i]:
                if increasing < 0:
                    curr += 1
                else:
                    curr = 2 
                    increasing = -1
            else:
                increasing = 0
                curr = 1
            res = max(res, curr) 

        return max(res, curr)


nums = [1,4,3,3,2]
solution = Solution()
print(solution.longestMonotonicSubarray(nums))