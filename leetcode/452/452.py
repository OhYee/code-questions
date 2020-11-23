from functools import cmp_to_key


def log(*args, **kwargs):
    # print(*args, **kwargs)
    pass


class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        points.sort(key=cmp_to_key(
            lambda a, b: (
                (- 1 if a[0] < b[0] else 1)
                if a[0] != b[0] else
                (-1 if a[1] < b[1] else 1)
            )
        ))
        log(points)

        count = 0
        lastX, lastY = -1, -1
        skip = False
        for x, y in points:
            if lastX == -1 and lastY == -1:
                lastX, lastY = x, y
            else:
                if x >= lastX and y <= lastY:
                    # 新的区间完全被老的区间覆盖
                    # 不需要考虑老区间，新区间删除的时候会同时删除老区间
                    log("覆盖，更新", lastX, lastY, x, y)
                    lastX, lastY = x, y

                elif x == lastX and y >= lastY:
                    # 老区间完全被新区覆盖
                    # 不需要考虑新区间
                    log("覆盖，不更新", lastX, lastY, x, y)
                    pass
                elif x >= lastX and x <= lastY and y > lastY:
                    # 新的和老的存在交叉
                    # 更新区间为交叉部分
                    log("交叉", lastX, lastY, x, y)
                    lastX, lastY = x, lastY
                else:
                    # 新的和老的没交叉，老的可以直接删除掉并计数
                    log("无交叉+", lastX, lastY, x, y)
                    count += 1
                    lastX, lastY = x, y
        if lastX != -1 and lastY != -1:
            log("结束+", lastX, lastY)
            count += 1

        return count
