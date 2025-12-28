class Solution:
    def validPalindrome(self, s: str) -> bool:
        def val(l, r):
            while l < r:
                if s[l] != s[r]:
                    return False
                l, r = l + 1, r - 1
            return True

        l, r = 0, len(s) - 1
        while l < r:
            if s[l] != s[r]:
                return val(l + 1, r) or val(l, r - 1)
            l, r = l + 1, r - 1
        return True
    
s = "abc"
solution = Solution()
print(solution.validPalindrome(s))