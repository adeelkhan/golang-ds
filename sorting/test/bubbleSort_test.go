package same

import (
	"fmt"
	"testing"
)

func Swap(a *int64, b *int64) {
	*a, *b = *b, *a
}

func bubbleSort(arr []int64) {
	for i := len(arr) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] { // ascending sort
				Swap(&arr[j], &arr[j+1])
			}
		}
	}
}

func testEq(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestBubbleSort(t *testing.T) {
	var tests = []struct {
		arr  []int64
		want []int64
	}{
		{arr: []int64{121, 5, 6, 2, 44}, want: []int64{2, 5, 6, 44, 121}},
		{arr: []int64{-121, 5, 6, 2, 44}, want: []int64{-121, 2, 5, 6, 44}},
		{arr: []int64{-121, 5, -6, 2, 44}, want: []int64{-121, -6, 2, 5, 44}},
	}
	fmt.Println("running a")

	for _, tt := range tests {
		testname := fmt.Sprintf("%d", tt.arr)
		t.Run(testname, func(t *testing.T) {
			original := make([]int64, len(tt.arr))
			copy(original, tt.arr)
			bubbleSort(tt.arr)
			if testEq(tt.arr, tt.want) != true {
				t.Errorf("original %d, got %d, want %d", original, tt.arr, tt.want)
			}
		})
	}

}
