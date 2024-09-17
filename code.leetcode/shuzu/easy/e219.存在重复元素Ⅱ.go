func containsNearbyDuplicate(nums []int, k int) bool {
	//建map
	m := make(map[int]int)
	//遍历
	for i, v := range nums {
		//存在，返回

		//不存在，添加
		if value, ok := m[v]; ok {

			if i-value <= k {
				return true
			}
		}

		m[v] = i
	}
	return false
	//存在，返回

	//不存在，添加
}