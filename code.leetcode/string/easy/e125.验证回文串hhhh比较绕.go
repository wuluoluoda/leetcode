package string

import "strings"

func isPalindrome(s string) bool {

	

	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') && !(s[i] >= 'A' && s[i] <= 'Z')&&!(s[i] >= '0' && s[i] <= '9') {
			//拼接损失性能
			s = s[:i] + s[i+1:]
			i--
		}
	}
	//删掉标点，空格，大小写，然后判断是否回文
	s = strings.ToLower(s)
	
	//左右指针
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
