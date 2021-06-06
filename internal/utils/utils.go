package utils

import (
	"fmt"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
)

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

func SplitAchievementsToBatches(sourceArr []achievement.Achievement, batch_size int, align_last bool) ([][]achievement.Achievement, bool) {
	if sourceArr == nil || batch_size <= 0 {
		return nil, false
	}
	src_len := len(sourceArr)
	batch_count := src_len / batch_size
	if src_len%batch_size > 0 {
		batch_count += 1
	}

	res := make([][]achievement.Achievement, batch_count)

	for ndx, src_val := range sourceArr {
		res_ndx := ndx / batch_size
		inside_ndx := ndx % batch_size
		if inside_ndx == 0 {
			if align_last {
				res[res_ndx] = make([]achievement.Achievement, batch_size)
			} else {
				res[res_ndx] = make([]achievement.Achievement, 0, batch_size)
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

func MapAchievements(sourceArr []achievement.Achievement) (map[uint]achievement.Achievement, bool) {
	res := make(map[uint]achievement.Achievement, len(sourceArr))
	for _, src_val := range sourceArr {
		if _, found := res[src_val.Id]; found {
			panic(fmt.Sprintf("Value %v exists twice!", src_val.String()))
			//return nil, false
		}
		res[src_val.Id] = src_val
	}
	return res, true
}

func SplitResumesToBatches(sourceArr []resume.Resume, batch_size int, align_last bool) ([][]resume.Resume, bool) {
	if sourceArr == nil || batch_size <= 0 {
		return nil, false
	}
	src_len := len(sourceArr)
	batch_count := src_len / batch_size
	if src_len%batch_size > 0 {
		batch_count += 1
	}

	res := make([][]resume.Resume, batch_count)

	for ndx, src_val := range sourceArr {
		res_ndx := ndx / batch_size
		inside_ndx := ndx % batch_size
		if inside_ndx == 0 {
			if align_last {
				res[res_ndx] = make([]resume.Resume, batch_size)
			} else {
				res[res_ndx] = make([]resume.Resume, 0, batch_size)
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

func MapResumes(sourceArr []resume.Resume) (map[uint]resume.Resume, bool) {
	res := make(map[uint]resume.Resume, len(sourceArr))
	for _, src_val := range sourceArr {
		if _, found := res[src_val.Id]; found {
			panic(fmt.Sprintf("Value %v exists twice!", src_val.String()))
			//return nil, false
		}
		res[src_val.Id] = src_val
	}
	return res, true
}

func SaveAchievements(dst []achievement.Achievement, src []achievement.Achievement, capacity int, samart_del bool) []achievement.Achievement {
	if len(dst)+len(src) > capacity {
		del_count := int(0)
		if samart_del {
			del_count = capacity - len(dst) - len(src)
		} else {
			del_count = len(dst) - 1
		}
		copy(dst, dst[del_count:])
		dst = dst[:len(dst)-del_count]
	}
	return append(dst, src...)
}

func SaveResumes(dst []resume.Resume, src []resume.Resume, capacity int, samart_del bool) []resume.Resume {
	if len(dst)+len(src) > capacity {
		del_count := int(0)
		if samart_del {
			del_count = capacity - len(dst) - len(src)
		} else {
			del_count = len(dst) - 1
		}
		copy(dst, dst[del_count:])
		dst = dst[:len(dst)-del_count]
	}
	return append(dst, src...)
}
