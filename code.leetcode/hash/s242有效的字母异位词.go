package hash

import (
	"strings"

	"golang.org/x/text/unicode/rangetable"
)

/*给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

字母异位词 是通过重新排列不同单词或短语的字母而形成的单词或短语，通常只使用所有原始字母一次。*/
func isAnagram(s string, t string) bool {
//分裂两个字符串
s1:=strings.Split(s, "")
t1:=strings.Split(t, "")
vim:=multimap[string]bool{}
i:=0
for _,word:=range s1{
 vim[word]=true
 i++
}
for _,char:=range t1{
	if vim[char]{
		i--
	}
}
if i==0{
	return true
}
return false

}