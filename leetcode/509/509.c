int f[] = {
    0, 1, 1, 2, 3, 5, 8, 13, 
    21, 34, 55, 89, 144, 233, 
    377, 610, 987, 1597, 2584, 
    4181, 6765, 10946, 17711, 
    28657, 46368, 75025, 121393, 
    196418, 317811, 514229, 832040
};

int fib(int n){
    // int f[31] = { 0, 1 };
    // for (int i=2; i<=30; i++) {
        // f[i] = f[i-1] + f[i-2];
        // printf("%d, ", f[i]);
    // }
    return f[n];
}
