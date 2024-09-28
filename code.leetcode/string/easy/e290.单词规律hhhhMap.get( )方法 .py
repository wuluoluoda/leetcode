class Solution(object):
    def wordPattern(self, pattern, s):
        words = s.split(" ")
        #长度不等直接返回False
        if len(pattern) != len(words) :
            return False
        
        patternMap = {}
        wordsMap = {}
        for i in range(len(pattern)):
            #存在值但不对应就返回False
            if (patternMap.get(pattern[i]) and patternMap[pattern[i]] != words[i]) or (wordsMap.get(words[i])  and wordsMap[words[i]] != pattern[i]) :
                return False

            
            patternMap[pattern[i]] = words[i]
            wordsMap[words[i]] = pattern[i]
        
        return True