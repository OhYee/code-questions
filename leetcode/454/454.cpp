#include <bits/stdc++.h>
using namespace std;

class Solution {
    unordered_map<int, int> m;

  public:
    int fourSumCount(vector<int> &A, vector<int> &B, vector<int> &C,
                     vector<int> &D) {
        m.clear();
        int res = 0;
        for (auto a : A)
            for (auto b : B)
                ++m[a + b];
        for (auto c : C)
            for (auto d : D)
                res += m[-c - d];
        return res;
    }
};