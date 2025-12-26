from typing import List

class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        result = []
        i = 0
        while i < len(nums):
            start = i
            j = i
            while j < len(nums) - 1 and nums[j] + 1 == nums[j + 1]:
                j += 1
            
            if start == j:
                result.append(str(nums[j]))
            else:
                result.append(f"{nums[start]}->{nums[j]}")
            
            i = j + 1
            
        return result    

nums = [0,2,3,4,6,8,9]
solution = Solution()
print(solution.summaryRanges(nums))