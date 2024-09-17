#include <string>
class Solution
{
public:
    int titleToNumber(string columnTitle)
    {
        int res=0;
        int loop = 0;
        // rune形式遍历string

        for (auto v : columnTitle)
        {
            // A-Z对应65-90
            // 进行进制转换
            v = v - 64;
            // v强制以原值转化为int
            int s = int(v);

            // 相加

            /*res = res + v*/
            // 错误在v是rune，res是int，已经改了var rune res = 0
            res = res * 26 + s;
            loop += 1;
        }
        return res;
    }
};