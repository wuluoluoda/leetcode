#include <vector>
class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
    //设map
	//******* */
	unordered_map<int, bool> m={};
	//遍历数组
	int left = 0;
	for (int v: nums ){ // v 即 nums[right]
		//********* */
		if (m.find(v) == m.end()) {
			nums[left] = v;
			++left;
			m[v]=true;
		}
	}
	
	return left ;   
    }
};