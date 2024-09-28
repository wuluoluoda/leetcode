func wordPattern(pattern string, s string) bool {
	
	words := strings.Split(s, " ")
	//长度不等直接返回false
	if len(pattern) != len(words) {
		return false
	}
	patternMap := make(map[byte]string)
	wordsMap := make(map[string]byte)
	for i := range pattern {
		//存在值但不对应就返回false
		if (patternMap[pattern[i]] != "" && patternMap[pattern[i]] != words[i]) || (wordsMap[words[i]] != 0 && wordsMap[words[i]] != pattern[i]) {
			return false

		}
		patternMap[pattern[i]] = words[i]
		wordsMap[words[i]] = pattern[i]
	}
	return true
}

/*
func isIsomorphic(s, t string) bool {
    s2t := map[byte]byte{}
    t2s := map[byte]byte{}
    for i := range s {
        x, y := s[i], t[i]
        if s2t[x] > 0 && s2t[x] != y || t2s[y] > 0 && t2s[y] != x {
            return false
        }
        s2t[x] = y
        t2s[y] = x
    }
    return true
}
*/