func isAnagram(s string, t string) bool {
	//分裂两个字符串
	if len(s) != len(t) {
		return false
	}
	s1 := strings.Split(s, "")
	t1 := strings.Split(t, "")
	vim := map[string]int{}

	for _, word := range s1 {
		vim[word]++

	}
	for _, char := range t1 {
		if num, ok := vim[char]; ok {
			/*num--*/
			num--
			vim[char] = num
		} else {
			return false
		}
	}
	//检查map是否都为空
	for _, num := range vim {
		if num != 0 {
			return false
		}
	}
	return true

}