

class Solution(object):
    def getIntersectionNode(self, headA, headB):
        #建立map
        m={}
        #遍历headA,将节点放入map
        while headA!=None:
            m[headA] =True  
            headA = headA.next
        
        #遍历headB，如果节点在map中，则返回
        while headB!=None:
            #*/*******map用法
            if  headB in m:
                return headB
                    
            headB = headB.next
                
        return None
