from typing import List

class Solution:
    def compress(self, chars: List[str]) -> int:
        chars.append("end")
        compressed_pos = 0
        count = 1
        current_char = chars[0]
        for i in range(1, len(chars)):
            if current_char != chars[i]:
                if count == 1:
                    chars[compressed_pos] = current_char
                    compressed_pos += 1
                    
                else:
                    chars[compressed_pos] = current_char
                    compressed_pos += 1
                    number = str(count)
                    number_length = len(number)
                    number_i = 0
                    while number_length > 0:
                        chars[compressed_pos] = number[number_i]
                        number_length -= 1
                        number_i += 1
                        compressed_pos += 1
                    count = 1
                current_char = chars[i]
            else:
                count += 1
        return compressed_pos

chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
solution = Solution()
result = solution.compress(chars)
print(result, chars[:result])
