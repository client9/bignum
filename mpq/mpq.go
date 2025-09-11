package mpq

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lgmp

#include <stdlib.h>
#include <stdio.h>

#include "gmp.h"
int macro_mpq_cmp_si(const mpq_t op, long int num2, unsigned long den2) {
	return mpq_cmp_si(op, num2, den2);
}
int macro_mpq_cmp_ui(const mpq_t op, unsigned long int num2, unsigned long int den2) {
	return mpq_cmp_ui(op, num2, den2);
}
int macro_mpq_sgn(const mpq_t op) {
        return mpq_sgn(op);
}
*/
import "C"

import (
	"unsafe"

	"github.com/client9/bignum/mpz"
)

type Rat C.mpq_t

// used internally to help CGO understand that mpz.Int is the same thing as C.mpz_t
type ratInt C.mpz_t

//
// 6.1 Initialization and Assignment Functions
//

func Init(x *Rat) {
	C.mpq_init(&x[0])
}

func Clear(x *Rat) {
	C.mpq_clear(&x[0])
}

func Set(rop *Rat, op *Rat) {
	C.mpq_set(&rop[0], &op[0])
}

// TODO mpq_set_z

func SetUi(rop *Rat, op1 uint, op2 uint) {
	C.mpq_set_ui(&rop[0], C.ulong(op1), C.ulong(op2))
}

func SetSi(rop *Rat, op1 int, op2 uint) {
	C.mpq_set_si(&rop[0], C.long(op1), C.ulong(op2))
}

func SetStr(rop *Rat, s string, base int) int {
	cstr := C.CString(s)
	ret := C.mpq_set_str(&rop[0], cstr, C.int(base))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

func Swap(rop1 *Rat, rop2 *Rat) {
	C.mpq_swap(&rop1[0], &rop2[0])
}

// 6.2 Conversion Functions

func GetD(rop *Rat) float64 {
	return float64(C.mpq_get_d(&rop[0]))
}

// TODO mpq_set_d
// TODO mpq_set_f

// GetStr is not part of the GMP library.
func GetStr(base int, op *Rat) string {
	p := C.mpq_get_str(nil, C.int(base), &op[0])
	if p == nil {
		return ""
	}
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

//
// 6.3 Arithmetic Functions
//

func Add(rop *Rat, op1 *Rat, op2 *Rat) {
	C.mpq_add(&rop[0], &op1[0], &op2[0])
}
func Sub(rop *Rat, op1 *Rat, op2 *Rat) {
	C.mpq_sub(&rop[0], &op1[0], &op2[0])
}
func Mul(rop *Rat, op1 *Rat, op2 *Rat) {
	C.mpq_mul(&rop[0], &op1[0], &op2[0])
}

// TODO mpq_mul_2exp

func Div(rop *Rat, op1 *Rat, op2 *Rat) {
	C.mpq_div(&rop[0], &op1[0], &op2[0])
}

// TODO mpq_mul_2exp

func Neg(rop *Rat, op *Rat) {
	C.mpq_neg(&rop[0], &op[0])
}
func Abs(rop *Rat, op *Rat) {
	C.mpq_abs(&rop[0], &op[0])
}

func Inv(rop *Rat, op *Rat) {
	C.mpq_inv(&rop[0], &op[0])
}

// 6.4 Comparison Functions
func Cmp(op1 *Rat, op2 *Rat) int {
	return int(C.mpq_cmp(&op1[0], &op2[0]))
}

func CmpZ(op1 *Rat, op2 *mpz.Int) int {
	// since mpz.Int is from a different package, it has a different signature in CGO and complains.
	// also the C vs. Go symatics of fixed size array are different so can't use Go-Style casts
	// have to use unsafe to force it to the right type.
	//
	// Hmm actually using mpz_ptr everywhere might be easier.. for another time.
	//
	var opint = C.mpz_ptr(unsafe.Pointer(&op2[0]))
	return int(C.mpq_cmp_z(&op1[0], opint))
}

func CmpUi(op1 *Rat, num2 uint, den2 uint) int {
	return int(C.macro_mpq_cmp_ui(&op1[0], C.ulong(num2), C.ulong(den2)))
}

func CmpSi(op1 *Rat, num2 int, den2 uint) int {
	return int(C.macro_mpq_cmp_si(&op1[0], C.long(num2), C.ulong(den2)))
}

func Sgn(op *Rat) int {
	return int(C.macro_mpq_sgn(&op[0]))
}

func Equal(op1 *Rat, op2 *Rat) int {
	return int(C.mpq_equal(&op1[0], &op2[0]))
}

//
// 6.5 Applying Integer Functions to Rationals
//

//
// 6.6 Input and Output Functions
//  (uses stdio FILE streams)
