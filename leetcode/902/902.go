
import (
	"sort"
	// "strconv"
)

// func pow(a, n int) int {
// 	switch n {
// 	case 0:
// 		return 1
// 	case 1:
// 		return a
// 	case 2:
// 		return a * a
// 	default:
// 		res := pow(a, n/2)
// 		res *= res
// 		if n%2 == 1 {
// 			res *= a
// 		}
// 		return res
// 	}

// }

// func atMostNGivenDigitSet2(digits []string, n int) int {
// 	nn := len(digits)
// 	nums := make([]int, nn)
// 	for i := 0; i < nn; i++ {
// 		temp, _ := strconv.ParseInt(digits[i], 10, 32)
// 		nums[i] = int(temp)
// 	}
// 	sort.Ints(nums)

// 	res := 0

// 	var dfs func(l, num int) bool
// 	dfs = func(l, num int) bool {
// 		if l == 0 {
// 			if num <= n {
// 				res++
// 				// fmt.Println(num)
// 				return true
// 			}
// 			return false
// 		}
// 		for _, cur := range nums {
// 			if !dfs(l-1, num*10+cur) {
// 				return false
// 			}
// 		}
// 		return true
// 	}

// 	l := 0
// 	temp := n
// 	for temp > 0 {
// 		temp /= 10
// 		l++
// 	}
// 	temp = 0
// 	for i := 1; i <= l; i++ {
// 		temp = temp*10 + 9
// 		if n >= temp {
// 			res += pow(nn, i)
// 		} else {
// 			dfs(i, 0)
// 		}
// 	}

// 	return res
// }

func s2i(s string) int {
	switch s {
	case "1":
		return 1
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	default:
		return 1
	}
}

func atMostNGivenDigitSet(digits []string, n int) int {
	nn := len(digits)
	nums := make([]int, nn)
	for i := 0; i < nn; i++ {
		nums[i] = s2i(digits[i])
	}

	// 获取位数
	target := make([]int, 0)
	temp := n
	for temp > 0 {
		target = append([]int{temp % 10}, target...)
		temp /= 10
	}
	m := len(target)

	sort.Ints(nums)

	// 0: 不选，共 1 种			0->0  0->3
	// 1: 选比当前数小的		1->3
	// 2: 选和当前数相等的		2->1  2->2
	// 3： 全选，有 nn 种	   3->3
	dp := make([][]int, 2)
	dp[0] = make([]int, 4)
	dp[1] = make([]int, 4)

	for i := 0; i < m; i++ {
		less := 0
		equal := 0
		for j := 0; j < nn; j++ {
			if nums[j] > target[i] {
				break
			}
			if nums[j] == target[i] {
				equal++
			} else {
				less++
			}
		}
		if i == 0 {
			dp[1][0] = 1
			dp[1][1] = less
			dp[1][2] = equal
			dp[1][3] = 0
		} else {
			dp[1][0] = 1 * dp[0][0]
			dp[1][1] = less * dp[0][2]
			dp[1][2] = equal * dp[0][2]
			dp[1][3] = nn * (dp[0][3] + dp[0][1] + dp[0][0])
		}

		// fmt.Println(dp[1])

		dp[0], dp[1] = dp[1], dp[0]
	}

	return dp[0][0] + dp[0][1] + dp[0][2] + dp[0][3] - 1
}
