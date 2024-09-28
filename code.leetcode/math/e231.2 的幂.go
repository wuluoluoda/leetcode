func isPowerOfTwo(n int) bool {
//估计有位运算的方法
/*
func isPowerOfTwo(n int) bool {
    // 小于 0 的数肯定不是 2 的幂
    if n < 0 {
        return false
    }
    // 2 的幂减 1 后，会产生一个所有位都是 1 的数
    // 例如：2^3 = 8，8-1 = 7，二进制为 111
    // 所以，2 的幂的数和它减 1 的数进行按位与操作后，结果应该为 0
	8 的二进制表示是 1000。
7 的二进制表示是 0111。
进行位与操作：

第1位：1 & 0 = 0
第2位：0 & 1 = 0
第3位：0 & 1 = 0
第4位：0 & 1 = 0
    return n > 0 && (n & (n - 1)) == 0
}
*/

func isPowerOfTwo(n int) bool {
    const big = 1 << 30
    return n > 0 && big%n == 0
}


}