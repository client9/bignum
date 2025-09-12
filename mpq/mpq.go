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

const mpz_ptr macro_mpq_numref(const mpq_t op) {
	// cast to mpz_ptr to avoid compiler warnings about const
	return (mpz_ptr)mpq_numref(op);
}

const mpz_ptr macro_mpq_denref(const mpq_t op) {
	// cast to mpz_ptr to avoid compiler warnings about const
	return (mpz_ptr) mpq_denref(op);
}

*/
import "C"

import (
	"unsafe"

	"github.com/client9/bignum/mpz"
)

//type RatPtr C.mpq_t

type RatPtr C.mpq_ptr
type Rat C.mpq_t

//
// 6.1 Initialization and Assignment Functions
//

func Init(x RatPtr) {
	C.mpq_init(x)
}

func Clear(x RatPtr) {
	C.mpq_clear(x)
}

func Set(rop RatPtr, op RatPtr) {
	C.mpq_set(rop, op)
}

// TODO mpq_set_z

func SetUi(rop RatPtr, op1 uint, op2 uint) {
	C.mpq_set_ui(rop, C.ulong(op1), C.ulong(op2))
}

func SetSi(rop RatPtr, op1 int, op2 uint) {
	C.mpq_set_si(rop, C.long(op1), C.ulong(op2))
}

func SetStr(rop RatPtr, s string, base int) int {
	cstr := C.CString(s)
	ret := C.mpq_set_str(rop, cstr, C.int(base))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

func Swap(rop1 RatPtr, rop2 RatPtr) {
	C.mpq_swap(rop1, rop2)
}

// 6.2 Conversion Functions

func GetD(rop RatPtr) float64 {
	return float64(C.mpq_get_d(rop))
}

func SetD(rop RatPtr, d float64) {
	C.mpq_set_d(rop, C.double( d))
}

// TODO mpq_set_f

// GetStr is not part of the GMP library.
func GetStr(base int, op RatPtr) string {
	p := C.mpq_get_str(nil, C.int(base), op)
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

func Add(rop RatPtr, op1 RatPtr, op2 RatPtr) {
	C.mpq_add(rop, op1, op2)
}
func Sub(rop RatPtr, op1 RatPtr, op2 RatPtr) {
	C.mpq_sub(rop, op1, op2)
}
func Mul(rop RatPtr, op1 RatPtr, op2 RatPtr) {
	C.mpq_mul(rop, op1, op2)
}

// TODO mpq_mul_2exp

func Div(rop RatPtr, op1 RatPtr, op2 RatPtr) {
	C.mpq_div(rop, op1, op2)
}

// TODO mpq_mul_2exp

func Neg(rop RatPtr, op RatPtr) {
	C.mpq_neg(rop, op)
}
func Abs(rop RatPtr, op RatPtr) {
	C.mpq_abs(rop, op)
}

func Inv(rop RatPtr, op RatPtr) {
	C.mpq_inv(rop, op)
}

//
// 6.4 Comparison Functions
//

func Cmp(op1 RatPtr, op2 RatPtr) int {
	return int(C.mpq_cmp(op1, op2))
}

func CmpZ(op1 RatPtr, op2 *mpz.Int) int {
	// since mpz.Int is from a different package, it has a different signature so CGO complains.
	// also the C vs. Go symatics of fixed size array are different so can't use Go-Style casts
	// have to use unsafe to force it to the right type.
	//
	// Hmm actually using mpz_ptr everywhere might be easier.. for another time.
	//
	var opint = C.mpz_ptr(unsafe.Pointer(op2))
	return int(C.mpq_cmp_z(op1, opint))
}

func CmpUi(op1 RatPtr, num2 uint, den2 uint) int {
	return int(C.macro_mpq_cmp_ui(op1, C.ulong(num2), C.ulong(den2)))
}

func CmpSi(op1 RatPtr, num2 int, den2 uint) int {
	return int(C.macro_mpq_cmp_si(op1, C.long(num2), C.ulong(den2)))
}

func Sgn(op RatPtr) int {
	return int(C.macro_mpq_sgn(op))
}

func Equal(op1 RatPtr, op2 RatPtr) int {
	return int(C.mpq_equal(op1, op2))
}

//
// 6.5 Applying Integer Functions to RatPtrionals
//
/*
func NumRef(op RatPtr) *mpz.Int {
	ptr := C.macro_mpq_numref(op)
	var x mpz.Int
	return x
}
func DenRef(op RatPtr) *Int {	
}
*/
//
// 6.6 Input and Output Functions
//  (uses stdio FILE streams)
