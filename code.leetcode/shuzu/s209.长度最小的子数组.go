package shuzu

import (
	"fmt"
	"math"
)

func MinSubArrayLen(target int, nums []int) int {
	//算出最大数要n个

	//算出最小数要m个
	summax := nums[0]
	for i := 0; i < len(nums); i++ {

		if summax <= nums[i] {
			summax = nums[i]

		}

	}
	fmt.Println(summax)
	m1 := math.Ceil(float64(target) / float64(summax))
	if int(m1) > len(nums) {
		return 0
	}
	fmt.Println(m1)

	//算出[n]int总和

	//算出[m]int总和
	var sums int
	minm := len(nums) + 1
	for m := int(m1); m <= len(nums); m++ {

		for index, _ := range nums[:len(nums)-m+1] {
			sum := nums[index : index+m]
			sums = 0
			for _, v := range sum {
				sums = sums + v

				if sums >= target {
					minm = m
					fmt.Println(minm)
					return minm //这里改成retrn minm
				}

			}

		}

	}
	if minm == len(nums)+1 { //如果minm没有被更新，说明没有找到满足条件的子数组
		return 0
	}
	return minm
}
