#s = [ord(c) for c in columnTitle]
#这段代码将遍历字符串 columnTitle 中的每个字符，
# 并将每个字符的 Unicode 码点（整数）添加到列表 s 中。
#如果你需要将 Unicode 码点转换回字符，你可以使用 chr() 函数：
#s = [chr(c) for c in columnTitle]

class Solution(object):
    def titleToNumber(self, columnTitle):
        res =0
        loop=0
        #rune形式遍历string
        s = [ord(c) for c in columnTitle]
        for v  in s :
            #A-Z对应65-90
            #进行进制转换
            v = v - 64
            #v强制以原值转化为int
            s = int(v)

            #相加

            #res = res + v
            #错误在v是rune，res是int，已经改了var rune res = 0
            res = res*26 + s
            loop+=1
        
        return res