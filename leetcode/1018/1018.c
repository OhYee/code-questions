/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
bool* prefixesDivBy5(int* A, int ASize, int* returnSize){
    *returnSize = ASize;
    bool *res = malloc(sizeof(bool) * (*returnSize));
    int num = 0;
    for (int i=0; i<ASize; ++i) 
        res[i] = (num = (num * 2 + A[i]) % 5) == 0;
    return res;
}
