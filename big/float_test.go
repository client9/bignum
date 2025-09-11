package big

import (
	"fmt"
	"testing"
)

func TestFloat(t *testing.T) {
	a := NewFloat(3.141498249073492873942793472 * 10000000000000)
	b := NewFloat(10.0)
	a.Add(a, b)
	b.Add(a, b)
	fmt.Println(a.String())
}
