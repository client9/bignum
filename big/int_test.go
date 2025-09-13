package big

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	a := NewInt(1000)
	b := NewInt(10)
	a.Add(a, b)
	b.Add(a, b)
	fmt.Println(a.String())
}

// should not explode
func TestIntZeroValue(t *testing.T) {
	var a, b Int
	a.Add(&a, &b)
}
