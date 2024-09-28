package shuzu

/*import (
	"container/list")
func main() {
	digits := list.New()
	digits.PushFront(1) // 在开头插入元素

	// 将 list 转换回切片
	var slice []int
	for e := digits.Front(); e != nil; e = e.Next() {
		slice = append(slice, e.Value.(int))
	}

	fmt.Println(slice) // 输出: [1]
}
//使用 container/list 包可以显著提高在列表开头插入元素时的性能，特别是在处理大量数据时。*/
func plusOne(digits []int) []int {
	n := len(digits)
	digits[n-1]++
	//对数组处理，进位处理
	//如何进位，首先从后往前遍历，直到某位不为9
	for right := n - 1; right >= 0; right-- {
		if digits[right] != 10 {
			break
		} else if right-1 >= 0 {
			digits[right] = 0
			digits[right-1]++
		} else {
			digits[right] = 0
			//首位加一位
			digits = append([]int{1}, digits...)
		}
	}
	return digits
}
