package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	options := map[string][]interface{}{
		"irregular_empty":    {0, nil},
		"regular_medium":     {1000, 1000},
		"irregular_negative": {-7, nil},
		"regular_small":      {4, 4},
		"regular_large":      {10000000, 10000000},
	}

	for name, option := range options {
		t.Run(name, func(t *testing.T) {
			size := option[0].(int)
			actual := generateRandomElements(size)
			switch expected := option[1].(type) {
			case nil:
				require.Nil(t, actual)
			case int:
				require.Equal(t, expected, len(actual))
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	options := map[string][]interface{}{
		"irregular_empty": {[]int{}, 0},
		"regular_single":  {[]int{2025}, 2025},
		"regular_small":   {[]int{4, 8, 15, 16, 23, 42}, 42},
		"regular_medium": {[]int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9,
			10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
			20, 21, 22, 23, 24, 72, 26, 27, 28, 29,
			30, 31, 32, 33, 34, 35, 36, 37, 38, 39,
			40, 41, 42, 43, 44, 45, 46, 47, 48, 49,
		}, 72},
		"regular_large": {[]int{
			0, 5, 12, 18, 27, 31, 44, 52, 68, 75, 83,
			97, 105, 116, 129, 141, 150, 162, 174, 185,
			199, 205, 217, 230, 244, 255, 269, 280, 294,
			305, 318, 330, 342, 356, 367, 379, 390, 404,
			415, 427, 439, 450, 462, 478, 489, 503, 514,
			526, 538, 550, 561, 577, 588, 600, 612, 624,
			635, 647, 659, 672, 684, 695, 707, 720,
		}, 720},
	}

	for name, option := range options {
		t.Run(name, func(t *testing.T) {
			expected := option[1].(int)
			switch data := option[0].(type) {
			case []int:
				require.Equal(t, expected, maximum(data))
			}
		})
	}
}
