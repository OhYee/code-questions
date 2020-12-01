#include <bits/stdc++.h>
using namespace std;

class Solution {
  public:
    vector<int> searchRange(vector<int> &nums, int target) {
        vector<int> res;
        int         begin = -1, end = -1;
        int         size = nums.size();
        if (size > 0) {
            auto it = lower_bound(nums.begin(), nums.end(), target);
            begin = it - nums.begin();
            if (it != nums.end() && *it == target) {
                it = lower_bound(nums.begin(), nums.end(), target + 1) - 1;
                end = it - nums.begin();
            } else
                begin = -1;
        }
        res.push_back(begin);
        res.push_back(end);
        return res;
    }
};