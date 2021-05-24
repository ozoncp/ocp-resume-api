package achievement

import (
	"fmt"
	"testing"
)

func TestAchievement(t *testing.T) {
	a := MustNew()
	a.Init(1, "First", "First struct!")
	defer a.Close()
	fmt.Println(a.String())
}
