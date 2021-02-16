class Solution:
	def matrixReshape(self, nums: List[List[int]], r: int, c: int) -> List[List[int]]:
	n = len(nums)
	if n == 0:
	return nums
	m = len(nums[0])
	if n*m != r*c:
	return nums

	res = [[ 0 for j in range(c) ] for i in range(r) ]
	for i in range(r):
		for j in range(c):
			idx = i * c + j
			res[i][j] = nums[idx // m] [idx % m]
			return res
