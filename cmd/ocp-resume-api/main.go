package main

import (
	"fmt"

	"github.com/ozoncp/ocp-resume-api/internal/utils"

	"github.com/enescakir/emoji"
)

func main() {
	fmt.Printf("It's resume API written by Pimenov Denis. Hello %v", emoji.WavingHand.Tone(emoji.Light))
	arr_to_split := []int{1, 2, 3, 4, 5}
	b_arr, isOk := utils.SplitBatches(arr_to_split, 2, false)
	fmt.Printf("%v, %v\n", b_arr, isOk)

	map_to_inv := map[uint]string{0: "Zero", 1: "One", 3: "Three"}
	res, isOk := utils.InverseMap(map_to_inv)
	fmt.Printf("%v, %v\n", res, isOk)

	arr_to_filter := []rune{'0', '1', '2', '3', '4', '5'}
	flt_for_arr := []rune{'0', '3', '5'}
	fmt.Printf("In: %v\n", arr_to_filter)
	fmt.Printf("Filter: %v\n", flt_for_arr)
	filtered, isOk := utils.FilterElements(arr_to_filter, flt_for_arr)
	fmt.Printf("%v, %v\n", filtered, isOk)

	_ = utils.LoopFileOpen([]string{"qwe.txt", "asd.txt", "zxc.txt"})
}
