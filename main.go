package main

import (
	"fmt"
	"math"
	"math/rand/v2" // v2 cuz the seed changes in every single "run"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	res := make([]int, 0, size)
	var value, sign int

	for i := 0; i < size; i++ {
		value = rand.Int()
		sign = rand.IntN(2)
		if sign == 0 {
			res = append(res, -value)
		} else {
			res = append(res, value)
		}
	}
	return res
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if data == nil || len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}

	ans := math.MinInt
	for _, value := range data {
		if value > ans {
			ans = value
		}
	}
	return ans
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	var wg sync.WaitGroup

	maxes := make([]int, CHUNKS)
	chunkLength := len(data) / CHUNKS
	var left, right int

	wg.Add(CHUNKS)
	for i := 0; i < CHUNKS; i++ { // i has iterable scope since Go 1.22 version (*)
		left = chunkLength * i
		right = left + chunkLength

		go func(chunk []int) {
			defer wg.Done()

			maxes[i] = maximum(chunk) // due to (*) there will not be any race-condition & "i" comes from closure
		}(data[left:right])
	}
	wg.Wait()

	return maximum(maxes)
}

func main() {
	var max int

	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max = maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
