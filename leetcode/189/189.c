void rotate(int* nums, int numsSize, int k){
    k %= numsSize;
    for (int i=0; i<numsSize-1; ++i) {
        int p = (i - k + numsSize) % numsSize;  
        while (p < i) p = (p - k + numsSize) % numsSize;
        int temp = nums[i]; nums[i] = nums[p]; nums[p] = temp;
    }
}
