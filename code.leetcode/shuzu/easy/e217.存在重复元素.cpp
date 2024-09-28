#include <unordered_map>
class Solution
{
public:
    bool containsDuplicate(vector<int> &nums)
    {
        // 建map
        unordered_map<int, bool> m = {};
        // 遍历
        for (auto v : nums)
        {
            // 存在，返回

            // 不存在，添加
            if (m[v]==true)
                {
                    return true;
                }
            m[v] = true;
        }
        return false;
        // 存在，返回

        // 不存在，添加
    }
};