class Solution {
public:
    void moveZeroes(vector<int>& nums) {
//先删掉再补全
	//双指针
int left=0;
for (int i=0;i<nums.size();++i){
    if (nums[i]!=0){
        nums[left]=nums[i];
        ++left;
    }
}
//补全
for (int j=left;j<nums.size();++j){
    nums[j]=0;
}
    }
};