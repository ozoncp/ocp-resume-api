package utils

import "math"

func SplitBatches(sourceArr []int, batch_size int, align_last bool) ([][]int, bool) {
	if sourceArr == nil || batch_size <= 0 {
		return nil, false
	}
	src_len := len(sourceArr)
	batch_count := int(math.Ceil(float64(src_len) / float64(batch_size)))

	res := make([][]int, batch_count, batch_count)

	for ndx, src_val := range sourceArr {
		res_ndx := ndx / batch_size
		inside_ndx := ndx % batch_size
		if inside_ndx == 0 {
			if align_last {
				res[res_ndx] = make([]int, batch_size, batch_size)
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
