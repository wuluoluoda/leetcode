func isUgly(n int) bool {
	if n == 1 {
		return true
	}

	if n <= 0 {
		return false
	}

	for n != 1 {
		bit := 0
		s2 := n % 2
		if s2 == 0 {
			n = n / 2
			bit++
		}
		s3 := n % 3
		if s3 == 0 {
			n = n / 3
			bit++
		}
		s5 := n % 5
		if s5 == 0 {
			n = n / 5
			bit++
		}
		if bit == 0 {
			return false
		}
	}
	return true
}