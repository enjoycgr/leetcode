package main

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == q || root == p {
		return root
	}

	left := lowestCommonAncestor2(root.Left, p, q)
	right := lowestCommonAncestor2(root.Right, q, p)

	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	return root
}
