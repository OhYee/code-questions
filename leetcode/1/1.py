class Node:
    def __init__(self, num, idx):
        self.num = num
        self.idx = idx

    def __repr__(self):
        return "Node({}, {})".format(self.num, self.idx)

    def get_value(self, other):
        other_value = 0
        if type(other) == int:
            other_value = other
        elif isinstance(other, Node):
            other_value = other.num
        return other_value

    def __lt__(self, other):
        return self.num < self.get_value(other)

    def __eq__(self, other):
        return self.num == self.get_value(other)


def lower_bound(_list, _target):
    idx = bisect.bisect_left(_list, _target)
    if idx >= len(_list) or _list[idx] != _target:
        idx = -1
    return idx


class Solution:
    def twoSum(self, nums: list, target: int):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        sorted_pair = [Node(num, idx) for (idx, num) in enumerate(nums)]
        sorted_pair.sort()
        res = []
        for pair in sorted_pair:
            surplus = target - pair.num
            idx = lower_bound(sorted_pair, Node(surplus, 0))
            if idx != -1 and sorted_pair[idx].idx != pair.idx:
                res = sorted([pair.idx, sorted_pair[idx].idx])
                break
        return res
