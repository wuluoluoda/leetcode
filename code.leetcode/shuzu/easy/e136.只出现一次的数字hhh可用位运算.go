func singleNumber(nums []int) int {
	/*//建map
	m := make(map[int]bool)
	//if 存在，反转bool
	//不存在，存入，下一个
	for _, v := range nums {
		if m[v] {
			m[v] = false
		} else {
			m[v] = true
		}
	}
	//遍历map，返回true的值
	for k, v := range m {
		if v {
			return k
		}
	}
	return 0*/


	
		single := 0
		for _, num := range nums {
			single ^= num
		}
		return single
	
	
	

}