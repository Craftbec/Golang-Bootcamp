package main

import (
	"fmt"
)

type Present struct {
	Weight int
	Price  int
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func trace_result(table [][]int, presents []Present, k, s int, res *[]Present) {
	if table[k][s] == 0 {
		return
	}
	if table[k-1][s] == table[k][s] {
		trace_result(table, presents, k-1, s, res)
	} else {
		*res = append(*res, presents[k-1])
		trace_result(table, presents, k-1, s-presents[k-1].Weight, res)
	}

}

func grabPresents(presents []Present, n int) []Present {
	var res []Present
	if len(presents) == 0 {
		return res
	}
	var table = make([][]int, len(presents)+1)
	for i := range table {
		table[i] = make([]int, n+1)
	}

	for i := range table {
		for j := range table[i] {
			if i == 0 || j == 0 {
				table[i][j] = 0
			} else {
				if j >= presents[i-1].Weight {
					table[i][j] = Max(table[i-1][j], (table[i-1][j-presents[i-1].Weight] + presents[i-1].Price))
				} else {
					table[i][j] = table[i-1][j]
				}
			}
		}
	}
	trace_result(table, presents, len(presents), n, &res)
	return res
}

func main() {
	var n int = 4
	presents := []Present{
		Present{Weight: 1, Price: 1500},
		Present{Weight: 4, Price: 3000},
		Present{Weight: 3, Price: 2000},
	}
	res := grabPresents(presents, n)
	fmt.Printf("%+v", res)
}
