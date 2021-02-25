class Solution:
    def transpose(self, matrix: List[List[int]]) -> List[List[int]]:
        n = len(matrix)
        m = len(matrix[0])
        return [
            [
                matrix[j][i]
                for j in range(n)
            ]
            for i in range(m)
        ]
