class Solution:
    def findNumOfValidWords(self, words: List[str], puzzles: List[str]) -> List[int]:
        def w2n(s):
            num = 0
            for c in s:
                num |= 1<<(ord(c) - ord('a'))
            return num

        m = {}
        for word in words:
            if len(ws := set(word)) <= 7:
                w = w2n(ws)
                m[w] = m.get(w, 0) + 1

        def getRes(puzzle):
            count = 0
            lst = [ord(c) - ord('a') for c in puzzle]
            # 枚举所有可能的谜面
            for i in range(1<<6):
                temp = 1 << lst[0]
                for j in range(6):
                    temp |= (((i>>j) & 1) << lst[j+1])
                count += m.get(temp, 0)
            return count
        return [getRes(puzzle) for puzzle in puzzles]
