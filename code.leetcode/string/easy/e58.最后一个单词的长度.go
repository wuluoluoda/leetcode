package string

func lengthOfLastWord(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			for j := i; j >= 0; j-- {
				if s[j] == ' ' {
					return i - j
				} else if j == 0 {
					return i + 1
				}
			}
		}

	}
	return 0
}
