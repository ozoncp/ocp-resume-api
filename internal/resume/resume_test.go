package resume

import (
	"fmt"
	"testing"
)

func TestResume(t *testing.T) {
	a := MustNew()
	a.Init(1, 33)
	defer a.Close()
	fmt.Println(a.String())
}
