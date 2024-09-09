package shuangzhizhen
/*// 暴力法
// 时间复杂度 O(n^2)
// 空间复杂度 O(1)
func removeElement(nums []int, val int) int {
    size := len(nums)
    for i := 0; i < size; i ++ {
        if nums[i] == val {
            for j := i + 1; j < size; j ++ {
                nums[j - 1] = nums[j]
            }
            i --
            size --
        }
    }
    return size
}
*/

//双指针法
func removeElement(nums []int, val int) int {
	// 初始化慢指针 slow
	slow := 0
	//快指针遇到val跳过，跳过时慢指针不动
	//慢指针因此就少走了跳过的步数，而得出了刚好剩下的元素
	for fast:=0;fast<len(nums);fast++{
	
		if nums[fast]==val{
	continue    
	}
	//这一步用来保证fast遇到val后续数值的更新，因为fast遇到val时，fast++，fast指向下一个元素，而slow指针不动	
nums[slow]=nums[fast]
	slow++

	}
	return slow
}
//双指针法
func removeElement2(nums []int, val int) int {
	left := 0
			for _, v := range nums { // v 即 nums[right]
				if v != val {
					nums[left] = v
					
					left++
				}
			}
			return left
	}
	
	