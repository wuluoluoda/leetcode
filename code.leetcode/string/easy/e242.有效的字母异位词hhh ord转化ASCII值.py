"""
func isAnagram(s, t string) bool {
    var c1, c2 [26]int
    for _, ch := range s {
        c1[ch-'a']++
    }
    for _, ch := range t {
        c2[ch-'a']++
    }
    return c1 == c2
}


"""
import string


class Solution(object):
    def isAnagram(self, s, t):
        if len(s) != len(t) :
            return False
	
        c1 = [0] * 26
        c2 = [0] * 26        
        for ch in s:
            c1[ord(ch.lower()) - ord('a')] += 1
        
        for ch in t:
            c2[ord(ch.lower()) - ord('a')] += 1
        
        return c1 == c2
