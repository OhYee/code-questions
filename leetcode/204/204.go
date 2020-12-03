package main

func countPrimes(n int) int {
	isPrime := make([]bool, n)
	prime := make([]int, 0)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			prime = append(prime, i)
		}
		for _, j := range prime {
			target := i * j
			if target >= n {
				break
			}
			isPrime[target] = false
			if (i % j) == 0 {
				break
			}
		}
	}

	return len(prime)
}
