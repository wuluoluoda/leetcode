package hash
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
//数组分别存入hash
count:=0
vis:=make(map[int]int)
//h构建nums1.2和的map
for _,v1:= range nums1{
	for _,v2:= range nums2{
		vis[v1+v2]++
	}
}
//检查v3v4的和里有没有前两个的相反数
for _,v3:=range nums3{
for _,v4:=range nums4{
	   sum:=-v3-v4
	   if countval,ok:=vis[sum];ok{
	   count+=countval
}
}
}

return count
}
