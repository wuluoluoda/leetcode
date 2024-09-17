class Solution
{
public:
    bool isUgly(int n)
    {
        if (n == 1)
        {
            return true;
        }

        if (n <= 0)
        {
            return false;
        }
        int s2;
        int s3;
        int s5;
        while (n != 1)
        {
            int bit = 0;
            s2 = n % 2;
            if (s2 == 0)
            {
                n = n / 2;
                bit++;
            }
            s3 = n % 3;
            if (s3 == 0)
            {
                n = n / 3;
                bit++;
            }
            s5 = n % 5;
            if (s5 == 0)
            {
                n = n / 5;
                bit++;
            }
            if (bit == 0)
            {
                return false;
            }
        }
        return true;
    }
};