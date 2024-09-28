//递归过程左中右
func inorderTraversal(root *TreeNode) []int {
	var s []int
    inorder(root, &s)
    return s 
}
//递归函数
func inorder(node *TreeNode,s *[]int) {

//基础情况
if node == nil {
	return
}	
//递归过程
		//遍历左子树s
		inorder(node.Left,s)
		//遍历根节点（输出数组）
		*s=append(*s,node.Val)//这里为什么需要用指针,这不是指针的指针吗,说是直接
								//s等于对指针的拷贝,但是指针还能拷贝吗
		//遍历右子树
		inorder(node.Right,s)
		    
}
//官方写法,不需要指针
/*
func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		res = append(res, node.Val)
		inorder(node.Right)
	}
	inorder(root)
	return
}

*/