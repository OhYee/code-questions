/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int* threeEqualParts(int* A, int ASize, int* returnSize){
    int *res = malloc(sizeof(int)*2);
    res[0] = -1; res[1] = -1; *returnSize = 2;

    // 数组无法分成非空的 3 份
    if (ASize < 3) return res;

    int count = 0;
    for (int i=0; i<ASize; ++i)
        if (A[i] == 1)
            ++count;

    // 1 的个数无法整除
    if (count % 3 != 0) return res;
    // 全是 0
    if (count == 0) {res[0] = 0; res[1] = 2; return res;}

    // 获取关键的 1 的位置
    int p[6] = {0, 0, 0, 0, 0, 0};
    int n1 = count / 3; // 每组中 1 的长度
    int temp = 0;
    for (int i=0; i<ASize; ++i) {
        if (A[i] == 1) {
            ++temp;
            if (temp == 1) p[0] = i;            // 第一组第一个 1
            if (temp == n1) p[1] = i;           // 第一组的最后一个 1
            if (temp == n1 + 1) p[2] = i;       // 第二组的第一个 1
            if (temp == n1 * 2) p[3] = i;       // 第二组的最后一个 1
            if (temp == n1 * 2 + 1) p[4] = i;   // 第三组的第一个 1
            p[5] = i;                           // 第三组的最后一个 1
        }
    }

    // 后导零个数
    int zero = ASize - p[5] - 1;
    
    // 确定 i 和 j 的位置
    int i = p[1] + zero;
    int j = p[3] + zero + 1;
    if (i >= p[2] || j > p[4] || i+1 >= j)
        return res;
    
    // 判断该位置分割的区间是否合法
    if (p[1] - p[0] != p[3] - p[2] || p[1] - p[0] != p[5] - p[4])
        return res;
    for (int i=0; i< p[1]-p[0]; ++i)
        if (A[p[0] + i] != A[p[2] + i] || A[p[0] + i] != A[p[4] + i]) 
            return res;

    res[0] = i; res[1] = j;
    return res;
}
