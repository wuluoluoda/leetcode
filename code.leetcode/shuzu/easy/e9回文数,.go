package shuzu

func isPalindrome(x int) bool {
	var num int
	biw := x
	//读出正负数

	if x >= 0 { //123%10=3 123/10=12 12%10=2 12/10=1 1%10=1 1/10=0
		for biw > 0 {

			num = num*10 + biw%10
			biw = biw / 10
		}
	}
	if num == x {
		return true
	}
	return false
}
