package util
const Name = "kss"
func MinSubArrayLen(target int, nums []int) int {
	//算出最大数要n个
	
	//算出最小数要m个
	summax:= nums[0]
 for i := 0; i < len(nums)-1; i++ {
	 
	 if summax<=nums[i]{
 summax=nums[i]
		 
	 }
	 
 }
 return 0
}
 //m1 := math.Ceil(float64(target) / float64(summax))
// 下面的全部 
 
	//算出[n]int总和
   
	//算出[m]int总和
	//minm := len(nums) + 1
	//for m:=int(m1);m< len(nums);m++{
	//for index,_ := range nums[:len(nums)-m]{
	// sum:= nums[index:index+m]
	// sums:=0
	 //for _,v := range sum{
		//sums=sums+v 
	 
	// if sums>=target{
	//	minm=m
	//	 break
	// }
	 
	//	}
	 
		 
	 //}
	 //if minm==len(nums) + 1{//如果minm没有被更新，说明没有找到满足条件的子数组
	//	 return 0
		 
	 //}
 //}
 //return minm	
 //}  