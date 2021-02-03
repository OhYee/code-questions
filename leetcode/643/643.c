double findMaxAverage(int* nums, int numsSize, int k){
    int sum = 0;
    double res = -999999;
    for (int i=0; i<numsSize; ++i) {
        sum += nums[i];
        if (i >= k-1) {
            double avg = (double)(sum) / k;
            res = res > avg ? res : avg;
            sum -= nums[i-k+1];
        }
    }
    return res;
}
