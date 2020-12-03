class Solution {
    public int countPrimes(int n) {
        int[] prime = new int[n];
        boolean[] isPrime = new boolean[n];
        int prime_num=0;

        Arrays.fill(isPrime, true);
        for (int i=2; i<n;++i) {
            if (isPrime[i])prime[prime_num++]=i;
            for (int j=0; j<prime_num; ++j) {
                int target = prime[j] * i;
                if (target >= n) break;
                isPrime[target] = false;
                if ((i % prime[j]) == 0) break;
            }
        }
        
        return prime_num; 
    }
}