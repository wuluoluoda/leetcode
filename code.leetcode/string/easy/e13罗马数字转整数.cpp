#include <unordered_map>
#include <string>
class Solution
{
    // 添加外部条件
private:
    unordered_map<char, int> symbolValues = {
        {'I', 1},
        {'V', 5},
        {'X', 10},
        {'L', 50},
        {'C', 100},
        {'D', 500},
        {'M', 1000},
    };

public:
    int romanToInt(string s)
    {
        int direction = 1;
        int sum = 0;
        int n = s.size();
        int value = 0;

        // 遍历取出byte
        for (int i = 0; i < n; i++)
        {
            value = symbolValues[s[i]];
            direction = 1;
            // 判断是否左右相反
            if (i + 1 < n && value < symbolValues[s[i + 1]])
            {
                direction = -1;
            }

            // 相反则减去
            // 否则加上
            sum += direction * value;
        }

        // 饭回数值
        return sum;
    }
};