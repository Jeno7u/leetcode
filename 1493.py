from typing import List

class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        result = 0
        pos_0 = -1
        l, r = 0, 0
        while r < len(nums):
            if nums[r] != 1:
                if pos_0 == -1:
                    pos_0 = r
                else:
                    result = max(result, r - l - 1)
                    l = pos_0 + 1
                    pos_0 = r
            r += 1
        result = max(result, r - l - 1)

        return result


nums = [0, 0]
solution = Solution()
print(solution.longestSubarray(nums))