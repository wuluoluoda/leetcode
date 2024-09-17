func reverseBits(num uint32) uint32 {
	//转化为str
	s := string(num)

	//转化为数组
	arr := []byte(s)

	//反转数组
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	//转化为str
	s = string(arr)
	//转化为int
	i, _ := strconv.ParseInt(s, 2, 64)
	//转化为uint32
	return uint32(i)

}