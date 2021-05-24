package utils

import (
	"fmt"
	"math"
)

func SplitBatches(sourceArr []int, batch_size int, align_last bool) ([][]int, bool) {
	if sourceArr == nil || batch_size <= 0 {
		return nil, false
	}
	src_len := len(sourceArr)
	batch_count := int(math.Ceil(float64(src_len) / float64(batch_size)))

	res := make([][]int, batch_count)

	for ndx, src_val := range sourceArr {
		res_ndx := ndx / batch_size
		inside_ndx := ndx % batch_size
		if inside_ndx == 0 {
			if align_last {
				res[res_ndx] = make([]int, batch_size)
			} else {
				res[res_ndx] = make([]int, 0, batch_size)
			}

		}
		if align_last {
			res[res_ndx][inside_ndx] = src_val
		} else {
			res[res_ndx] = append(res[res_ndx], src_val)
		}

	}

	return res, true
}

func InverseMap(sourceMap map[uint]string) (map[string]uint, bool) {
	res := make(map[string]uint, len(sourceMap))
	for src_key, src_val := range sourceMap {
		if _, found := res[src_val]; found {
			panic(fmt.Sprintf("Value %v exists twice!", src_val))
			//return nil, false
		}
		res[src_val] = src_key
	}
	return res, true
}

func FilterElements(sourceArray []rune, filterArray []rune) ([]rune, bool) {
	res := make([]rune, 0, len(sourceArray))
	for _, src_val := range sourceArray {
		found := false
		for _, flt_val := range filterArray {
			if src_val == flt_val {
				found = true
				break
			}
		}
		if !found {
			res = append(res, src_val)
		}
	}
	return res, true
}
