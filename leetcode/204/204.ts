function countPrimes(n: number): number {
    var isPrime = (Array<boolean>(n) as boolean[] & {fill(val :boolean)}).fill(true);
    var prime = [];

    for (var i=2; i<n; ++i) {
        if (isPrime[i]) prime.push(i);
        for (var j = 0; j < prime.length; ++j) {
        // for (var j in prime) {
            var target = prime[j] * i;
            if (target >= n) break;
            isPrime[target] = false;
            if ((i % prime[j]) == 0) break;
        }
    }
    
    return prime.length; 
};