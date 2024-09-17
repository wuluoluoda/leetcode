package shuzu

func merge(nums1 []int, m int, nums2 []int, n int) {
	//遍历nums2放入nums1的n部分
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	//排序
	sort.Ints(nums1)
}

/*直接合并
func merge(nums1 []int, m int, nums2 []int, _ int) {
    copy(nums1[m:], nums2)
    sort.Ints(nums1)
}

*/
