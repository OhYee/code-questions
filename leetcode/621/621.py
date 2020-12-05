class Solution:
    def leastInterval(self, tasks: List[str], n: int) -> int:
        count = {}
        for t in tasks:
            count[t] = count.get(t, 0) + 1
        items = sorted(count.items(), key=lambda x: x[1], reverse=True)
        total = len(tasks)
        res = 0
        finish = {}
        while True:
            if len(items) > 0:
                t = -1
                for idx, item in enumerate(items):
                    k, v = item
                    if finish.get(k, -1) < res:
                        t = idx
                        break
                if t == -1:
                    # 该轮没有任务可执行
                    # 等待一轮
                    res += 1
                else:
                    # 该轮执行第 t 个任务
                    k, v = items[t]
                    finish[k] = res + n
                    items[t] = (k, v-1)
                    if v-1 <= 0:
                        items.pop(t)
                    else:
                        for i in range(t + 1, len(items)):
                            if items[i][1] >= items[i-1][1]:
                                items[i], items[i-1] = items[i-1], items[i]
                            else:
                                break
                    res += 1
            else:
                return res
