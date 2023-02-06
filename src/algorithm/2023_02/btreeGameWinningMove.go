package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// n2 二号玩家的节点个数，lsz 左子树，rsz 右子树 父节点 n-1-lsz-rsz
func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	lsz, rsz := 0, 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ls := dfs(node.Left)
		rs := dfs(node.Right)
		if node.Val == x {
			lsz, rsz = ls, rs
		}
		return ls + rs + 1
	}
	dfs(root)
	return max(max(lsz, rsz), n-1-lsz-rsz)*2 > n

}
func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {

}
