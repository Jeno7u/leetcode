from typing import List

class Solution:
    def intervalIntersection(self, firstList: List[List[int]], secondList: List[List[int]]) -> List[List[int]]:
        res = []
        p1, p2 = 0, 0
        while p1 < len(firstList) and p2 < len(secondList):
            if firstList[p1][1] < secondList[p2][0]:
                p1 += 1
            elif secondList[p2][1] < firstList[p1][0]:
                p2 += 1
            else:
                if firstList[p1][0] < secondList[p2][0]:
                    if firstList[p1][1] < secondList[p2][1]:
                        res.append([secondList[p2][0], firstList[p1][1]])
                        p1 += 1
                    else:
                        res.append(secondList[p2])
                        p2 += 1
                else:
                    if firstList[p1][1] > secondList[p2][1]:
                        res.append([firstList[p1][0], secondList[p2][1]])
                        p2 += 1
                    else:
                        res.append(firstList[p1])
                        p1 += 1
        return res


firstList = [[0,2],[5,10],[13,23],[24,25]]
secondList = [[1,5],[8,12],[15,24],[25,26]]

solution = Solution()
print(solution.intervalIntersection(firstList, secondList))


firstList = [[17,20]]
secondList = [[2,3],[6,8],[12,14],[19,20]]

solution = Solution()
print(solution.intervalIntersection(firstList, secondList))


# TODO the solution could be shrinked using max and min + simple check 
# a_start, a_end = A[i]
# b_start, b_end = B[j]
# if a_start <= b_end and b_start <= a_end:                       # Criss-cross lock
#     result.append([max(a_start, b_start), min(a_end, b_end)])   # Squeezing