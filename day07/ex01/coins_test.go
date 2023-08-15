package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAverage(t *testing.T) {
	test_coin := [][]int{
		{1, 3, 4, 7, 13, 15, 1, 3, 4, 7},
		{1, 5, 8, 10},
		{2, 3, 4, 6, 8, 10, 12, 14,

			16, 18, 20, 22, 24, 26,

			29, 31, 33, 35, 37, 39,

			41, 43, 45, 47, 49, 51,

			52, 54, 56, 58, 60, 62,

			64, 66, 68, 70, 72, 74, 75,

			77, 79, 81, 83, 85, 87,

			89, 91, 93, 95, 97, 99},
		{},
		{7, 8, 9, 12, 23, 43},
		{2, 4, 6},
	}
	result_coin := [][]int{
		{7, 15, 3},
		{5, 8},
		{85, 3},
		{},
		{},
		{4},
	}

	n := []int{
		25, 13, 88, 12, 6, 5,
	}

	assert := assert.New(t)
	for i, _ := range test_coin {
		assert.Equal(len(minCoins(n[i], test_coin[i])), len(result_coin[i]), "minCoins")
		assert.Equal(len(minCoins2(n[i], test_coin[i])), len(result_coin[i]), "minCoins2")
	}
}
