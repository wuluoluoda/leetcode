//我的方法，复杂了

/*func isIsomorphicdemo(s string, t string) bool {
	//记录mapint
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		//value, ok := m[s[i]]
		_, ok := m[s[i]]
		if ok {
			//检查这里的t是否一致，检查索引越界

			if t[m[s[i]]] == t[i] {
				continue
			} else {
				return false
			}

		} else {
			m[s[i]] = i
		}
	}

	//检查ok
	//ok则记录
	//检查记录()点~
	//
	return true
}
func isIsomorphic(s string, t string) bool {
	if isIsomorphicdemo(s, t) && isIsomorphicdemo(t, s) {
		return true

	}
	return false

}*/



//别人的方法，简洁
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

