int minCostClimbingStairs(int* cost, int costSize){
    int dp[1005];
    dp[0] = cost[0];
    if (costSize >= 1) dp[1] = cost[1];
    for (int i=2; i<=costSize; ++i) {
        dp[i] = fmin(dp[i-2], dp[i-1]);
        if (i < costSize) dp[i] += cost[i];
    }
    return dp[costSize];
}
