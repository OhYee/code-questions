int cmp(const void *a, const void *b) {
   return (*(int*)b) - (*(int*)a);
}

int largestPerimeter(int* A, int ASize) {
    qsort(A, ASize, sizeof(int), cmp);
    for (int i=2; i<ASize; ++i) 
        if (A[i-2] < A[i-1] + A[i])
            return A[i] + A[i-1] + A[i-2];
    return 0;
}