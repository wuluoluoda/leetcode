

'''/*罗马数字包含以下七种字符: I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000*/###'''
class Solution:
    #以map设定规则
    SYMBOL_VALUES = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    def romanToInt(self, s) :
        direction=1
        sum=0



    #遍历取出byte
        for i, ch in enumerate(s) :
            value = Solution.SYMBOL_VALUES[ch]
            direction = 1
            #判断是否左右相反
            if i+1<len(s) and value< Solution.SYMBOL_VALUES[s[i + 1]]:
                direction=-1 
        
            #相反则减去
            #否则加上
            sum += direction*value

    #饭回数值
        return sum

