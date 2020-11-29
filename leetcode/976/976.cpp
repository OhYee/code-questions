#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    int largestPerimeter(vector<int> &A) {
        sort(A.begin(), A.end(), greater<int>());
        for (int i = 2; i < A.size(); ++i)
            if (A[i - 2] < A[i - 1] + A[i])
                return A[i - 2] + A[i - 1] + A[i];
        return 0;
    }
};