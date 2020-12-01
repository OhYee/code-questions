/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
int *searchRange(int *nums, int numsSize, int target, int *returnSize) {
    int *res = malloc(sizeof(int) * 2);
    int  begin = -1, end = -1;
    *returnSize = 2;

    if (numsSize > 0) {
        begin = binary_search(nums, numsSize, target);
        if (begin < numsSize && nums[begin] == target)
            end = binary_search(nums, numsSize, target + 1) - 1;
        else
            begin = -1;
    }
    res[0] = begin;
    res[1] = end;
    return res;
}

int binary_search(int *nums, int size, int target) {
    int l = 0, r = size;
    int mid = (l + r) / 2;
    while (l < r) {
        if (target <= nums[mid]) {
            // 左侧
            r = mid;
            mid = (l + r) / 2;
        } else {
            // 右侧
            l = mid + 1;
            mid = (l + r) / 2;
        }
    }
    return l;
}