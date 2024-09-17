#include <vector>
class Solution
{
public:
    // max 返回两个整数中较大的一个
    int max(int a, int b)
    {
        if (a > b)
        {
            return a;
        }
        return b;
    }

    // min 返回两个整数中较小的一个
    int min(int a, int b)
    {
        if (a < b)
        {
            return a;
        }
        return b;
    }
    int maxProfit(vector<int> &prices)
    {
        int inf = 10000;
        int minPrice = inf;
        int maxProfit = 0;

        for (int price : prices)
        {
            maxProfit = max(maxProfit, price - minPrice);
            minPrice = min(minPrice, price);
        }
        return maxProfit;
    }
};