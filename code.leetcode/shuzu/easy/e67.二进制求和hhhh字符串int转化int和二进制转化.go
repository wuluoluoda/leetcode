package shuzu

import "strconv"

func addBinary(a string, b string) string {
	suma, _ := strconv.ParseInt(a, 2, 64)
	sumb, _ := strconv.ParseInt(b, 2, 64)
	sum := suma + sumb
	return strconv.FormatInt(sum, 2)

}
