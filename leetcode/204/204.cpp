#include <cstring>

class Solution {
  public:
    int countPrimes(int n) {
        int * prime = new int[n];
        bool *isPrime = new bool[n];
        int   prime_num = 0;

        memset(isPrime, true, sizeof(bool) * n);
        for (int i = 2; i < n; ++i) {
            if (isPrime[i])
                prime[prime_num++] = i;
            for (int j = 0; j < prime_num; ++j) {
                int target = prime[j] * i;
                if (target >= n)
                    break;
                isPrime[target] = false;
                if (!(i % prime[j]))
                    break;
            }
        }

        delete[] prime;
        delete[] isPrime;
        return prime_num;
    }
};