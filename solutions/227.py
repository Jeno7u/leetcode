# у на есть стек. Если число закидываем в стек. Если нужно сделать вычитание закидываем число, но отрицаиельное. А в конце возвращаем
# просто сумму стека. Но если у нас деление, то мы вычесляем сразу значение умножения/деления и закидываем в стек.
class Solution:
    def calculate(self, s):    
        def update(op, v):
            if op == "+": stack.append(v)
            if op == "-": stack.append(-v)
            if op == "*": stack.append(stack.pop() * v)
            if op == "/": stack.append(int(stack.pop() / v))
    
        num, stack, sign = 0, [], "+"
        
        it = 0
        while it < len(s):
            if s[it].isdigit():
                num = num * 10 + int(s[it])
            elif s[it] in "+-*/":
                update(sign, num)
                num, sign = 0, s[it]
            it += 1
        update(sign, num)
        return sum(stack)
        
        

s = " 3+5 / 2 "
solution = Solution()
print(solution.calculate(s))