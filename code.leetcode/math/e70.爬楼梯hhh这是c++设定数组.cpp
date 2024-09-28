#include <vector>
class Solution
{
public:
	int climbStairs(int n)
	{
		std::vector<int> f(n + 1);
		// f:=0/楼梯须爬n阶，一次1或2
		// f(n) = f(n-1) + f(n-2)
		if (n == 1)
		{
			return 1;
		}
		if (n == 2)
		{
			return 2;
		}

		if (n > 2)
		{
			f[1] = 1;
			f[2] = 2;
			// 遍历
			for (int i = 3; i <= n; ++i)
			{

				f[i] = f[i - 1] + f[i - 2];
			}
		}
		return f[n];
	}
};