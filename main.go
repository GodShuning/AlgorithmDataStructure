package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	helper(&res, root)
	return res
}

func helper(res *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	helper(res, root.Left)
	helper(res, root.Right)
}
