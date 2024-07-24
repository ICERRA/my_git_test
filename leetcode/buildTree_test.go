// Package decode 该文件下均为树的题目
package leetcode

import (
	"math"
	"slices"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// region buildTree 前序中序构建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	// 获取左子树的长度
	leftTreeSize := slices.Index(inorder, preorder[0])
	leftTree := buildTree(preorder[1:1+leftTreeSize], inorder[:leftTreeSize])
	rightTree := buildTree(preorder[1+leftTreeSize:], inorder[1+leftTreeSize:])

	var rootNode = &TreeNode{
		Val:   preorder[0],
		Left:  leftTree,
		Right: rightTree,
	}

	return rootNode
}

//endregion

// region 二叉树的最大深度

// maxDepth 深度优先
func maxDepth(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return 1
	}

	if root.Left == nil {
		return maxDepth(root.Right) + 1
	}

	if root.Right == nil {
		return maxDepth(root.Left) + 1
	}

	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right)))) + 1
}

// maxDepth1 广度优先
func maxDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var (
		deepth int
		queue  = make([]*TreeNode, 0)
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			rootNode := queue[0]
			queue = queue[1:]

			if rootNode.Left != nil {
				queue = append(queue, rootNode.Left)
			}
			if rootNode.Right != nil {
				queue = append(queue, rootNode.Right)
			}
			size--
		}
		deepth++
	}
	return deepth
}

//endregion

// region 二叉搜索树中的第k小元素
func kthSmallest(root *TreeNode, k int) int {
	var inorder []int
	var doFunc func(root *TreeNode)
	doFunc = func(root *TreeNode) {
		if root == nil {
			return
		}
		doFunc(root.Left)
		inorder = append(inorder, root.Val)
		doFunc(root.Right)
	}
	doFunc(root)
	if len(inorder) >= k {
		return inorder[k-1]
	}

	return 0
}

//endregion
