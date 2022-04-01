package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res = make([][]int, 0)
	if root == nil {
		return res
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		length := len(queue)
		tmp := make([]int, length)
		for i := 0; i < length; i++ {
			if queue[i] == nil {
				continue
			}
			if len(res)%2 == 0 {
				tmp[i] = queue[i].Val
			} else {
				tmp[length-1-i] = queue[i].Val
			}

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
		res = append(res, tmp)
	}
	return res
}
