int *f;

int getF(int t) {
    return f[t] = (t == f[t] ? t : getF(f[t]));
}

int findCircleNum(int** isConnected, int isConnectedSize, int* isConnectedColSize){
    int n = isConnectedSize;

    f = malloc(sizeof(int) * n);
    for (int i=0; i<n; ++i) f[i] = i;
    
    for (int i=0; i<n; ++i) { 
        for (int j=0; j<i; ++j) {
            if (isConnected[i][j]) {
                f[i] = f[j] = f[getF(j)] = getF(i);
            }
        }
    }
    int count = 0;
    for (int i=0; i<n; ++i) {
        // printf("%d %d\n", i, getF(i));
        if (getF(i) == i) ++count;
    }

    free(f);
    return count;
}
