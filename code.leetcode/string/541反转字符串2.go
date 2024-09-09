package string

/*
给定一个字符串 s 和一个整数 k，从字符串开头算起，每计数至 2k 个字符，就反转这 2k 字符中的前 k 个字符。

如果剩余字符少于 k 个，则将剩余字符全部反转。
如果剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符，其余字符保持原样。
*/

//我的方法


/*
func reverseStr(s string, k int) string {
	//將s轉化爲[]bytes
	sbyte := []byte(s)
	leftnow:=byte(0)
	//每2k个字符反转前k个字符
	left, right := 0, len(sbyte)-1
	if right+1<k{
		left=0
		
		for  {
			leftnow=sbyte[left]
			sbyte[left]=sbyte[right]
			sbyte[right]=leftnow
			left++
			right--
			if left>=right{
				return string(sbyte)
			}	
		} 
	}
	
	if 2*k > right+1&&k<=right+1 {
		left=0
		right=k-1
		for  {
			leftnow=sbyte[left]
			sbyte[left]=sbyte[right]
			sbyte[right]=leftnow
			left++
			right--
			if left>=right{
				return string(sbyte)
			}	
		} 

	} else {
		for i:=1;i<=len(sbyte)/(2*k);i++ {
		    
		left=2*(i-1)*k
		right=2*i*k-k-1
		for  {
			leftnow=sbyte[left]
			sbyte[left]=sbyte[right]
			sbyte[right]=leftnow
			left++
			right--
			if left>=right{
				break
			}	
		}
		//
		if i+1>len(sbyte)/(2*k)&&i!=len(sbyte)/(2*k) {
			left=2*i*k
			sub:=len(sbyte)-1-left
			if sub<k{
				right=len(sbyte)-1
				
				for  {
					leftnow=sbyte[left]
					sbyte[left]=sbyte[right]
					sbyte[right]=leftnow
					left++
					right--
					if left>=right{
						return string(sbyte)
					}	
				} 
			}
			
			if 2*k >sub&&k<=sub {
				left=2*i*k
				right=2*i*k+k-1
				for  {
					leftnow=sbyte[left]
					sbyte[left]=sbyte[right]
					sbyte[right]=leftnow
					left++
					right--
					if left>=right{
						return string(sbyte)
					}	
				} 	    
		/*left=2*i*k
		right=len(sbyte)-1
			for  {
				sbyte[left], sbyte[right] = sbyte[right], sbyte[left]
                left++
                right--
				if left>=right{
					break
				}	
			}*//*
		}
	}
	//剩余字符少于k个，则将剩余字符全部反转
	
	}
	
}
return string(sbyte)
}*/

//官方方法
func reverseStr(s string, k int) string {
    ss := []byte(s)
    length := len(s)
    for i := 0; i < length; i += 2 * k {
     // 1. 每隔 2k 个字符的前 k 个字符进行反转
     // 2. 剩余字符小于 2k 但大于或等于 k 个，则反转前 k 个字符
        if i + k <= length {
            reverse(ss[i:i+k])
        } else {
            reverse(ss[i:length])
        }
    }
    return string(ss)
}

func reverse(b []byte) {
    left := 0
    right := len(b) - 1
    for left < right {
        b[left], b[right] = b[right], b[left]
        left++
        right--
    }
}