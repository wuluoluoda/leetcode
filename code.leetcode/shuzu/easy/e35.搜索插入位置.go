package shuzu

func searchInsert(nums []int, target int) int {
	/*给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
	如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

	请必须使用时间复杂度为 O(log n) 的算法。*/
	mid := 0
	//二分查找01234567
	left, right := 0, len(nums)-1
	for left <= right {
		mid = left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		}
		if nums[mid] > target {
			right = mid - 1
		}
		left++
		right--
	}
	return mid + 1
}
