class Solution
{
public:
    int removeElement(vector<int> &nums, int val)
    {
     int left= 0;
    for (int v : nums)
        { // v 即 nums[right]
            if (v!= val)
                {
                    nums[left] = v;
                        ++left;
                }
        }
        return left;
    }
};