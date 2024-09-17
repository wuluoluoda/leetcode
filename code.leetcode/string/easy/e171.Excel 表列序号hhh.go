func titleToNumber(columnTitle string) int {
	var  res int
	loop:=0
	//rune形式遍历string
	s := []rune(columnTitle)
	for _, v := range s {
		//A-Z对应65-90
		//进行进制转换
		v = v - 64
		//v强制以原值转化为int
		s := int(v)

		//相加

		/*res = res + v*/
		//错误在v是rune，res是int，已经改了var rune res = 0
		res = res*26 + s
		loop++
	}
return res
}