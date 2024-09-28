package math

func mySqrt(x int) int {
	//小于1直接返回0
	if x < 1 {
		return 0
	}
	//deng'yu
	//大于1，二分查找
	left := 0

	right := x
	for left <= right {
		mid := left + (right-left)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			if (mid+1)*(mid+1) > x {
				return mid
			}
			left = mid + 1
		}
		if mid*mid > x {
			if (mid-1)*(mid-1) < x {
				return mid - 1
			}
			right = mid - 1
		}
	}
	return 0
}
