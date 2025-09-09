package bigmath

import (
	"fmt"
	"testing"
)

func TestSqrtGMP(t *testing.T) {
	var op FloatPtr
	var rop FloatPtr
	Init2(&op, 100)
	Init2(&rop, 100)
	SetD(&op, float64(2.0), RNDN)
	Sqrt(&rop, &op, RNDN)

	s, n := GetStr(10, 0, &rop, RNDN)
	fmt.Println("1 -->", s, n)
	Printf2("2 --> %.100Rf\n", &rop)

}
func TestSqrt(t *testing.T) {
	tp := 100
	target := 2.0

	op := FloatPtr{}

	//var op FloatPtr
	var rop FloatPtr
	Init2(&op, tp)
	Init2(&rop, tp)
	SetD(&op, target, RNDN)
	Sqrt(&rop, &op, RNDN)
	s, n := GetStr(10, 0, &rop, RNDN)
	fmt.Println("3 -->", s, n)
	//Printf2("yyyyyyyyyyy %10.100Rf\n", &op)
	fmt.Println("^^^^yyyy")
}
func BenchmarkSqrt(b *testing.B) {
	tp := 100
	target := 2.0
	op := FloatPtr{}

	//var op FloatPtr
	var rop FloatPtr
	Init2(&op, tp)
	Init2(&rop, tp)
	SetD(&op, target, RNDN)

	for b.Loop() {
		Sqrt(&rop, &op, RNDN)
	}
}
