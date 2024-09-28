package string

func isValid(s string) bool {

	stack := make([]byte, 0)
	// len奇数直接false
	if len(s)%2 == 1 {
		return false
	}
	//用栈来处理好一些
	//写一个map

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	// 遍历
	for i := 0; i < len(s); i++ {

		// 如果得到右括号，就出栈
		// map存在代表着有左括号，出栈
		// map存在意味着map值大于0（如果没有值byte为空）
		if pairs[s[i]] > 0 {
			//不能出栈或右括号与栈中左括号不对应
			if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
				return false
			}

			//出栈
			stack = stack[:len(stack)-1]

		} else { //如果得到了一个左括号，就入栈
			stack = append(stack, s[i])
		}

	}
	// 栈中无数即可
	return len(stack) == 0
}

/*func isValid(s string) bool {
ss:=strings.Split(s, "")
	//map[]
//字符串长度不对return
if len(ss)%2==1{
return false
}
zuoyou:=make(map[string]int)
zuoyou["("]=0
zuoyou["["]=0
zuoyou["{"]=0
zuoyou["}"]=1
zuoyou[")"]=1
zuoyou["]"]=1
lei:=make(map[string]int)
lei["("]=0
lei[")"]=0
lei["{"]=1
lei["}"]=1
lei["]"]=2
lei["["]=2
//遍历
for i:=0;i<len(ss);i++{
if i%2==0&&zuoyou[ss[i]]==1{
return false
}
if i%2==1&&zuoyou[ss[i]]==0{
	return false
	}
if i%2==0&&lei[ss[i]]!=lei[ss[i+1]]{
	return false
}
}
return true
}*/
/*func isValid(s string) bool {
	//分裂字符串
	ss := strings.Split(s, "")
	//判断长度
	length := len(ss)
	if length%2 == 1 {
		return false
	}

	//设定map
	lei := make(map[string]string)
	lei["("] = ")"

	lei["{"] = "}"

	lei["["] = "]"

	//双指针法
	//找到中点//123456
	left := len(ss) / 2
	right := len(ss)/2 + 1
	//向两边分
	for left >= 0 && right <= length-1 {
		if ss[right] != lei[ss[left]] {
			return false
		}
		left--
		right++
	}

	return true
}*/
