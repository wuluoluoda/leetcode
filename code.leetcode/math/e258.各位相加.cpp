class Solution
{
public:
    int addDigits(int num)
    {
        // 遍历
        /*sum := 0不该在这*/
        while (num >= 10)
        {
            int sum = 0;
            if (num == 10)
            {
                num = 1;
            }
            while (num > 0)
            {
                sum += num % 10;
                num /= 10;
            }
            num = sum;
        }
        return num;
    }
};