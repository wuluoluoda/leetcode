package fanben

import (
	"fmt"
	"sort"
)
func main() {
 //nums := []int{-4, -1, 0, 3, 10}
 nums := []int{-7,-3,2,3,11}
 sortedSquares(nums)
}
func sortedSquares(nums []int) []int {
 for index,value:=range nums{
  nums[index]=value*value
 }
 sort.Ints(nums)
 fmt.Println(nums)
 return nums
 
}