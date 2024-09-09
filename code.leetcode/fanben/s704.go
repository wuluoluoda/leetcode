package fanben

// 二分查找
func search(nums []int, target int) int {
	fash:=-1
for index,value := range nums{
	//若存在
    if value == target{
		fash=index
        return fash
    }else{
        continue
	}

}
return fash
}
//可补充二分法写法