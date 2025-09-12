package mpfr

/*
#cgo CFLAGS: -I/opt/homebrew/include
#cgo LDFLAGS: -L/opt/homebrew/lib -lmpfr -lgmp

#include <stdlib.h>
#include <stdio.h> // For printf in C

#include "mpfr.h"

int fn_mpfr_init_set(mpfr_t rop, mpfr_t op, mpfr_rnd_t rnd) {
	return mpfr_init_set(rop, op, rnd);
}
int fn_mpfr_init_set_ui(mpfr_t rop, unsigned long int op, mpfr_rnd_t rnd) {
	return mpfr_init_set_ui(rop, op, rnd);
}
int fn_mpfr_init_set_si(mpfr_t rop, long int op, mpfr_rnd_t rnd) {
	return mpfr_init_set_si(rop, op, rnd);
}
int fn_mpfr_init_set_d (mpfr_t rop, double op, mpfr_rnd_t rnd) {
	return mpfr_init_set_d(rop, op, rnd);
}

// mpfr_print is va-arg function which can't be called by cgo.
// Wrap basic case of a template string for a mpfr_t
int mpfr_printf2(const char* template, mpfr_t x) {
	return mpfr_printf(template, x);
}
// mpfr_print is va-arg function which can't be called by cgo.
// return value must be free with mpfr_free_str
char* mpfr_sprintf3(const char* template, mpfr_rnd_t rnd, mpfr_t x)  {
	char* buf;
	int len;

	len = mpfr_asprintf(&buf, template, rnd, x);
	if (len < 0) {
		return NULL;
	}
	return buf;
}
// mpfr_print is va-arg function which can't be called by cgo.
// return value must be free with mpfr_free_str
char* mpfr_sprintf2(const char* template, mpfr_t x)  {
	char* buf;
	int len;

	len = mpfr_asprintf(&buf, template, x);
	if (len < 0) {
		return NULL;
	}
	return buf;
}
*/
import "C"

import (
	"unsafe"
)

type RoundMode int

const (
	RNDN RoundMode = C.MPFR_RNDN
	RNDD RoundMode = C.MPFR_RNDD
	RNDU RoundMode = C.MPFR_RNDU
	RNDZ RoundMode = C.MPFR_RNDZ
	RNDA RoundMode = C.MPFR_RNDA
	RNDF RoundMode = C.MPFR_RNDF
)

type Float C.mpfr_t

//
// 5.1 Initialization Functions
//

func Init(z *Float) {
	C.mpfr_init(&z[0])
}
func Init2(z *Float, prec int) {
	C.mpfr_init2(&z[0], C.mpfr_prec_t(prec))
}

func Clear(z *Float) {
	C.mpfr_clear(&z[0])
}

func SetDefaultPrec(prec int) {
	C.mpfr_set_default_prec(C.mpfr_prec_t(prec))
}

func GetDefaultPrec() int {
	return int(C.mpfr_get_default_prec())
}

func SetPrec(z *Float, prec int) {
	C.mpfr_set_prec(&z[0], C.mpfr_prec_t(prec))
}

func GetPrec(z *Float) int {
	return int(C.mpfr_get_prec(&z[0]))
}

//
// 5.2 Assignment Functions
//

func Set(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_set(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func SetUi(z *Float, op uint64, rnd RoundMode) int {
	return int(C.mpfr_set_ui(&z[0], C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func SetSi(z *Float, op int64, rnd RoundMode) int {
	return int(C.mpfr_set_si(&z[0], C.long(op), C.mpfr_rnd_t(rnd)))
}
func SetFlt(z *Float, val float32, rnd RoundMode) int {
	return int(C.mpfr_set_flt(&z[0], C.float(val), C.mpfr_rnd_t(rnd)))
}
func SetD(z *Float, val float64, rnd RoundMode) int {
	return int(C.mpfr_set_d(&z[0], C.double(val), C.mpfr_rnd_t(rnd)))
}

// TODO OTHER VARIATIONS

func SetNan(x *Float) {
	C.mpfr_set_nan(&x[0])
}

func SetInf(x *Float, sign int) {
	C.mpfr_set_inf(&x[0], C.int(sign))
}

func SetZero(x *Float, sign int) {
	C.mpfr_set_zero(&x[0], C.int(sign))
}

func Swap(x *Float, y *Float) {
	C.mpfr_swap(&x[0], &y[0])
}

//
// 5.3 Combined Initialization and Assignment Functions
//

func InitSet(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func InitSetUi(z *Float, op uint64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_ui(&z[0], C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func InitSetSi(z *Float, op int64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_si(&z[0], C.long(op), C.mpfr_rnd_t(rnd)))
}
func InitSetD(z *Float, val float64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_d(&z[0], C.double(val), C.mpfr_rnd_t(rnd)))
}

// TODO: mpfr_set_ld "long double" is probably a 64-bit integer, but in some
// place it might be 80bit.

// TODO: mpfr_set_z, set_q, set_f

func InitSetStr(x *Float, s string, base int, rnd RoundMode) int {
	//cstr := unsafe.Pointer(C.CString(s))
	cstr := C.CString(s)
	ret := C.mpfr_set_str(&x[0], cstr, C.int(base), C.mpfr_rnd_t(rnd))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

// 5.4 Conversion Functions
func GetFlt(op *Float, rnd RoundMode) float32 {
	return float32(C.mpfr_get_flt(&op[0], C.mpfr_rnd_t(rnd)))
}
func GetD(op *Float, rnd RoundMode) float64 {
	return float64(C.mpfr_get_d(&op[0], C.mpfr_rnd_t(rnd)))
}

func GetSi(op *Float, rnd RoundMode) int {
	return int(C.mpfr_get_si(&op[0], C.mpfr_rnd_t(rnd)))
}
func GetUi(op *Float, rnd RoundMode) int {
	return int(C.mpfr_get_ui(&op[0], C.mpfr_rnd_t(rnd)))
}

func GetStrNdigits(b int, prec int) int {
	return int(C.mpfr_get_str_ndigits(C.int(b), C.mpfr_prec_t(prec)))
}

// char *str, mpfr_exp_t *expptr, int base, size_t n, mpfr_t op, mpfr_rnd_t rnd)
func GetStr(base int, n int, op *Float, rnd RoundMode) (string, int) {
	var exp int
	cIntPtr := (*C.mpfr_exp_t)(unsafe.Pointer(&exp))
	p := C.mpfr_get_str(nil, cIntPtr, C.int(base), C.size_t(n), &op[0], C.mpfr_rnd_t(rnd))
	s := C.GoString(p)
	C.mpfr_free_str(p)
	return s, exp
}

// 5.5 Arithmetic Functions
func Add(rop *Float, op1 *Float, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_add(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}
func AddUi(rop *Float, op1 *Float, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_add_ui(&rop[0], &op1[0], C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func AddSi(rop *Float, op1 *Float, op2 int, rnd RoundMode) int {
	return int(C.mpfr_add_si(&rop[0], &op1[0], C.long(op2), C.mpfr_rnd_t(rnd)))
}
func AddD(rop *Float, op1 *Float, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_add_d(&rop[0], &op1[0], C.double(op2), C.mpfr_rnd_t(rnd)))
}

func Sub(rop *Float, op1 *Float, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_sub(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}
func SubUi(rop *Float, op1 *Float, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_sub_ui(&rop[0], &op1[0], C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func SubSi(rop *Float, op1 *Float, op2 int, rnd RoundMode) int {
	return int(C.mpfr_sub_si(&rop[0], &op1[0], C.long(op2), C.mpfr_rnd_t(rnd)))
}
func SubD(rop *Float, op1 *Float, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_sub_d(&rop[0], &op1[0], C.double(op2), C.mpfr_rnd_t(rnd)))
}

func Mul(rop *Float, op1 *Float, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_mul(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}

func MulUi(rop *Float, op1 *Float, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_mul_ui(&rop[0], &op1[0], C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func MulSi(rop *Float, op1 *Float, op2 int, rnd RoundMode) int {
	return int(C.mpfr_mul_si(&rop[0], &op1[0], C.long(op2), C.mpfr_rnd_t(rnd)))
}
func MulD(rop *Float, op1 *Float, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_mul_d(&rop[0], &op1[0], C.double(op2), C.mpfr_rnd_t(rnd)))
}
func Sqr(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_sqr(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Div(rop *Float, op1 *Float, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_div(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}

func DivUi(rop *Float, op1 *Float, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_div_ui(&rop[0], &op1[0], C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func DivSi(rop *Float, op1 *Float, op2 int, rnd RoundMode) int {
	return int(C.mpfr_div_si(&rop[0], &op1[0], C.long(op2), C.mpfr_rnd_t(rnd)))
}
func DivD(rop *Float, op1 *Float, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_div_d(&rop[0], &op1[0], C.double(op2), C.mpfr_rnd_t(rnd)))
}
func UiDiv(rop *Float, op1 uint, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_ui_div(&rop[0], C.ulong(op1), &op2[0], C.mpfr_rnd_t(rnd)))
}
func SiDiv(rop *Float, op1 int, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_si_div(&rop[0], C.long(op1), &op2[0], C.mpfr_rnd_t(rnd)))
}
func DDiv(rop *Float, op1 float64, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_d_div(&rop[0], C.double(op1), &op2[0], C.mpfr_rnd_t(rnd)))
}
func Sqrt(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_sqrt(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func SqrtUi(rop *Float, op uint, rnd RoundMode) int {
	return int(C.mpfr_sqrt_ui(&rop[0], C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func RecSqrt(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_rec_sqrt(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Cbrt(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_cbrt(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RootnUi(rop *Float, op *Float, n uint, rnd RoundMode) int {
	return int(C.mpfr_rootn_ui(&rop[0], &op[0], C.ulong(n), C.mpfr_rnd_t(rnd)))
}
func RootnSi(rop *Float, op *Float, n int, rnd RoundMode) int {
	return int(C.mpfr_rootn_si(&rop[0], &op[0], C.long(n), C.mpfr_rnd_t(rnd)))
}

/* DEPRECATED
func Root(rop *Float, op *Float, n uint, rnd RoundMode) int {
	return int(C.mpfr_rootn_si(&rop[0], &op[0], C.ulong(n), C.mpfr_rnd_t(rnd)))
}
*/

func Neg(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_neg(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Abs(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_abs(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

// TOOD other functions

//
// 5.6 Comparison Functions
//

func Cmp(op1 *Float, op2 *Float) int {
	return int(C.mpfr_cmp(&op1[0], &op2[0]))
}
func Cmpabs(op1 *Float, op2 *Float) int {
	return int(C.mpfr_cmpabs(&op1[0], &op2[0]))
}
func NanP(op *Float) int {
	return int(C.mpfr_nan_p(&op[0]))
}
func InfP(op *Float) int {
	return int(C.mpfr_inf_p(&op[0]))
}
func NumberP(op *Float) int {
	return int(C.mpfr_number_p(&op[0]))
}
func ZeroP(op *Float) int {
	return int(C.mpfr_zero_p(&op[0]))
}
func RegularP(op *Float) int {
	return int(C.mpfr_regular_p(&op[0]))
}
func Sgn(op *Float) int {
	return int(C.mpfr_sgn(&op[0]))
}
func GreaterP(op1 *Float, op2 *Float) int {
	return int(C.mpfr_greater_p(&op1[0], &op2[0]))
}
func GreaterequalP(op1 *Float, op2 *Float) int {
	return int(C.mpfr_greaterequal_p(&op1[0], &op2[0]))
}
func LessP(op1 *Float, op2 *Float) int {
	return int(C.mpfr_less_p(&op1[0], &op2[0]))
}
func LessequalP(op1 *Float, op2 *Float) int {
	return int(C.mpfr_lessequal_p(&op1[0], &op2[0]))
}
func EqualP(op1 *Float, op2 *Float) int {
	return int(C.mpfr_equal_p(&op1[0], &op2[0]))
}
func LessgreaterP(op1 *Float, op2 *Float) int {
	return int(C.mpfr_lessgreater_p(&op1[0], &op2[0]))
}
func UnorderedP(op1 *Float, op2 *Float) int {
	return int(C.mpfr_unordered_p(&op1[0], &op2[0]))
}
func TotalOrderP(op1 *Float, op2 *Float) int {
	return int(C.mpfr_total_order_p(&op1[0], &op2[0]))
}

//
// 5.7 Transcendental Functions
//

func Log(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_log(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func LogUi(rop *Float, op uint64, rnd RoundMode) int {
	return int(C.mpfr_log_ui(&rop[0], C.ulong(op), C.mpfr_rnd_t(rnd)))
}

func Log2(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_log2(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Log10(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_log10(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Log1p(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_log1p(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Log2p1(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_log2p1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Log10p1(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_log10p1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_exp(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp2(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_exp2(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp10(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_exp10(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Expm1(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_expm1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp2m1(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_exp2m1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp10m1(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_exp10m1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

//
// Powers
//

func Pow(rop *Float, op1 *Float, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_pow(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}

func Powr(rop *Float, op1 *Float, op2 *Float, rnd RoundMode) int {
	return int(C.mpfr_powr(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}

// Trigonometry
func Cos(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_cos(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Sin(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_sin(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Tan(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_tan(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Cospi(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_cospi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Sinpi(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_sinpi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Tanpi(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_tanpi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Sec(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_sec(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Csc(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_csc(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Cot(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_cot(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Acos(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_acos(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Asin(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_asin(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Atan(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_atan(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Acospi(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_acospi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Asinpi(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_asinpi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Atanpi(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_atanpi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

// Special constants
func ConstLog2(rop *Float, rnd RoundMode) int {
	return int(C.mpfr_const_log2(&rop[0], C.mpfr_rnd_t(rnd)))
}
func ConstPi(rop *Float, rnd RoundMode) int {
	return int(C.mpfr_const_pi(&rop[0], C.mpfr_rnd_t(rnd)))
}
func ConstEuler(rop *Float, rnd RoundMode) int {
	return int(C.mpfr_const_euler(&rop[0], C.mpfr_rnd_t(rnd)))
}
func ConstCatalan(rop *Float, rnd RoundMode) int {
	return int(C.mpfr_const_catalan(&rop[0], C.mpfr_rnd_t(rnd)))
}

//
// 5.8 Input and Output Functinos
//

func Dump(x *Float) {
	C.mpfr_dump(&x[0])
}

//
// 5.9 Formatted Output Functions
//

func Printf2(template string, x *Float) int {
	return int(C.mpfr_printf2(C.CString(template), &x[0]))
}

func Sprintf3(template string, rnd RoundMode, x *Float) string {
	p := C.mpfr_sprintf3(C.CString(template), C.mpfr_rnd_t(rnd), &x[0])
	if p == nil {
		return ""
	}
	s := C.GoString(p)
	C.mpfr_free_str(p)
	return s
}

func Sprintf2(template string, x *Float) string {
	p := C.mpfr_sprintf2(C.CString(template), &x[0])
	if p == nil {
		return ""
	}
	s := C.GoString(p)
	C.mpfr_free_str(p)
	return s
}

//
// 5.10 Integer and Remainder Related Functions
//

func Rint(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_rint(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Ceil(rop *Float, op *Float) int {
	return int(C.mpfr_ceil(&rop[0], &op[0]))
}
func Floor(rop *Float, op *Float) int {
	return int(C.mpfr_floor(&rop[0], &op[0]))
}
func Round(rop *Float, op *Float) int {
	return int(C.mpfr_round(&rop[0], &op[0]))
}
func Roundeven(rop *Float, op *Float) int {
	return int(C.mpfr_roundeven(&rop[0], &op[0]))
}
func Trunc(rop *Float, op *Float) int {
	return int(C.mpfr_trunc(&rop[0], &op[0]))
}

func RintCeil(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_rint_ceil(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RintFloor(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_rint_floor(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RintRound(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_rint_round(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RintRoundeven(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_rint_roundeven(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RintTrunc(rop *Float, op *Float, rnd RoundMode) int {
	return int(C.mpfr_rint_trunc(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

// 5.11 Rounding-Related Functions
func MinPrec(x *Float) uint {
	return uint(mpfr_min_prec(&x.ptr))
}

//
// 5.12 Miscellaneous Functions
//

func Signbit(op *Float) int {
	return int(C.mpfr_signbit(&op[0]))
}

func GetVersion() string {
	p := C.mpfr_get_version()
	s := C.GoString(p)
	return s
}

func GetPatches() string {
	p := C.mpfr_get_patches()
	s := C.GoString(p)
	return s
}

//
// 5.13 Exception Related Functions
//

//
// 5.14 Memory Handling Functions
//

//
// 5.15 Compatibility With MPF
//

//
// 5.16 Custom Interface
//
