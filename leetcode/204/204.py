class Solution:
    def countPrimes(self, n: int) -> int:
        prime = []
        isPrime = [True] * n

        for i in range(2, n): 
            if isPrime[i]: 
                prime.append(i)
            for j in prime:
                target = i * j
                if target >= n:
                    break
                isPrime[target] = False
                if (i % j) == 0:
                    break
        
        return len(prime)