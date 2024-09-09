package main

import (
	"fmt"

	"code.leetcode/shuzu"
	"code.leetcode/util"
)

var s string = util.Name

func main() {

	res := shuzu.MinSubArrayLen(15, []int{1, 2, 3, 4, 5})
	fmt.Println(s, res) //[0.4]

}
