#include <stdbool.h>

int countPrimes(int n) {
    int * prime = malloc(sizeof(int) * n);
    bool *isPrime = malloc(sizeof(bool) * n);
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
            if ((i % prime[j]) == 0)
                break;
        }
    }

    free(prime);
    free(isPrime);
    return prime_num;
}