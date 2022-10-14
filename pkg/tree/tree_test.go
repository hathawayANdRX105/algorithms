package tree_test

import (
	"algorithms/pkg/tree"
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := tree.NewBinarySearchTree(tree.CmpNumber[int])

	bst.Insert(5)
	tree.Inorder(bst.Root, func(tn *tree.TreeNode[tree.BSTNode, int]) {
		fmt.Printf("%d ->", tn.Val)
	})
	println()

	bst.RemoveOnce(2)
	bst.Insert(2)
	tree.Inorder(bst.Root, func(tn *tree.TreeNode[tree.BSTNode, int]) {
		fmt.Printf("%d ->", tn.Val)
	})
	println()

	bst.Insert(6)
	tree.Inorder(bst.Root, func(tn *tree.TreeNode[tree.BSTNode, int]) {
		fmt.Printf("%d ->", tn.Val)
	})
	println()

	bst.Insert(1)
	tree.Inorder(bst.Root, func(tn *tree.TreeNode[tree.BSTNode, int]) {
		fmt.Printf("%d ->", tn.Val)
	})
	println()
	bst.RemoveOnce(5)

	bst.Insert(3)
	tree.Inorder(bst.Root, func(tn *tree.TreeNode[tree.BSTNode, int]) {
		fmt.Printf("%d ->", tn.Val)
	})
	println()

	// t.Log(tree.IsOrder(bst))
}

func TestAVLTree(t *testing.T) {
	avTree := tree.NewAVLTree(tree.CmpNumber[int])

	avTree.InsertOnce(5)
	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))
	avTree.InsertOnce(2)
	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))
	avTree.InsertOnce(3)
	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))
	avTree.InsertOnce(4)
	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))
	avTree.InsertOnce(4)
	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))
	avTree.InsertOnce(6)

	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))
	// tree.Preorder(avTree.Root, func(node *tree.TreeNode[tree.AVLNode, int]) {
	// 	fmt.Printf("%v - %d  -> ", node.Val, node.Node.Height)
	// })
	// fmt.Println()

	println("523446")
	tree.Preorder(avTree.Root, func(node *tree.TreeNode[tree.AVLNode, int]) {
		fmt.Printf("[val:%v, height:%d, time:%d]  -> \n", node.Val, node.Node.Height, node.Time)
	})
	println()

	avTree.RemoveOnce(4)
	println("52346")
	tree.Preorder(avTree.Root, func(node *tree.TreeNode[tree.AVLNode, int]) {
		fmt.Printf("[val:%v, height:%d, time:%d]  -> \n", node.Val, node.Node.Height, node.Time)
	})
	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))
	println()

	avTree.RemoveOnce(2)
	println("5346")
	tree.Preorder(avTree.Root, func(node *tree.TreeNode[tree.AVLNode, int]) {
		fmt.Printf("[val:%v, height:%d, time:%d]  -> \n", node.Val, node.Node.Height, node.Time)
	})
	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))
	println()

	avTree.RemoveOnce(6)
	println("534")
	tree.Preorder(avTree.Root, func(node *tree.TreeNode[tree.AVLNode, int]) {
		fmt.Printf("[val:%v, height:%d, time:%d]  -> \n", node.Val, node.Node.Height, node.Time)
	})
	t.Logf("is AVL Tree :%t", tree.IsAVLTree(avTree))

}
