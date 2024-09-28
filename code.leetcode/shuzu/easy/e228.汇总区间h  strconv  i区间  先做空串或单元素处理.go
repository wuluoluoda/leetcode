func summaryRanges(nums []int) []string {
	//先做空串处理
	if len(nums) == 0 {
		return []string{}
	}
	//做单元素处理
	if len(nums) == 1 {
		return []string{strconv.Itoa(nums[0])}
	}
	var res []string
	var start, end int
	//遍历
	for i := 0; i < len(nums)-1; i++ {
		//开头生成start和end
		if i == 0 {
			start = nums[i]
			end = nums[i]
		}
		//如果相邻，变更标签
		if nums[i]+1 == nums[i+1] {
			end = nums[i+1]
		}

		//如果不相邻，根据标签输出string并放入切片
		if nums[i]+1 != nums[i+1] {
			//输出首尾res
			res = append(res, listornum(start, end))
			//确定下一个首尾
			start = nums[i+1]
			end = nums[i+1]
			
		}
	}
	//输出最后一个
	res = append(res, listornum(start, end))
	return res
}
func listornum(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	}
	return strconv.Itoa(start) + "->" + strconv.Itoa(end)
}