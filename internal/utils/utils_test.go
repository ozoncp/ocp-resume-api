package utils

import (
	"fmt"
	"testing"

	"github.com/ozoncp/ocp-resume-api/internal/achievement"
	"github.com/ozoncp/ocp-resume-api/internal/resume"
)

func TestSplitBatches(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}
	b_arr, isOk := SplitBatches(arr, 2, false)
	fmt.Printf("%v, %v\n", b_arr, isOk)
	b_arr, isOk = SplitBatches(arr, 2, true)
	fmt.Printf("%v, %v\n", b_arr, isOk)
}

func TestInverseMap(t *testing.T) {
	defer func() {
		if obj := recover(); obj != nil {
			fmt.Println(obj)
		}
	}()
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

func TestLoopFileOpen(t *testing.T) {
	isOk := LoopFileOpen([]string{"qwe.txt", "asd.txt", "zxc.txt"})
	fmt.Printf("%v\n", isOk)
}

func TestSplitAchievementToBatches(t *testing.T) {
	arr := []achievement.Achievement{}
	i := uint(0)
	for {
		if i == 10 {
			break
		}
		tmp_achiv := achievement.New()
		tmp_achiv.Init(i, fmt.Sprintf("Ach%d", i), "Some ach")
		arr = append(arr, *tmp_achiv)
		i += 1
	}
	res, isOk := SplitAchievementToBatches(arr, 3, true)
	fmt.Printf("%v, %v\n", res, isOk)
}

func TestMapAchievements(t *testing.T) {
	arr := []achievement.Achievement{}
	i := uint(0)
	for {
		if i == 10 {
			break
		}
		tmp_achiv := achievement.New()
		tmp_achiv.Init(i, fmt.Sprintf("Ach%d", i), "Some ach")
		arr = append(arr, *tmp_achiv)
		i += 1
	}
	res, isOk := MapAchievements(arr)
	fmt.Printf("%v, %v\n", res, isOk)
}

func TestSplitResumesToBatches(t *testing.T) {
	arr := []resume.Resume{}
	i := uint(0)
	for {
		if i == 10 {
			break
		}
		tmp_achiv := resume.New()
		tmp_achiv.Init(i, i+100)
		arr = append(arr, *tmp_achiv)
		i += 1
	}
	res, isOk := SplitResumesToBatches(arr, 3, true)
	fmt.Printf("%v, %v\n", res, isOk)
}

func TestMapResumes(t *testing.T) {
	arr := []resume.Resume{}
	i := uint(0)
	for {
		if i == 10 {
			break
		}
		tmp_achiv := resume.New()
		tmp_achiv.Init(i, i+100)
		arr = append(arr, *tmp_achiv)
		i += 1
	}
	res, isOk := MapResumes(arr)
	fmt.Printf("%v, %v\n", res, isOk)
}
