class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if s == "":
            return True
        pos = 0
        for i in range(len(t)):
            if t[i] == s[pos]:
                pos += 1
            if len(s) == pos:
                return True
        return len(s) == pos

s = "abc"
t = "ahbgdc"

solution = Solution()
print(solution.isSubsequence(s, t))