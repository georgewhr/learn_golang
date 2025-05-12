package internal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (this *TreeNode) InsertLeft(data int, root *TreeNode) *TreeNode {

	if root == nil {
		return &TreeNode{Val: data}
	}

	root.Left = this.InsertLeft(data, root.Left)
	return root

}

func (this *TreeNode) InsertRight(data int, root *TreeNode) *TreeNode {
	if root == nil {
		return &TreeNode{Val: data}
	}

	root.Right = this.InsertLeft(data, root.Right)
	return root
}

func (this *TreeNode) ConstructBST(data int, root *TreeNode) *TreeNode {

	if root == nil {
		return &TreeNode{Val: data}
	}

	if data > root.Val {
		root.Right = this.ConstructBST(data, root.Right)
	} else {
		root.Left = this.ConstructBST(data, root.Left)
	}

	return root

}
