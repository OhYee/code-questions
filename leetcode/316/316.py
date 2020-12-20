class Solution:
    def removeDuplicateLetters(self, s: str) -> str:
        n = len(s)
        l = len(set(s)) # 最终的字符串长度
        stack = l * [' ']
        top = -1

        # 不同字符个数
        cnt = {}
        for c in s:
            cnt[c] = cnt.get(c, 0) + 1

        for c in s:
            # print("当前栈", stack[:top+1], " 判断", c)
            while top >= 0 and stack[top] > c and c not in stack[:top+1]:
                suffix = len(list(map(
                    lambda x: x[1], 
                    filter(
                        lambda x:x[0] not in stack[:top] and x[1] > 0,
                        cnt.items()
                    )
                ))) # 后面可选的字符个数
                # print(" 可选字符还有", suffix)
                if top + 1 + suffix <= l:
                    break 
                top -= 1
                # print(" 出栈", stack[:top+1])                
            if top < l - 1 and c not in stack[:top+1]:
                top += 1
                stack[top] = c
                # print(" 入栈", stack[:top+1])
            else:
                # print(" 不插入")
                pass

            cnt[c] -= 1
            if cnt[c] <= 0:
                del cnt[c]

        return "".join(stack)
