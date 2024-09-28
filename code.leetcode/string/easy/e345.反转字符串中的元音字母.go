func reverseVowels(s string) string {
	//创造一个map来存储元音字母
	ss := []rune(s)
	vowels := map[rune]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	//如果存在元音字母，则记录其位置
	var pos []int
	for i, c := range s {
		if vowels[c] {
			pos = append(pos, i)
		}
	}
	//对记录的位置进行反转
	for i := 0; i < len(pos)-1-i; i++ {
		ss[pos[i]], ss[pos[len(pos)-1-i]] = ss[pos[len(pos)-1-i]], ss[pos[i]]

	}
	return string(ss)
}