package big

import (
	"fmt"
	"testing"
)

func TestRat(t *testing.T) {
	a := NewRat(1, 3)
	b := NewRat(1, 2)
	a.Add(a, b)
	fmt.Println(a.String())
}
