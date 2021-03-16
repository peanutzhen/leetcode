package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 思路是这样子的
// 利用BST的中序遍历有序 我们发现
// 因为只是其中两个节点值错误 那么就肯定存在两次逆序
// 比如测试用例1：[1,3,null,null,2]
// 中序遍历 3 2 1 两次逆序在 3->2 2->1
// 我们只需在遍历的时候保存这两个节点的指针 之后交换值即可
// 注意：有时两个逆序对是同一个
// 如：[3,1,4,null,null,2] 遍历结果 1 3 2 4
// 我们可以先记录着 3与2 假设后面还有逆序对 再把2修改成那个节点即可
func recoverTree(root *TreeNode) {
	count := 0
	var prev *TreeNode // 记录遍历时的前一个节点
	var foo, bar *TreeNode

	goahead := true
	var inorderTraversal func(root *TreeNode)
	inorderTraversal = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorderTraversal(root.Left)

		if prev != nil && prev.Val > root.Val {
			if count == 0 { // 第一次逆序对
				foo, bar = prev, root
				count++
			} else {
				bar = root
				goahead = false
			}
		}
		prev = root

		// 已经找到两个逆序对 不必再遍历
		if goahead {
			inorderTraversal(root.Right)
		}
	}
	inorderTraversal(root)
	foo.Val, bar.Val = bar.Val, foo.Val
}
