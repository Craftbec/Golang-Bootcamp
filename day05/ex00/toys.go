package main

import (
	"errors"
	"fmt"
	"log"
)

type TreeNode struct {
	HasToy bool
	Left   *TreeNode
	Right  *TreeNode
}

func count(root *TreeNode, a *int) {
	if root == nil {
		return
	}
	if root.HasToy == true {
		*a += 1
	}
	count(root.Left, a)
	count(root.Right, a)
}

func areToysBalanced(root *TreeNode) (bool, error) {
	var left, right int
	if root == nil {
		return false, errors.New("empty tree")
	}
	count(root.Left, &left)
	count(root.Right, &right)

	if left == right {
		return true, nil
	} else {
		return false, nil
	}
}

func main() {
	var root TreeNode
	root.HasToy = true
	root.Left = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right = &TreeNode{HasToy: true}
	root.Right.Left = &TreeNode{HasToy: true}
	res, err := areToysBalanced(&root)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
