package utils

import (
	"fmt"
	"testing"
)

func TestSplitBatches(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	b_arr, isOk := SplitBatches(arr, 2, false)
	fmt.Printf("%v, %v\n", b_arr, isOk)
	b_arr, isOk = SplitBatches(arr, 2, true)
	fmt.Printf("%v, %v\n", b_arr, isOk)
}

func TestInverseMap(t *testing.T) {
	tst := map[uint]string{0: "Zero", 1: "One", 3: "Three"}
	res, isOk := InverseMap(tst)
	fmt.Printf("%v, %v\n", res, isOk)
	tst = map[uint]string{}
	res, isOk = InverseMap(tst)
	fmt.Printf("%v, %v\n", res, isOk)
	tst = map[uint]string{0: "Zero", 1: "Zero"}
	res, isOk = InverseMap(tst)
	fmt.Printf("%v, %v\n", res, isOk)
}

func TestFilterElements(t *testing.T) {
	arr := []rune{'0', '1', '2', '3', '4', '5'}
	flt := []rune{'0', '3', '5'}
	fmt.Printf("In: %v\n", arr)
	fmt.Printf("Filter: %v\n", flt)
	arr, isOk := FilterElements(arr, flt)
	fmt.Printf("%v, %v\n", arr, isOk)
}
