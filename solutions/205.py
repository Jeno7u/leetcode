class Solution:
    def isIsomorphic(self, s: str, t: str) -> bool:
        if len(t) != len(s):
            return False

        mapST = {}
        mapTS = {}
        for i in range(len(s)):
            if s[i] not in mapST and t[i] not in mapTS:
                mapST[s[i]] = t[i]
                mapTS[t[i]] = s[i]
            elif s[i] in mapST and t[i] == mapST[s[i]]:
                continue
            else:
                return False
        return True

s = "paper"
t = "title"
solution = Solution()
print(solution.isIsomorphic(s, t))