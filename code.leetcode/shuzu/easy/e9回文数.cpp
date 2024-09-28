class Solution
{
public:
    bool isPalindrome(int x)
    {
        int num = 0;
        int biw = x;
        // 读出正负数
        if (x % 10 == 0 && x != 0)
        {
            return false;
        }
        if (x >= 0)
        { // 123%10=3 123/10=12 12%10=2 12/10=1 1%10=1 1/10=0
            while (biw > num)
            {

                num = num * 10 + biw % 10;
                biw = biw / 10;
            }
        }
        if (num == biw || biw == num / 10)//处理特殊是怕超出整数区间
        {
            return true;
        }
        return false;
    }
};