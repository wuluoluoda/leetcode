package hash

import "strings"
func canConstruct(ransomNote string, magazine string) bool {
	m:=make(map[string]int)
	ransomNotes:=strings.Split(ransomNote,"")
	magazines:=strings.Split(magazine,"")
	for _,v1:= range magazines{
	m[v1]++
	}
	for _,v2:= range ransomNotes{
	if m[v2]!=0{
	return false
	}else{
		m[v2]--
	}

	}
	return true


}
