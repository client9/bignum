package mpfr

import (
	"testing"
)

func TestManyDigits1(t *testing.T) {
	prec := int(100000 * 3.4)

	var op FloatPtr
	var rop FloatPtr

	Init2(&rop, prec)
	Init2(&op, prec)

	SetD(&op, float64(1.0), RNDN)
	Cos(&rop, &op, RNDN)
	Tan(&rop, &rop, RNDN)
	Sin(&rop, &rop, RNDN)
	/*
		s, n := GetStr(10, 0, &rop, RNDN)
		fmt.Println(n, len(s), s)
	*/

}
func BenchmarkManyDigits1(b *testing.B) {
	prec := int(100000 * 3.4)

	var op FloatPtr
	var rop FloatPtr

	Init2(&rop, prec)
	Init2(&op, prec)
	for b.Loop() {

		SetD(&op, float64(1.0), RNDN)
		Cos(&rop, &op, RNDN)
		Tan(&rop, &rop, RNDN)
		Sin(&rop, &rop, RNDN)
		GetStr(10, 0, &rop, RNDN)
	}

}
