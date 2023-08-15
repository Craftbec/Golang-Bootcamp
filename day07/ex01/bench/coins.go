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
	var res, tmp_res []int
	lenght := MaxInt
	flag := false
	tmp_l := MaxInt
	for _, l := range table {
		if len(l) < tmp_l {
			lenght = len(l)
			tmp_res = l
		}
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
		res = tmp_res
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

func minCoins2Optimized(val int, coins []int) []int {
	if len(coins) == 0 {
		return coins
	}
	if !check(coins) {
		fmt.Println("slice must contain only positive values")
		return nil
	}
	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})
	if coins[len(coins)-1] > val {
		return []int{}
	}
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
