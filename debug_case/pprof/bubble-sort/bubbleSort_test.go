package main

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	//	arr := []int{64, 34, 25, 12, 22, 11, 90}

	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 3, 2, 1, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{5, 6, 2, 8, 4}, []int{2, 4, 5, 8, 6}},
	}

	for _, test := range tests {
		if got := bubbleSort(test.nums); !reflect.DeepEqual(got, test.want) {
			t.Errorf("bubbleSort(%v) = %v, want %v", test.nums, got, test.want)
		}
	}

}
