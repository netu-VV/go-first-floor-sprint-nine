package main

import (
	"fmt"
	"math/rand"
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
		return []int{}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, size)
	for i := range nums {
		nums[i] = r.Intn(size * 10)
	}
	return nums
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) <= 0 {
		return 0
	}

	max := data[0]

	if len(data) == 1 {
		return max
	}

	for _, m := range data[1:] {
		if m > max {
			max = m
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) <= 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup
	maxes := make([]int, CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			start := v * chunkSize
			end := start + chunkSize
			if v == CHUNKS-1 {
				end = len(data)
			}

			maxes[v] = maximum(data[start:end])
		}(i)
	}
	wg.Wait()
	return maximum(maxes)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
