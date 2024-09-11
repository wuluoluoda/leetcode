package string

import (
	"strings"
)

/*
罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
*/
func romanToInt(s string) int {
	direction := 1
	sum := 0
	//以map设定规则
	m := make(map[string]int, 7)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000
	//分裂字符串
	ss := strings.Split(s, "")
	//遍历
	for i := 0; i < len(ss); i++ {

		direction = 1
		//判断是否左右相反
		if i+1 < len(ss) && m[ss[i]] < m[ss[i+1]] {
			direction = -1
		}
		//相反则减去
		//否则加上
		sum = direction*(m[ss[i]]) + sum
	}
	//饭回数值
	return sum
}
