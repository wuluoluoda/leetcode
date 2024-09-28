#include <vector>
class Solution
{
public:
    bool isAnagram(string s, string t)
    {
        if (s.length() != t.length()) {
            return false;
        }
        vector<int> c1(26, 0);
        vector<int> c2(26, 0);
        for (auto& ch : s)
        {
            c1[ch - 'a']++;
        }
        for (auto& ch : t)
        {
            c2[ch - 'a']++;
        }
        return c1 == c2;
    }
};