class Solution(object):
    def isPalindrome(self, s):
        i = 0
        sgood = ""
        while  i < len(s):
            if  (s[i] >= 'a' and s[i] <= 'z') or (s[i] >= 'A' and s[i] <= 'Z') or (s[i] >= '0' and s[i] <= '9') :
                #拼接损失性能
                #s = s[:i] + s[:i+2]
                #不拼接，生成新数组，性能高
                #每个符合条件的都加入一个新数组 islower
                
                sgood += s[i].lower()
                
                
                
            i+=1
        n = len(sgood)
		
	#islower转化为字符串
        
	#删掉标点，空格，大小写，然后判断是否回文
        
        
        #左右指针
        left = 0
        right = n - 1
        while left < right :
            if sgood[left] is not sgood[right] :
                return False
            
            left+=1
            right-=1
        

        return True
