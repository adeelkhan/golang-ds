package main

import "fmt"

func Swap(a *int64, b *int64) {
	*a, *b = *b, *a
}

func insertionSort(arr []int64) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j-1] > arr[j] {
				Swap(&arr[j-1], &arr[j])
			}
		}
	}
}

func main() {
	arr := []int64{4, 5, 2, 6, 2, -9}
	insertionSort(arr)
	fmt.Println(arr)
}
