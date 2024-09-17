#include <string>
class Solution
{
public:
    bool isPalindrome(string s)
    {
        int i = 0;
        string sgood = "";
        while (i < s.size())
        {
            if ((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9'))
            {
                // 拼接损失性能
                // s = s[:i] + s[:i+2]
                // 不拼接，生成新数组，性能高
                // 每个符合条件的都加入一个新数组 islower

                /*sgood += s[i].lower();*/
                sgood += tolower(s[i]);
            }

            i += 1;
        }
        int n = sgood.size();

        // islower转化为字符串

        // 删掉标点，空格，大小写，然后判断是否回文

        // 左右指针
        int left = 0;
        int right = n - 1;
        while (left < right)
        {
            if (sgood[left] != sgood[right])
            {
                return false;
            }

            left += 1;
            right -= 1;
        }
        return true;
    }
};