package hash

import "sort"

/*import (
	"sort"
	"strconv"
)

func threeSum(nums []int) [][]int {
	//初始化数组
	s := []int{1, 2, 3}
	si := []int{1, 2, 3}
	var thenums [][]int

	m := make(map[int]int)
	m2 := make(map[string]bool)
	var str1 string
	var str2 string
	var str3 string

	/*存入m的v1，
	  如果-v2-v3存在且i1！=i2

	  输出append（）
	for i, v1 := range nums {
		m[v1]++

		for j, v2 := range nums {
			for k, v3 := range nums {
				sum := -v2 - v3
				si= []int{i, j, k}
				sort.Ints(si)
				str1 = strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(k)

				if m[sum] > 0 && i != j && i != k && j != k &&m[str1]{
					m[str1]=true
					m[str2]=true
					m[str3]=true
					m[sum]--
					s = []int{v1, v2, v3}

					thenums = append(thenums, s)
				}
			}

		}
	}
	return thenums
}
*/
func threeSum(nums []int) [][]int {
sort.Ints(nums)
	res := [][]int{}
	// 找出a + b + c = 0
	// a = nums[i], b = nums[left], c = nums[right]
	for i := 0; i < len(nums)-2; i++ {
		// 排序之后如果第一个元素已经大于零，那么无论如何组合都不可能凑成三元组，直接返回结果就可以了
		n1 := nums[i]
		if n1 > 0 {
			break
		}
		// 去重a
		if i > 0 && n1 == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3 == 0 {
				res = append(res, []int{n1, n2, n3})
				// 去重逻辑应该放在找到一个三元组之后，对b 和 c去重
				for l < r && nums[l] == n2 {
					l++
				}
				for l < r && nums[r] == n3 {
					r--
				}
			} else if n1+n2+n3 < 0 {
				l++
			} else {
				r--
			}
		}
	}
	return res
	}
