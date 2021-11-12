package main

import "fmt"

func Swap(a *int64, b *int64) {
	*a, *b = *b, *a
}

func find_pivot(arr []int64, start int, end int) int {
	pivot := end
	i := start
	j := pivot - 1

	for i <= j {
		if arr[pivot] < arr[i] {
			Swap(&arr[i], &arr[j])
			j--
		} else if arr[pivot] > arr[j] {
			Swap(&arr[i], &arr[j])
			i++
		} else {
			i++
			j--
		}
	}
	Swap(&arr[i], &arr[pivot])
	return i
}

func qs(arr []int64, start int, end int) {
	if start >= end {
		return
	}
	pivot := find_pivot(arr, start, end)

	qs(arr, start, pivot-1)
	qs(arr, pivot+1, end)

}
func quickSort(arr []int64) {
	qs(arr, 0, len(arr)-1)
}

func main() {
	arr := []int64{4, 5, 2, 6, 12, -9}
	quickSort(arr)
	fmt.Println(arr)
}
