void rotate(int** matrix, int matrixSize, int* matrixColSize){
    int n = matrixSize;
    for (int i=0; i<n/2; ++i) {
        for (int j=i; j<n-i-1; ++j) {
            // i=0 j=1 n=4
            //  0,0 [0,1] 0,2  0,3
            //  1,0  1,1  1,2 [1,3]
            // [2,0] 2,1  2,2  2,3
            //  3,0  3,1 [3,2] 3,3
            int temp = matrix[i][j];
            matrix[i][j] = matrix[n-j-1][i];
            matrix[n-j-1][i] = matrix[n-i-1][n-j-1];
            matrix[n-i-1][n-j-1] = matrix[j][n-i-1];
            matrix[j][n-i-1] = temp;
        }
    }
}
