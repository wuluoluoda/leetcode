package fanben

import (
	"sort"
	"strings"
)
func groupAnagrams(strs []string) [][]string {
   
    
	data := make([]string, len(strs))
    
    list := make([][]string, 0)
	list[0]=append(list[0], strs[1]) 
    havelisted := make([]int, 0)
	havelisted[0] = 1
    havelistedbool := false
    
    for index,str:= range strs{
       //取最小单元
        
        minstrs:=strings.Split(str, "")
        //最小单元排序
		
		sort.Strings(minstrs)
		//这里能重复声明（var）吗
		minstr := ""

		minstr = strings.Join(minstrs, "")
		
		data[index] = minstr
		
    }
     //最小单元对比，相同归类
        i, j, k := 0, len(strs)-1, 0
       for {
            for i<j{
            //检查是否已被分组，若未被分组，建立单个字符串的切片，先将自己归为一组，
            //检查i是否已在 havelisted中
             for _,value:= range havelisted{
                if i==value{
            havelistedbool=true
            //这一步最终要体现为回到flase状态***
                }
                
             }
			 //在这段代码中，你试图将一个切片赋值给一个变量，
			 //但是 Go 语言不允许这样做。在 Go 语言中，
			 //你不能直接使用 {} 来创建一个切片，
			 //你需要使用 make 函数或者 append 函数。
             if !havelistedbool{
             list[k]=append(list[k], strs[i]) 
             } 
			 //体现为回到flase状态***
             havelistedbool=false
            //下一波预先留字符串已分组标记

            if data[i]==data[j]{          
            //预留j已被分组标记
            havelisted=append(havelisted,j)
             //相同归类， //同类归一数组
            list[k]=append(list[k],strs[j])
           
           
            }
         //换下一个元素适配i   
        j--
     }
     //list数组整合
     list=append(list,list[k])
     i++
     k++
     j=len(strs)-1
     
 }      

}