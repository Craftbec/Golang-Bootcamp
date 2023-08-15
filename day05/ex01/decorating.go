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

func right_left(tmp []TreeNode, right bool) []TreeNode {
	var next []TreeNode
	for len(tmp) != 0 {
		x := tmp[len(tmp)-1]
		tmp = tmp[:len(tmp)-1]
		if right {
			if x.Right != nil {
				next = append(next, *x.Right)
			}
		}
		if x.Left != nil {
			next = append(next, *x.Left)
		}
		if !right {
			if x.Right != nil {
				next = append(next, *x.Right)
			}
		}
	}
	return next
}

func unrollGarland(root *TreeNode) ([]bool, error) {
	var tmp []TreeNode
	if root == nil {
		return []bool{}, errors.New("empty tree")
	}
	res := []bool{}
	tmp = append(tmp, *root)
	for i := 0; ; i++ {
		for _, q := range tmp {
			res = append(res, q.HasToy)
		}
		if i%2 == 0 {
			tmp = right_left(tmp, false)
		} else {
			tmp = right_left(tmp, true)
		}
		if tmp == nil {
			break
		}
	}

	return res, nil
}
func main() {
	var root TreeNode
	root.HasToy = true
	root.Left = &TreeNode{HasToy: true}
	root.Left.Left = &TreeNode{HasToy: true}
	root.Left.Left.Left = &TreeNode{HasToy: true}
	root.Left.Left.Right = &TreeNode{HasToy: true}
	root.Left.Right = &TreeNode{HasToy: false}
	root.Right = &TreeNode{HasToy: false}
	root.Right.Left = &TreeNode{HasToy: true}
	root.Right.Right = &TreeNode{HasToy: true}
	root.Right.Right.Left = &TreeNode{HasToy: true}
	res, err := unrollGarland(&root)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
