
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
     unordered_map<int, int> m={};
    for (int index=0;index<nums.size();++index) {
        int val=nums[index];
        if (m.find(target-val) != m.end())  {
            return { m[target-val], index};
        } else {
            m[val] = index;
        }
    }
    return {};   
    }
};