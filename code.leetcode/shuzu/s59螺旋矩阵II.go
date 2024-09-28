package shuzu

func generateMatrix(n int) [][]int {
	//循环传值
	//若k为ou数，对列操作
	//若k为ji数，对行操作
	var res = make([][]int, n)
	//除了做整体初始化，还要做行的初始化
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	var startx int = 0 //行起始位置
	var starty int = 0 //列起始位置
	var count int = 1  //给空格处赋值
	var offset int = 1 // 需要控制每一条边遍历的长度，每次循环右边界收缩一位

	var center int = n / 2
	//多少圈
	var loop int = n / 2
	//主动判断循环结束条件，不知道具体用处**********.反正用时少
	for loop > 0 {
		i, j := startx, starty

		for j = startx; j < n-offset; j++ {
			res[i][j] = count
			count++
		}
		for i = starty; i < n-offset; i++ {
			res[i][j] = count
			count++
		}
		//这里分号很重要********
		for ; j > offset; j-- {
			res[i][j] = count
			count++
		}
		for ; i > offset; i-- {
			res[i][j] = count
			count++
		}
		loop--
		startx++
		starty++
		offset++

	}

	// 判断n是否为奇数
	if n%2 == 1 {
		//res[n/2][n/2] = count+1,不能这么做，索引不能是小数
		res[center][center] = count + 1
	}
	return res
}
