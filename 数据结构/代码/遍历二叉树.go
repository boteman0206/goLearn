package main

import "fmt"

/**
https://juejin.cn/post/7025577339119468557

*/
//数据结构
type TreeNode struct {
	Val   string
	Left  *TreeNode
	Right *TreeNode
}

// 前序遍历的遍历顺序为： 根节点->左子树->右子树
func preorderTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Println(node.Val)
	preorderTraveral(node.Left)
	preorderTraveral(node.Right)
}

// 中序遍历 中序遍历的遍历顺序为： 左子树->根节点->右子树
func midTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	midTraveral(node.Left)

	fmt.Println(node.Val)
	midTraveral(node.Right)
}

//后序遍历的遍历顺序为： 左子树->右子树->根节点
func lastTraveral(node *TreeNode) {
	if node == nil {
		return
	}
	lastTraveral(node.Left)
	lastTraveral(node.Right)
	fmt.Println(node.Val)
}

func main() {

	data := TreeNode{
		Val: "A",
		Left: &TreeNode{
			Val: "B",
			Left: &TreeNode{
				Val:   "D",
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val: "E",
				Left: &TreeNode{
					Val:   "H",
					Left:  nil,
					Right: nil,
				},
				Right: &TreeNode{
					Val:   "I",
					Left:  nil,
					Right: nil,
				},
			},
		},
		Right: &TreeNode{
			Val: "C",
			Left: &TreeNode{
				Val:   "F",
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   "G",
				Left:  nil,
				Right: nil,
			},
		},
	}
	//preorderTraveral(&data)
	//midTraveral(&data)
	lastTraveral(&data)
}
