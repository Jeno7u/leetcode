# sliding window problem, но есть два варианта. Закидывать посещенные в set и добавлять когда встретили, но тогда чтобы понять когда 
# мы последний раз встречали этот элемент мы должны его еще найти. Так что лучше использовать словарь и туда запихивать когда последний
# раз мы видели этот элемент чтобы с него начать наш window
class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if len(s) == 0:
            return 0

        result = 1
        seen_chars = {}
        start, i = 0, 0
        while i < len(s):
            print("------")
            print(start, i, result)
            print(s[start:i])
            if s[i] not in seen_chars:
                seen_chars[s[i]] = i
            elif s[i] in seen_chars:
                if start <= seen_chars[s[i]] <= i:
                    start = seen_chars[s[i]] + 1
                seen_chars[s[i]] = i
            result = max(result, i - start + 1)
            i += 1
        return result

s = "ab!bb"
solution = Solution()
print(solution.lengthOfLongestSubstring(s))