class Solution:
    def predictPartyVictory(self, senate: str) -> str:
        n = len(senate)
        rc = len(list(filter(lambda x: x == "R", senate)))
        dc = n - rc
        forbidden = [False] * n
        state = 0
        pos = 0

        while 1:
            t = senate[pos]
            if rc == 0:
                return "Dire"
            if dc == 0:
                return "Radiant"
            if forbidden[pos] == True:
                pos = (pos + 1) % n
                continue

            if t == "R":
                if state >= 0:
                    state += 1
                else:
                    state += 1
                    forbidden[pos] = True
                    rc -= 1
            else:
                if state <= 0:
                    state -= 1
                else:
                    state -= 1
                    forbidden[pos] = True
                    dc -= 1
            pos = (pos + 1) % n
