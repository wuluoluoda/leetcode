func addDigits(num int) int {
	//遍历
	/*sum := 0不该在这*/
	for num >= 10 {
		sum := 0
		if num == 10 {
			num = 1
		}
		for num > 0 {
			sum += num % 10
			num /= 10
		}
		num = sum
	}
	return num
}