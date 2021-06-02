package utils

import (
	"fmt"
	"math"
	"os"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
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

func LoopFileOpen(paths []string) error {
	open_and_close := func(file_path string, write_string string) error {
		fp, err := os.OpenFile(file_path, os.O_RDWR|os.O_CREATE, 0777)
		if err == nil {
			defer fp.Close()
			need_to_write := len([]byte(write_string))
			writed, err := fp.Write([]byte(write_string))
			if err != nil {
				return err
			}
			if writed != need_to_write {
				fmt.Printf("Writed %v bytes out of %v", writed, need_to_write)
			}
		}
		return err
	}
	for _, path := range paths {
		err := open_and_close(path, path)
		if err != nil {
			return err
		}
	}
	return nil
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
