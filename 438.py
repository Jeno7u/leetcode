class Solution:
    def findAnagrams(self, s: str, p: str) -> List[int]:
        if len(p) > len(s):
            return []

        # get initial dict
        char_amount_s = {}
        char_amount_p = {}
        for i in range(len(p)):
            if s[i] not in char_amount_s:
                char_amount_s[s[i]] = 1
            else:
                char_amount_s[s[i]] += 1

            if p[i] not in char_amount_p:
                char_amount_p[p[i]] = 1
            else:
                char_amount_p[p[i]] += 1
        
        # do sliding window
        result = []
        i = 0
        while i < len(s) - len(p):
            if char_amount_s == char_amount_p:
                result.append(i)

            # remove start element
            if char_amount_s[s[i]] == 1:
                char_amount_s.pop(s[i])
            else:
                char_amount_s[s[i]] -= 1

            # add end + 1 elemnt
            if s[i + len(p)] not in char_amount_s:
                char_amount_s[s[i + len(p)]] = 1
            else:
                char_amount_s[s[i + len(p)]] += 1

            i += 1

        if char_amount_s == char_amount_p:
                result.append(i)

        return result

s = "cbaebabacd"
p = "abc"
solution = Solution()
print(solution.findAnagrams(s, p))
