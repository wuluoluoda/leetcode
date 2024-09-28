/*func summaryRanges(nums []int) (ans []string) :
    for i, n := 0, len(nums); i < n; :
        left := i
        for i++; i < n && nums[i-1]+1 == nums[i]; i++ :

        s := strconv.Itoa(nums[left])
        if left < i-1 :
            s += "->" + strconv.Itoa(nums[i-1])

        ans = append(ans, s)

    return*/

#include <vector>
class Solution
{
public:
    vector<string> summaryRanges(vector<int> &nums)
    {
        vector<string> ans;
        string s;
        for (int i = 0, n = nums.size(); i < n;)
        {
            int left = i;
            ++i;
            while (i < n && nums[i - 1] + 1 == nums[i])
            {
                ++i;
            }
            s= to_string(nums[left]);
            if (left < i - 1)
            {
                s += "->" + to_string(nums[i - 1]);
            }
            ans.push_back(s);
        }

        return ans;
    }
};