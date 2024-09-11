package string

func strStr(haystack string, needle string) int {
	//needle更长，直接返回
	if len(needle) > len(haystack) {
		return -1
	}
	//ha

	//计算needle的长度1234  345
	//把needle反向放入栈中
	//遍历haystack，如果haystack[i] == needle[0]，则从i开始，比较haystack[i:i+len(needle)]和needle是否相等

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] {
			if haystack[i:i+len(needle)] == needle {
				return i
			}
		}
	}
	return -1
}
