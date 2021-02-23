class Solution:
    def flipAndInvertImage(self, A: List[List[int]]) -> List[List[int]]:
        return [reversed([item ^ 1 for item in line]) for line in A]
