class Solution:
    def accountsMerge(self, acconts: List[List[str]]) -> List[List[str]]:
        n = len(acconts)
        
        pos = 0
        emailList = {}
        emailList2 = []
        for accont in acconts:
            for email in accont[1:]:
                if email not in emailList:
                    emailList2.append(email)
                    emailList[email] = pos
                    pos += 1
        f = [i for i in range(pos)]            
        def getF(x):
            f[x] = x if x == f[x] else getF(f[x])
            return f[x]
        def connect(x, y):
            f[getF(x)] = getF(y)

        u = ["" for i in range(pos)]
        for idx, accont in enumerate(acconts):
            username, emails = accont[0], accont[1:]

            index = emailList[emails[0]]
            u[index] = username

            for email in emails[1:]:
                index2 = emailList[email]
                u[index2] = username
                connect(index, index2)
        
        users = {}
        for i in range(pos):
            index = getF(i)
            email = emailList2[i]
            if index in users:
                users[index].append(emailList2[i])
            else:
                users[index] = [u[index], emailList2[i]]
        
        for k in users.keys():
            users[k] = [users[k][0]] + sorted(users[k][1:])

        return list(users.values())

