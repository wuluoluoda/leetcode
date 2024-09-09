package fanben

func RemoveElement(nums []int, val int) int {
i:=0
lennums:=len(nums)   
for index,value := range nums{
	//若存在
    if value == val{
		
        nums= append(nums[:index],nums[index+1:]...)
		i++
		continue
    }
	}
	return lennums-i
}

