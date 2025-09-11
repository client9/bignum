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

func BenchmarkInt(b *testing.B) {
	for b.Loop() {
		a := NewInt(1000)
		b := NewInt(100)
		a.Add(a, b)
		b.Add(a, b)
	}
}
func BenchmarkIntTmp(b *testing.B) {
	for b.Loop() {
		a := NewIntTmp(1000)
		b := NewIntTmp(100)
		defer a.Clear()
		defer b.Clear()
		a.Add(a, b)
		b.Add(a, b)
	}
}
