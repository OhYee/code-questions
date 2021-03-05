import re

class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        return re.match("^%s$" % p, s) != None
