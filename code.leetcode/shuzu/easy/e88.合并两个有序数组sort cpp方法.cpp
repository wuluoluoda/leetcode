class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
//遍历nums2放入nums1的n部分
	for(int i = 0; i < n; ++i) {
		nums1[m+i] = nums2[i];
	}
	//排序************
	std::sort(nums1.begin(), nums1.end());

    }
};