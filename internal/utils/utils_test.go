package utils

import (
	"fmt"
	"testing"
)

func TestSplitBatches(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	b_arr, isOk := SplitBatches(arr, 2, false)
	fmt.Printf("%v, %v", b_arr, isOk)
	b_arr, isOk = SplitBatches(arr, 2, false)
	fmt.Printf("%v, %v", b_arr, isOk)
}
