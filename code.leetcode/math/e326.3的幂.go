/*简化
func isPowerOfThree(n int) bool {
    for n > 0 && n%3 == 0 {
        n /= 3
    }
    return n == 1
}

作者：力扣官方题解
链接：https://leetcode.cn/problems/power-of-three/solutions/1011809/3de-mi-by-leetcode-solution-hnap/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func isPowerOfThree(n int) bool {
if n <= 0 {
		return false
}
if n == 1 {
		return true
}
for n%3 == 0 {
		n /= 3
}
if n == 1 {
		return true
}
return false
}