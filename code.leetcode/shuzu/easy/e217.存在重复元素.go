func containsDuplicate(nums []int) bool {
//建map
	m:=make(map[int]bool)
//遍历
for _,v:=range nums{
    //存在，返回

	//不存在，添加
	if m[v]{
        return true
    }
    m[v]=true
}	
return false
	//存在，返回

	//不存在，添加
}