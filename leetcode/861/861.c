int matrixScore(int** A, int ASize, int* AColSize){
    for (int i=0; i<ASize; ++i)
        if (A[i][0] == 0) 
            for (int j=0; j<AColSize[i]; ++j) 
                A[i][j] = !A[i][j];

    // for (int i=0; i<ASize; ++i) {
    //     for (int j=0; j<AColSize[i]; ++j)
    //         printf("%d ", A[i][j]);
    //     printf("\n");
    // }
    // printf("\n");

    for (int j=0; j<AColSize[0]; ++j) {
        int num1 = 0;
        for (int i=0; i<ASize; ++i) 
            if (A[i][j]) ++num1;
        if (num1*2 < ASize)
            for (int i=0; i<ASize; ++i) 
                A[i][j] = !A[i][j];
    }
    
    // for (int i=0; i<ASize; ++i) {
    //     for (int j=0; j<AColSize[i]; ++j)
    //         printf("%d ", A[i][j]);
    //     printf("\n");
    // }
    // printf("\n");

    int res = 0;
    for (int i=0; i<ASize; ++i)
        for (int j=0; j<AColSize[i]; ++j)
            res += (A[i][j] << (AColSize[i]-j-1));
    return res;
}