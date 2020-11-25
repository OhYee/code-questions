class Solution:
    def sortString(self, s: str) -> str:
        lst = list(s)
        res = ""

        reverse = False
        while len(lst) > 0:
            part = sorted(list(set(lst)), reverse=reverse)
            reverse = not reverse

            for c in part:
                lst.remove(c)
            res += "".join(part)
            
        return res
