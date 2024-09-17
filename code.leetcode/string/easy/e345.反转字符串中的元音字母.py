import string


class Solution(object):
    def reverseVowels(self, s):
        """
        type s str
        rtype str
        """
        ss = list(s)

        vowels ={
            'a':True,
            'e':True,
            'i':True,
            'o':True,
            'u':True,
            'A':True,
            'E':True,
            'I':True,
            'O':True,
            'U':True,}
        
        #如果存在元音字母，则记录其位置
        pos=[]
        for i, c in enumerate(s) :
            if vowels.get(c) :

                pos.append(i)
            
        
        #对记录的位置进行反转
        i = 0
        while i < len(pos)-1-i :
            ss[pos[i]], ss[pos[len(pos)-1-i]] = ss[pos[len(pos)-1-i]], ss[pos[i]]
            i += 1

        #拼接字符串
        return ''.join(ss)

           