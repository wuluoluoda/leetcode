#include <unordered_map>
class Solution
{
public:
    bool containsNearbyDuplicate(vector<int>& nums, int k)
    {
        // 建map
        unordered_map<int, int> m = {};
        // 遍历
        for (int i = 0; i < nums.size(); i++)
        {
            // 存在，返回

            // 不存在，添加
            if (m.count(nums[i]) && i - m[nums[i]] <= k)
            {
                return true;
            }
            m[nums[i]] = i;
        }
        return false;
        // 存在，返回

        // 不存在，添加
    }
};