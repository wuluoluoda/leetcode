class Solution
{
public:
    int mySqrt(int x)
    {
        // 小于1直接返回0
        if (x < 1)
        {
            return 0;
        }
        // deng'yu
        // 大于1，二分查找
        int left = 0;

        int right = x;
        while (left <= right)
        {
            int mid = left + (right - left) / 2;
            if ((long long)mid * mid == x)
            {
                return mid;
            }
            if ((long long)mid * mid < x)
            {
                if ((long long)(mid + 1) * (mid + 1) > x)
                {
                    return mid;
                }
                left = mid + 1;
            }
            if ((long long)mid * mid > x)
            {
                if ((long long)(mid - 1) * (mid - 1) < x)
                {
                    return mid - 1;
                }
                right = mid - 1;
            }
        }
        return 0;
    }
};