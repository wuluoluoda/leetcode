package shuzu

/*给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。*/
func moveZeroes(nums []int) {
	//先删掉再补全
	//双指针
	left := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[left] = nums[i]
			left++
		}
	}
	//补全
	for j := left; j < len(nums); j++ {
		nums[j] = 0
	}
}
