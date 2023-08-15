package test

import (
	"testing"
)

var (
	test_coin = [][]int{
		{1, 3, 4, 7, 13, 15, 1, 3, 4, 7},
		{1, 5, 8, 10},
		{2, 3, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 29, 31, 33, 35, 37, 39,
			41, 43, 45, 47, 49, 51, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 75,
			77, 79, 81, 83, 85, 87, 89, 91, 93, 95, 97, 99},
		{7, 8, 9, 12, 23, 43},
		{},
		{9, 89, 678, 7, 34, 45, 57, 34, 56, 765, 32, 9, 89, 678, 73, 334, 454, 577, 364, 556, 7635, 322},
	}
	n = []int{25, 13, 88, 6, 123, 5}
)

// func BenchmarkMinCoins(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for j := 0; j < 6; j++ {
// 			minCoins(n[j], test_coin[j])
// 		}

// 	}
// }

func BenchmarkMinCoins2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 6; j++ {
			minCoins2(n[j], test_coin[j])
		}
	}
}
func BenchmarkMinCoins2Optimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for j := 0; j < 6; j++ {
			minCoins2Optimized(n[j], test_coin[j])
		}
	}
}
