#include <unordered_map>
class Solution {
public:
    int missingNumber(vector<int>& nums) {
unordered_map<int,bool>has={};
    for (auto v :nums) {
        has[v] = true;
    }
    for (int i = 0; ; i++) {
        if (!has[i]) {
            return i;
        }
    }
    }
};
/*
class Solution {
public:
    int missingNumber(vector<int>& nums) {
        unordered_set<int> set;
        int n = nums.size();
        for (int i = 0; i < n; i++) {
            set.insert(nums[i]);
        }
        int missing = -1;
        for (int i = 0; i <= n; i++) {
            if (!set.count(i)) {
                missing = i;
                break;
            }
        }
        return missing;
    }
};


 */