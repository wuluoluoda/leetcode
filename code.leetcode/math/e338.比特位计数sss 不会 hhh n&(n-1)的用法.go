func countBits(n int) []int {
//计算有几位，化成二进制数时0001 0010 0011 0100 0101 0110 0111 1000 20
/*when满位,+1,+1.
+2,-1+1
+3,+1
+4,-2+1*/
if n==0{
	return []int{0}
}
weishu:=0
//检验2进制几位
for i:=0;;i++{
if n%(2^i)&&n/(2^i)==0{
weishu=i
break
}
    
}
if n%2==0{
    n/2
}
//2+4+8
}