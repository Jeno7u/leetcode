from typing import List

# вариает с поиском ближайшего числа через бинарный поиск и двумя поинтерами ищем бижайшие числа
class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        if len(arr) == k:
            return arr
        
        def find_idx_x(arr, x):
            # binary search
            l, r = 0, len(arr) - 1
            while l <= r:
                mid = (l + r) // 2
                if arr[mid] == x:
                    return mid
                if arr[mid] < x:
                    l = mid + 1
                if arr[mid] > x:
                    r = mid - 1
            
            if l == len(arr):
                l -= 1

            # closest value
            if 0 < l < len(arr) and abs(x - arr[l - 1]) <= abs(x - arr[l]):
                l -= 1
            elif 0 <= l < len(arr) - 1 and abs(x - arr[l + 1]) < abs(x - arr[l]):
                l += 1
            return l
        
        idx = find_idx_x(arr, x)

        l, r = idx, idx
        k -= 1
        while k > 0:
            if (r == len(arr) - 1) or (l > 0 and r < len(arr) - 1 and (abs(x - arr[l - 1]) <= abs(x - arr[r + 1]))):
                l -= 1
                k -= 1
            elif (r < len(arr) - 1):
                r += 1
                k -=1

        return arr[l:r + 1] 

arr = [0,0,1,2,3,3,4,7,7,8]
k = 3
x = 5
solution = Solution()
print(solution.findClosestElements(arr, k, x))


# и нормальный вариант где мы делаем lower bound бинарный поиск где ищем срез (ответ) проверяя крайние точки среза на близость к нужному числу
# типо если левая граница ближе чем правая, то движемся влево
class Solution:
    def findClosestElements(self, arr: List[int], k: int, x: int) -> List[int]:
        l, r = 0, len(arr) - k
        while l < r:
            mid = (l + r) // 2
            if x - arr[mid] > arr[mid + k] - x:
                l = mid + 1
            else:
                r = mid
        return arr[l:l + k]