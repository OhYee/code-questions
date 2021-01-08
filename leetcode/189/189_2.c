void swap(int *a, int *b) {
    int temp = *a;
    *a = *b;
    *b = temp;
}

void reverse(int *nums, int numsSize) {
    for (int i=0; i<numsSize/2; ++i)
        swap(&nums[i], &nums[numsSize-i-1]);
}

void rotate(int* nums, int numsSize, int k){
    k %= numsSize;
    reverse(nums, numsSize);
    reverse(nums, k);
    reverse(nums + k, numsSize - k);
}
