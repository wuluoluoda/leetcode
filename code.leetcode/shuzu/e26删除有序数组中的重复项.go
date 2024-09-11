package shuzu
func removeDuplicates(nums []int) int {
	//设map
	m:=make(map[int]bool)
	//遍历数组
	left := 0
	for _, v := range nums { // v 即 nums[right]
		if !m[v] {
			nums[left] = v
			left++
			m[v]=true
		}
	}
	
	return left
	
	}