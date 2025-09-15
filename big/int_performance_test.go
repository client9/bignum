package big

import (
	"fmt"
	"runtime"
	"sync"
	"testing"

	stdlib "math/big"
)

// Compares many digits addition, multiplication, and serialization
//
// TODO: division would be interesting too
//
//
// Inspired from tests in https://github.com/ncw/gmp
//

// Converting binary representation to decimal is hard work!
func BenchmarkStdToString(b *testing.B) {
	b.Skip()
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberStd(i)
			for b.Loop() {
				x.String()
			}
		})
	}
}

func BenchmarkGmpToString(b *testing.B) {
	b.Skip()
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberGmp(i)
			for b.Loop() {
				x.String()
			}
		})
	}
}

func BenchmarkStdAddRaw(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberStd(i)
			y := nDigitNumberStd(i)
			z := stdlib.NewInt(0)
			for b.Loop() {
				z.Add(x, y)
			}
		})
	}
}
func BenchmarkGmpAddRaw(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberGmp(i)
			y := nDigitNumberGmp(i)
			z := NewInt(0)
			for b.Loop() {
				z.Add(x, y)
			}
		})
	}
}
func BenchmarkStdAddAlloc(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberStd(i)
			y := nDigitNumberStd(i)
			for b.Loop() {
				z := stdlib.NewInt(0)
				z.Add(x, y)
			}
		})
	}
}
func BenchmarkGmpAddAlloc(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberGmp(i)
			y := nDigitNumberGmp(i)
			for b.Loop() {
				z := NewInt(0)
				z.Add(x, y)
			}
		})
	}
}
func BenchmarkStdMulRaw(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			runtime.GC()
			runtime.GC()
			x := nDigitNumberStd(i)
			y := nDigitNumberStd(i)
			for b.Loop() {
				z := stdlib.NewInt(0)
				z.Mul(x, y)
			}
		})
	}
}
func BenchmarkGmpMulRaw(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			runtime.GC()
			runtime.GC()
			x := nDigitNumberGmp(i)
			y := nDigitNumberGmp(i)
			z := NewInt(0)
			for b.Loop() {
				z.Mul(x, y)
			}
		})
	}
}
func BenchmarkStdMulAlloc(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			runtime.GC()
			runtime.GC()
			x := nDigitNumberStd(i)
			y := nDigitNumberStd(i)
			for b.Loop() {
				z := stdlib.NewInt(0)
				z.Mul(x, y)
			}
		})
	}
}

func BenchmarkGmpMulAlloc(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			runtime.GC()
			runtime.GC()
			x := nDigitNumberGmp(i)
			y := nDigitNumberGmp(i)
			for b.Loop() {
				z := NewInt(0)
				z.Mul(x, y)
			}
		})
	}
}

func BenchmarkGmpMulDefer(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberGmp(i)
			y := nDigitNumberGmp(i)
			//z := NewInt(0)
			for b.Loop() {
				z := NewIntTmp(0)
				defer z.Clear()
				z.Mul(x, y)
			}
		})
	}
}

func BenchmarkGmpMulPool(b *testing.B) {

	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("%d", i)
		b.Run(name, func(b *testing.B) {
			runtime.GC()
			runtime.GC()
			bufferPool := sync.Pool{
				New: func() any { return NewInt(0) },
			}

			x := nDigitNumberGmp(i)
			y := nDigitNumberGmp(i)
			for b.Loop() {
				z := bufferPool.Get().(*Int)
				z.Mul(x, y)
				bufferPool.Put(z)
			}
		})
	}
}

// Make a n digit number
func nDigitNumberGmp(digits int64) *Int {
	x := NewInt(10)
	n := NewInt(digits)
	one := NewInt(1)
	x.Exp(x, n, nil)
	x.Sub(x, one)
	return x
}

// Make a n digit number
func nDigitNumberStd(digits int64) *stdlib.Int {
	x := stdlib.NewInt(10)
	n := stdlib.NewInt(digits)
	one := stdlib.NewInt(1)
	x.Exp(x, n, nil)
	x.Sub(x, one)
	return x
}
