package mpz

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lgmp

#include <stdlib.h>
#include <stdio.h>

#include "gmp.h"
int macro_mpz_cmp_si(const mpz_t op, signed long int op2) {
	return mpz_cmp_si(op, op2);
}
int macro_mpz_cmp_ui(const mpz_t op, unsigned long int op2) {
	return mpz_cmp_ui(op, op2);
}
int macro_mpz_sgn(const mpz_t op) {
        return mpz_sgn(op);
}
*/
import "C"

import (
	"unsafe"
)

type Int C.mpz_t
type IntPtr C.mpz_ptr

//
// 5.1 Initialization Functions
//

func Init(x IntPtr) {
	C.mpz_init(x)
}
func Init2(x IntPtr, n int) {
	C.mpz_init2(x, C.mp_bitcnt_t(n))
}

func Clear(x IntPtr) {
	C.mpz_clear(x)
}

func Realloc2(x IntPtr, n int) {
	C.mpz_realloc2(x, C.mp_bitcnt_t(n))
}

// 5.2 Assignment Functions
func Set(rop IntPtr, op IntPtr) {
	C.mpz_set(rop, op)
}
func SetUi(rop IntPtr, op uint) {
	C.mpz_set_ui(rop, C.ulong(op))
}

func SetSi(rop IntPtr, op int) {
	C.mpz_set_si(rop, C.long(op))
}

func SetD(rop IntPtr, op float64) {
	C.mpz_set_d(rop, C.double(op))
}

// TODO mpz_set_q --> depends on mpq

func SetStr(rop IntPtr, s string, base int) int {
	cstr := C.CString(s)
	ret := C.mpz_set_str(rop, cstr, C.int(base))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

func Swap(rop1 IntPtr, rop2 IntPtr) {
	C.mpz_swap(rop1, rop2)
}

//
// 5.3 Combined Initization and Assignment Functions
//

func InitSet(rop IntPtr, op IntPtr) {
	C.mpz_init_set(rop, op)
}
func InitSetUi(rop IntPtr, op uint) {
	C.mpz_init_set_ui(rop, C.ulong(op))
}

func InitSetSi(rop IntPtr, op int) {
	C.mpz_init_set_si(rop, C.long(op))
}

func InitSetD(rop IntPtr, op float64) {
	C.mpz_init_set_d(rop, C.double(op))
}
func InitSetStr(rop IntPtr, s string, base int) int {
	cstr := C.CString(s)
	ret := C.mpz_init_set_str(rop, cstr, C.int(base))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

// 5.4 Conversion Functions
func GetUi(rop IntPtr) uint {
	return uint(C.mpz_get_ui(rop))
}

func GetSi(rop IntPtr) int {
	return int(C.mpz_get_si(rop))
}

func GetD(rop IntPtr) float64 {
	return float64(C.mpz_get_d(rop))
}

// TODO mpz_get_d_2exp

func GetStr(base int, op IntPtr) string {
	p := C.mpz_get_str(nil, C.int(base), op)
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

//
// 5.5 Arithmetic Functions
//

func Add(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_add(rop, op1, op2)
}
func AddUi(rop IntPtr, op1 IntPtr, op2 uint) {
	C.mpz_add_ui(rop, op1, C.ulong(op2))
}
func Sub(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_sub(rop, op1, op2)
}
func SubUi(rop IntPtr, op1 IntPtr, op2 uint) {
	C.mpz_sub_ui(rop, op1, C.ulong(op2))
}
func UiSub(rop IntPtr, op1 uint, op2 IntPtr) {
	C.mpz_ui_sub(rop, C.ulong(op1), op2)
}
func Mul(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_mul(rop, op1, op2)
}
func MulSi(rop IntPtr, op1 IntPtr, op2 int) {
	C.mpz_mul_si(rop, op1, C.long(op2))
}
func MulUi(rop IntPtr, op1 IntPtr, op2 uint) {
	C.mpz_mul_ui(rop, op1, C.ulong(op2))
}
func AddMul(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_addmul(rop, op1, op2)
}
func AddmulUi(rop IntPtr, op1 IntPtr, op2 uint) {
	C.mpz_addmul_ui(rop, op1, C.ulong(op2))
}
func SubMul(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_submul(rop, op1, op2)
}
func SubmulUi(rop IntPtr, op1 IntPtr, op2 uint) {
	C.mpz_submul_ui(rop, op1, C.ulong(op2))
}

// TODO mpz_mul_2exp

func Neg(rop IntPtr, op IntPtr) {
	C.mpz_neg(rop, op)
}
func Abs(rop IntPtr, op IntPtr) {
	C.mpz_abs(rop, op)
}

//
// 5.6 Division Functions
//

func CdivQ(q IntPtr, n IntPtr, d IntPtr) {
	C.mpz_cdiv_q(q, n, d)
}
func CdivR(r IntPtr, n IntPtr, d IntPtr) {
	C.mpz_cdiv_r(r, n, d)
}
func CdivQr(q IntPtr, r IntPtr, n IntPtr, d IntPtr) {
	C.mpz_cdiv_qr(q, r, n, d)
}

func FdivQ(q IntPtr, n IntPtr, d IntPtr) {
	C.mpz_fdiv_q(q, n, d)
}
func FdivR(r IntPtr, n IntPtr, d IntPtr) {
	C.mpz_fdiv_r(r, n, d)
}
func FdivQr(q IntPtr, r IntPtr, n IntPtr, d IntPtr) {
	C.mpz_fdiv_qr(q, r, n, d)
}
func TdivQ(q IntPtr, n IntPtr, d IntPtr) {
	C.mpz_tdiv_q(q, n, d)
}
func TdivR(r IntPtr, n IntPtr, d IntPtr) {
	C.mpz_tdiv_r(r, n, d)
}
func TdivQr(q IntPtr, r IntPtr, n IntPtr, d IntPtr) {
	C.mpz_tdiv_qr(q, r, n, d)
}
func Mod(r IntPtr, n IntPtr, d IntPtr) {
	C.mpz_mod(r, n, d)
}

// 5.7 Exponentiation Functions
func Powm(rop IntPtr, base IntPtr, exp IntPtr, mod IntPtr) {
	C.mpz_powm(rop, base, exp, mod)
}

func PowmUi(rop IntPtr, base IntPtr, exp uint, mod IntPtr) {
	C.mpz_powm_ui(rop, base, C.ulong(exp), mod)
}
func PowmSec(rop IntPtr, base IntPtr, exp IntPtr, mod IntPtr) {
	C.mpz_powm_sec(rop, base, exp, mod)
}

func PowUi(rop IntPtr, base IntPtr, exp uint) {
	C.mpz_pow_ui(rop, base, C.ulong(exp))
}
func UiPowUi(rop IntPtr, base uint, exp uint) {
	C.mpz_ui_pow_ui(rop, C.ulong(exp), C.ulong(exp))
}

//
// 5.8 Root Extraction Functions
//

func Sqrt(rop IntPtr, op IntPtr) {
	C.mpz_sqrt(rop, op)
}

// 5.9 Number Theoretic Functions
func ProbabPrimeP(n IntPtr, reps int) int {
	return int(C.mpz_probab_prime_p(n, C.int(reps)))
}

func Gcd(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_gcd(rop, op1, op2)
}

func Invert(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_invert(rop, op1, op2)
}

// 5.10 Comparison Functions
func Cmp(op1 IntPtr, op2 IntPtr) int {
	return int(C.mpz_cmp(op1, op2))
}
func CmpD(op1 IntPtr, op2 float64) int {
	return int(C.mpz_cmp_d(op1, C.double(op2)))
}
func CmpSi(op1 IntPtr, op2 int) int {
	return int(C.macro_mpz_cmp_si(op1, C.long(op2)))
}
func CmpUi(op1 IntPtr, op2 uint) int {
	return int(C.macro_mpz_cmp_ui(op1, C.ulong(op2)))
}

func Cmpabs(op1 IntPtr, op2 IntPtr) int {
	return int(C.mpz_cmp(op1, op2))
}

func CmpabsD(op1 IntPtr, op2 float64) int {
	return int(C.mpz_cmpabs_d(op1, C.double(op2)))
}

func CmpabsUi(op1 IntPtr, op2 uint) int {
	return int(C.mpz_cmpabs_ui(op1, C.ulong(op2)))
}
func Sgn(op IntPtr) int {
	return int(C.macro_mpz_sgn(op))
}

//
// 5.11 Logical and Bit Manipulation Functions
//

func And(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_and(rop, op1, op2)
}

func Ior(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_ior(rop, op1, op2)
}

func Xor(rop IntPtr, op1 IntPtr, op2 IntPtr) {
	C.mpz_xor(rop, op1, op2)
}

func Com(rop IntPtr, op IntPtr) {
	C.mpz_com(rop, op)
}

//
// 5.12 Input and Output Functions
//

//
// 5.13 Input and Output Functions
//

//
// 5.14 IntPtreger Import and Export
//

// 5.15 Miscellaneous Functions
func FitsUlongP(op IntPtr) int {
	return int(C.mpz_fits_ulong_p(op))
}

func FitsSlongP(op IntPtr) int {
	return int(C.mpz_fits_slong_p(op))
}

func Sizeinbase(op IntPtr, base int) int {
	return int(C.mpz_sizeinbase(op, C.int(base)))
}

//
// 5.16 Special Functions
//

func Size(op IntPtr) int {
	return int(C.mpz_size(op))
}
