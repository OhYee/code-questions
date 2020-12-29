int minPatches(int* nums, int numsSize, int n){
    unsigned int m = 1, res = 0, i = 0;
    while (m <= n) m += (i < numsSize && nums[i] <= m) ? nums[i++] : (res++, m);
    return res;
}
