from typing import List

class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        def find_idx_x(arr, x):
            l, r = 0, len(arr) - 1
            while l <= r:
                mid = (l + r) // 2
                if arr[mid] == x:
                    return mid
                if arr[mid] < x:
                    l = mid + 1
                else:
                    r = mid - 1
            return l
        
        idx = find_idx_x(arr, x)
        radius = k // 2
        dr = 0
        
        if len(arr) > idx + radius:
            dr = len(arr)

arr = [1,2,3,4,5]
k = 4
x = 4

solution = Solution()
print(solution.findClosestElements(arr, k, x))