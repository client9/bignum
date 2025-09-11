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

//
// 5.1 Initialization Functions
//

func Init(x *Int) {
	C.mpz_init(&x[0])
}
func Init2(x *Int, n int) {
	C.mpz_init2(&x[0], C.mp_bitcnt_t(n))
}

func Clear(x *Int) {
	C.mpz_clear(&x[0])
}

func Realloc2(x *Int, n int) {
	C.mpz_realloc2(&x[0], C.mp_bitcnt_t(n))
}

// 5.2 Assignment Functions
func Set(rop *Int, op *Int) {
	C.mpz_set(&rop[0], &op[0])
}
func SetUi(rop *Int, op uint) {
	C.mpz_set_ui(&rop[0], C.ulong(op))
}

func SetSi(rop *Int, op int) {
	C.mpz_set_si(&rop[0], C.long(op))
}

func SetD(rop *Int, op float64) {
	C.mpz_set_d(&rop[0], C.double(op))
}

// TODO mpz_set_q --> depends on mpq

func SetStr(rop *Int, s string, base int) int {
	cstr := C.CString(s)
	ret := C.mpz_set_str(&rop[0], cstr, C.int(base))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

func Swap(rop1 *Int, rop2 *Int) {
	C.mpz_swap(&rop1[0], &rop2[0])
}

//
// 5.3 Combined Initization and Assignment Functions
//

func InitSet(rop *Int, op *Int) {
	C.mpz_init_set(&rop[0], &op[0])
}
func InitSetUi(rop *Int, op uint) {
	C.mpz_init_set_ui(&rop[0], C.ulong(op))
}

func InitSetSi(rop *Int, op int) {
	C.mpz_init_set_si(&rop[0], C.long(op))
}

func InitSetD(rop *Int, op float64) {
	C.mpz_init_set_d(&rop[0], C.double(op))
}
func InitSetStr(rop *Int, s string, base int) int {
	cstr := C.CString(s)
	ret := C.mpz_init_set_str(&rop[0], cstr, C.int(base))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

// 5.4 Conversion Functions
func GetUi(rop *Int) uint {
	return uint(C.mpz_get_ui(&rop[0]))
}

func GetSi(rop *Int) int {
	return int(C.mpz_get_si(&rop[0]))
}

func GetD(rop *Int) float64 {
	return float64(C.mpz_get_d(&rop[0]))
}

// TODO mpz_get_d_2exp

func GetStr(base int, op *Int) string {
	p := C.mpz_get_str(nil, C.int(base), &op[0])
	s := C.GoString(p)
	C.free(unsafe.Pointer(p))
	return s
}

//
// 5.5 Arithmetic Functions
//

func Add(rop *Int, op1 *Int, op2 *Int) {
	C.mpz_add(&rop[0], &op1[0], &op2[0])
}
func AddUi(rop *Int, op1 *Int, op2 uint) {
	C.mpz_add_ui(&rop[0], &op1[0], C.ulong(op2))
}
func Sub(rop *Int, op1 *Int, op2 *Int) {
	C.mpz_sub(&rop[0], &op1[0], &op2[0])
}
func SubUi(rop *Int, op1 *Int, op2 uint) {
	C.mpz_sub_ui(&rop[0], &op1[0], C.ulong(op2))
}
func UiSub(rop *Int, op1 uint, op2 *Int) {
	C.mpz_ui_sub(&rop[0], C.ulong(op1), &op2[0])
}
func Mul(rop *Int, op1 *Int, op2 *Int) {
	C.mpz_mul(&rop[0], &op1[0], &op2[0])
}
func MulSi(rop *Int, op1 *Int, op2 int) {
	C.mpz_mul_si(&rop[0], &op1[0], C.long(op2))
}
func MulUi(rop *Int, op1 *Int, op2 uint) {
	C.mpz_mul_ui(&rop[0], &op1[0], C.ulong(op2))
}
func AddMul(rop *Int, op1 *Int, op2 *Int) {
	C.mpz_addmul(&rop[0], &op1[0], &op2[0])
}
func AddmulUi(rop *Int, op1 *Int, op2 uint) {
	C.mpz_addmul_ui(&rop[0], &op1[0], C.ulong(op2))
}
func SubMul(rop *Int, op1 *Int, op2 *Int) {
	C.mpz_submul(&rop[0], &op1[0], &op2[0])
}
func SubmulUi(rop *Int, op1 *Int, op2 uint) {
	C.mpz_submul_ui(&rop[0], &op1[0], C.ulong(op2))
}

// TODO mpz_mul_2exp

func Neg(rop *Int, op *Int) {
	C.mpz_neg(&rop[0], &op[0])
}
func Abs(rop *Int, op *Int) {
	C.mpz_abs(&rop[0], &op[0])
}

//
// 5.6 Division Functions
//

func CdivQ(q *Int, n *Int, d *Int) {
	C.mpz_cdiv_q(&q[0], &n[0], &d[0])
}
func CdivR(r *Int, n *Int, d *Int) {
	C.mpz_cdiv_r(&r[0], &n[0], &d[0])
}
func CdivQr(q *Int, r *Int, n *Int, d *Int) {
	C.mpz_cdiv_qr(&q[0], &r[0], &n[0], &d[0])
}

func FdivQ(q *Int, n *Int, d *Int) {
	C.mpz_fdiv_q(&q[0], &n[0], &d[0])
}
func FdivR(r *Int, n *Int, d *Int) {
	C.mpz_fdiv_r(&r[0], &n[0], &d[0])
}
func FdivQr(q *Int, r *Int, n *Int, d *Int) {
	C.mpz_fdiv_qr(&q[0], &r[0], &n[0], &d[0])
}
func TdivQ(q *Int, n *Int, d *Int) {
	C.mpz_tdiv_q(&q[0], &n[0], &d[0])
}
func TdivR(r *Int, n *Int, d *Int) {
	C.mpz_tdiv_r(&r[0], &n[0], &d[0])
}
func TdivQr(q *Int, r *Int, n *Int, d *Int) {
	C.mpz_tdiv_qr(&q[0], &r[0], &n[0], &d[0])
}

// 5.7 Exponentiation Functions
func Powm(rop *Int, base *Int, exp *Int, mod *Int) {
	C.mpz_powm(&rop[0], &base[0], &exp[0], &mod[0])
}

func PowmUi(rop *Int, base *Int, exp uint, mod *Int) {
	C.mpz_powm_ui(&rop[0], &base[0], C.ulong(exp), &mod[0])
}
func PowmSec(rop *Int, base *Int, exp *Int, mod *Int) {
	C.mpz_powm_sec(&rop[0], &base[0], &exp[0], &mod[0])
}

func PowUi(rop *Int, base *Int, exp uint) {
	C.mpz_pow_ui(&rop[0], &base[0], C.ulong(exp))
}
func UiPowUi(rop *Int, base uint, exp uint) {
	C.mpz_ui_pow_ui(&rop[0], C.ulong(exp), C.ulong(exp))
}

//
// 5.8 Root Extraction Functions
//

func Sqrt(rop *Int, op *Int) {
	C.mpz_sqrt(&rop[0], &op[0])
}

// 5.9 Number Theoretic Functions
func ProbabPrimeP(n *Int, reps int) int {
	return int(C.mpz_probab_prime_p(&n[0], C.int(reps)))
}

func Gcd(rop *Int, op1 *Int, op2 *Int) {
	C.mpz_gcd(&rop[0], &op1[0], &op2[0])
}

// 5.10 Comparison Functions
func Cmp(op1 *Int, op2 *Int) int {
	return int(C.mpz_cmp(&op1[0], &op2[0]))
}
func CmpD(op1 *Int, op2 float64) int {
	return int(C.mpz_cmp_d(&op1[0], C.double(op2)))
}
func CmpSi(op1 *Int, op2 int) int {
	return int(C.macro_mpz_cmp_si(&op1[0], C.long(op2)))
}
func CmpUi(op1 *Int, op2 uint) int {
	return int(C.macro_mpz_cmp_ui(&op1[0], C.ulong(op2)))
}

func Cmpabs(op1 *Int, op2 *Int) int {
	return int(C.mpz_cmp(&op1[0], &op2[0]))
}

func CmpabsD(op1 *Int, op2 float64) int {
	return int(C.mpz_cmpabs_d(&op1[0], C.double(op2)))
}

func CmpabsUi(op1 *Int, op2 uint) int {
	return int(C.mpz_cmpabs_ui(&op1[0], C.ulong(op2)))
}
func Sgn(op *Int) int {
	return int(C.macro_mpz_sgn(&op[0]))
}

//
// 5.11 Logical and Bit Manipulation Functions
//

//
// 5.12 Input and Output Functions
//

//
// 5.13 Input and Output Functions
//

//
// 5.14 Integer Import and Export
//

// 5.15 Miscellaneous Functions
func Sizeinbase(op *Int, base int) int {
	return int(C.mpz_sizeinbase(&op[0], C.int(base)))
}

//
// 5.16 Special Functions
//

func Size(op *Int) int {
	return int(C.mpz_size(&op[0]))
}
