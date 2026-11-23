package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 100} {
		data := generateRandomElements(size)
		want := max(size, 0)
		if len(data) != want {
			t.Errorf("generateRandomElements(%d): длина %d, ожидалась %d", size, len(data), want)
		}

		for i, value := range data {
			if value <= 0 {
				t.Errorf("generateRandomElements(%d)[%d] = %d, ожидалось положительное число", size, i, value)
			}
		}
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"пустой слайс", []int{}, 0},
		{"один элемент", []int{5}, 5},
		{"максимум в начале", []int{9, 1, 3}, 9},
		{"повтор максимума", []int{1, 8, 2, 8}, 8},
		{"отрицательные числа", []int{-7, -2, -5}, -2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		data []int
		want int
	}{
		{nil, 0},
		{[]int{5}, 5},
		{[]int{-3, -1, -2}, -1},
		{[]int{1, 2, 9, 4, 5, 6, 7, 8}, 9},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 9},
	}

	for _, tt := range tests {
		if got := maxChunks(tt.data); got != tt.want {
			t.Errorf("maxChunks(%v) = %d, ожидалось %d", tt.data, got, tt.want)
		}
	}
}
