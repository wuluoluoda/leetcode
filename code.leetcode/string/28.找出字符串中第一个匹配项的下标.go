package string
/*
给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串的第一个匹配项的下标（下标从 0 开始）。
如果 needle 不是 haystack 的一部分，则返回  -1 。
*/
func strStr(haystack string, needle string) int {
s:=[]byte(haystack)
t:=[]byte(needle)
m:=make(map[byte]int)

for i,v:=range s{
    
}
for j:=0;j<len(t);j++{
 m[t[j]]
}
}