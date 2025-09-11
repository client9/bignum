package big

import (
	"fmt"
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
func BenchmarkStdLibToString(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("StdLibStr%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberStdLib(i)
			for b.Loop() {
				x.String()
			}
		})
	}
}

func BenchmarkGMPToString(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("GmpStr%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberGmp(i)
			for b.Loop() {
				x.String()
			}
		})
	}
}

func BenchmarkStdlibAdd(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("StdLibAdd%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberStdLib(i)
			y := nDigitNumberStdLib(i)
			z := stdlib.NewInt(0)
			for b.Loop() {
				z.Add(x, y)
			}
		})
	}
}
func BenchmarkGmpAdd(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("GmpAdd%d", i)
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
func BenchmarkStdlibMul(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("StdLibMul%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberStdLib(i)
			y := nDigitNumberStdLib(i)
			z := stdlib.NewInt(0)
			for b.Loop() {
				z.Mul(x, y)
			}
		})
	}
}

func BenchmarkGmpMul(b *testing.B) {
	for i := int64(10); i <= 1000000; i *= 10 {
		name := fmt.Sprintf("GmpMul%d", i)
		b.Run(name, func(b *testing.B) {
			x := nDigitNumberGmp(i)
			y := nDigitNumberGmp(i)
			z := NewInt(0)
			for b.Loop() {
				z.Mul(x, y)
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
func nDigitNumberStdLib(digits int64) *stdlib.Int {
	x := stdlib.NewInt(10)
	n := stdlib.NewInt(digits)
	one := stdlib.NewInt(1)
	x.Exp(x, n, nil)
	x.Sub(x, one)
	return x
}
