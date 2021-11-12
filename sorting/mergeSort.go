package main

import "fmt"

func ms(arr []int64, a, b int) []int64 {
	if a >= b {
		a := []int64{arr[a]}
		return a
	}
	mid := (a + b) / 2
	left := ms(arr, a, mid)
	right := ms(arr, mid+1, b)
	return merge(left, right)
}

func merge(left []int64, right []int64) []int64 {
	arr := make([]int64, len(left)+len(right))
	i := 0
	j := 0
	pos := 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			arr[pos] = left[i]
			i++
			pos++
		} else {
			arr[pos] = right[j]
			j++
			pos++
		}
	}
	if i < len(left) {
		for ; i < len(left); i++ {
			arr[pos] = left[i]
			pos++
		}
	}
	if j < len(right) {
		for ; j < len(right); j++ {
			arr[pos] = right[j]
			pos++
		}
	}
	return arr
}
func mergeSort(arr []int64) []int64 {
	return ms(arr, 0, len(arr)-1)
}

func main() {
	arr := []int64{4, 5, 2, 16, 2, -9}
	fmt.Println(mergeSort(arr))
}
