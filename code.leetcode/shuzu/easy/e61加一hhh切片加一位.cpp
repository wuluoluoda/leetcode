class Solution {
public:
    vector<int> plusOne(vector<int>& digits) {
int n = digits.size();
	digits[n-1]++;
	//对数组处理，进位处理
	//如何进位，首先从后往前遍历，直到某位不为9
	for (int right = n - 1; right >= 0; --right) {
		if (digits[right] != 10) {
			break;
		} else if (right-1 >= 0) {
			digits[right] = 0;
			digits[right-1]++;
		} else {
			digits[right] = 0;
			//首位加一位
			digits.insert(digits.begin(), 1);
			
		}
	}
	return digits;
    }
};