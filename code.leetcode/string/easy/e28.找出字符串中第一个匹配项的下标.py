class Solution(object):
    def strStr(self, haystack, needle):
        #needle更长，直接返回
        if len(needle)>len(haystack):
            return -1

            #ha

            #计算needle的长度1234  345
        #把needle反向放入栈中
        #遍历haystack，如果haystack[i] == needle[0]，则从i开始，比较haystack[i:i+len(needle)]和needle是否相等
            i=0
        
        for i in range(len(haystack)-len(needle)+1):
            if haystack[i]== needle[0]: 
                if haystack[i:i+len(needle)]==needle:
                    return i
            i+=1
            

        return -1
