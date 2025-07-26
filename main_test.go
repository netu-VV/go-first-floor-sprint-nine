package main

import (
	"testing"
	"time"
)

func TestGenerateRandomElements(t *testing.T) {
	testStruct := []struct {
		name string
		size int
	}{
		{"Zero", 0},
		{"One", 1},
		{"Small", 3},
		{"Large", 10000},
	}

	for _, tt := range testStruct {
		t.Run(tt.name, func(t *testing.T) {
			start := time.Now()
			data := generateRandomElements(tt.size)
			elapsed := time.Since(start)

			if len(data) != tt.size {
				t.Errorf("expected length %d, got %d", tt.size, len(data))
			}

			t.Logf("generated %d elements in %v", tt.size, elapsed)
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"Empty", []int{}, 0},
		{"Single", []int{50}, 50},
		{"Identical", []int{5, 5, 5, 5}, 5},
		{"Positive", []int{1, 5, 2, 9, 3}, 9},
		{"Negatives", []int{-5, -2, -10, -1}, -1},
		{"Mixed", []int{-10, 0, 10, 5}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}
