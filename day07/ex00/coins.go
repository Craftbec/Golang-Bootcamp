package test

import (
	"fmt"
	"sort"
)

const MaxInt = 2147483648

func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] > res[j]
	})
	return res
}

func check(coins []int) bool {
	for _, i := range coins {
		if i <= 0 {
			return false
		}
	}
	return true
}

func check_res(table [][]int, val int) []int {
	var res []int
	lenght := MaxInt
	flag := false
	for _, l := range table {
		if len(l) < lenght {
			var sum int
			for _, i := range l {
				sum += i
			}
			if sum == val {
				lenght = len(l)
				res = l
				flag = true
			}

		}
	}
	if !flag {
		for _, l := range table {
			if len(l) < lenght {
				lenght = len(l)
				res = l
			}
		}
	}
	return res
}

func minCoins2(val int, coins []int) []int {
	if !check(coins) {
		fmt.Println("slice must contain only positive values")
		return nil
	}
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	coins = remove_duplicates(coins)
	var table = [][]int{}
	for k, _ := range coins {
		tmp := val
		table = append(table, []int{})
		for i := k; i < len(coins); i++ {
			for tmp >= coins[i] {
				tmp -= coins[i]
				table[k] = append(table[k], coins[i])
			}
		}
	}
	coins = check_res(table, val)
	return coins
}

func remove_duplicates(a []int) (ret []int) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if i > 0 && a[i-1] == a[i] {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

// func main() {
// 	s := []int{4, 1, 3, 3}
// 	// s := []int{2, 3, 4, 6, 8, 10, 12, 14,

// 	// 	16, 18, 20, 22, 24, 26,

// 	// 	29, 31, 33, 35, 37, 39,

// 	// 	41, 43, 45, 47, 49, 51,

// 	// 	52, 54, 56, 58, 60, 62,

// 	// 	64, 66, 68, 70, 72, 74, 75,

// 	// 	77, 79, 81, 83, 85, 87,

// 	// 	89, 91, 93, 95, 97, 99}
// 	// fmt.Println(remove_duplicates(s))
// 	// minCoins2(25, s)
// 	fmt.Println(minCoins(6, s))
// 	fmt.Println(minCoins2(6, s))

// }
