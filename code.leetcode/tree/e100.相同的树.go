func isSameTree(p *TreeNode, q *TreeNode) bool {
	var inorderTraversal func(root *TreeNode)
	inorderTraversal=func (root *TreeNode) (res []int) {
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
	return reflect.DeepEqual(inorderTraversal(p),inorderTraversal(q))
}//切片相同，则树相同