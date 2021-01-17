class Solution:
    def checkStraightLine(self, coordinates: List[List[int]]) -> bool:
        if len(coordinates) <= 1:
            return true

        x0, y0 = coordinates[0]
        x1, y1 = coordinates[1]

        for p in coordinates[2:]:
            x, y = p
            if (y-y0) * (x-x1) != (y-y1) * (x-x0):
                return False
        return True

