package hash

//暴力法
/*
func twoSum(nums []int, target int) []int {

	for i, v1:= range nums {
		for j, v2 := range nums {

			if target == v1+v2 && i != j {
				return []int{i, j}
			}

		}


	}
	return []int{}
}*/

//hash法
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		//如果m[target-val]存在，则返回
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}

}
