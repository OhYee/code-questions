#define MAXN (50005)
int dp[MAXN][2];

int maxProfit(int* prices, int pricesSize, int fee){
    int dp[2][2];
    dp[0][0]=0; dp[0][1] = -prices[0];
    for (int i=1; i<pricesSize; ++i) {
        int t=i&1,p=!t;
        dp[t][0] = fmax(dp[p][0], dp[p][1]+prices[i]-fee);
        dp[t][1] = fmax(dp[p][1], dp[p][0]-prices[i]);
        // printf("%d: %d %d\n", i, dp[t][0], dp[t][1]);
    }
    int lst = (pricesSize-1)&1;
    return fmax(0, fmax(dp[lst][0], dp[lst][1]));
}
