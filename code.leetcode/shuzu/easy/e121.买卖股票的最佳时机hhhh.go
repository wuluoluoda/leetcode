package shuzu

func maxProfit(prices []int) int {
	inf := 10000
	minPrice := inf
	maxProfit := 0

	for _, price := range prices {
		maxProfit = max(maxProfit, price-minPrice)
		minPrice = min(minPrice, price)
	}
	return maxProfit
}

// max 返回两个整数中较大的一个
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min 返回两个整数中较小的一个
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*//先找最小
	min:=prices[0]
	i:=1
	for ;i<len(prices)-1;i++{
	    if prices[i]<min{
	        min=prices[i]
			break
	    }
	}
	max:=prices[len(prices)-1]
	//往右找，找最大，同时记录更小
	for j:=i+1;j<len(prices);j++{
	    if prices[j]>max{
	        max=prices[j]
	    }
		if prices[j]<min{
	        min=prices[i]
		}
	}
}

	//记录差值

	//换更小为起点
	//重复*/

/*/map法
//排序升序
sort.Ints(prices)
//遍历，找差值最大
for left, right :=0, len(prices)-1; left < right; {

}*/
