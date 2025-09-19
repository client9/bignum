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
int fn_mpfr_init_set_z (mpfr_t rop, mpz_t op, mpfr_rnd_t rnd) {
	return mpfr_init_set_z(rop, op, rnd);
}
int fn_mpfr_init_set_q (mpfr_t rop, mpq_t op, mpfr_rnd_t rnd) {
	return mpfr_init_set_q(rop, op, rnd);
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
	"github.com/client9/bignum/mpq"
	"github.com/client9/bignum/mpz"
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
type FloatPtr C.mpfr_ptr

//
// 5.1 Initialization Functions
//

func Init(z FloatPtr) {
	C.mpfr_init(z)
}
func Init2(z FloatPtr, prec int) {
	C.mpfr_init2(z, C.mpfr_prec_t(prec))
}

func Clear(z FloatPtr) {
	C.mpfr_clear(z)
}

func SetDefaultPrec(prec int) {
	C.mpfr_set_default_prec(C.mpfr_prec_t(prec))
}

func GetDefaultPrec() int {
	return int(C.mpfr_get_default_prec())
}

func SetPrec(z FloatPtr, prec int) {
	C.mpfr_set_prec(z, C.mpfr_prec_t(prec))
}

func GetPrec(z FloatPtr) int {
	return int(C.mpfr_get_prec(z))
}

//
// 5.2 Assignment Functions
//

func Set(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_set(rop, op, C.mpfr_rnd_t(rnd)))
}
func SetUi(z FloatPtr, op uint64, rnd RoundMode) int {
	return int(C.mpfr_set_ui(z, C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func SetSi(z FloatPtr, op int64, rnd RoundMode) int {
	return int(C.mpfr_set_si(z, C.long(op), C.mpfr_rnd_t(rnd)))
}
func SetFlt(z FloatPtr, val float32, rnd RoundMode) int {
	return int(C.mpfr_set_flt(z, C.float(val), C.mpfr_rnd_t(rnd)))
}
func SetD(z FloatPtr, val float64, rnd RoundMode) int {
	return int(C.mpfr_set_d(z, C.double(val), C.mpfr_rnd_t(rnd)))
}

func SetZ(z FloatPtr, op mpz.IntPtr, rnd RoundMode) int {
	return int(C.mpfr_set_z(z, C.mpz_ptr(unsafe.Pointer(op)), C.mpfr_rnd_t(rnd)))
}

func SetQ(z FloatPtr, op mpq.RatPtr, rnd RoundMode) int {
	return int(C.mpfr_set_q(z, C.mpq_ptr(unsafe.Pointer(op)), C.mpfr_rnd_t(rnd)))
}

func SetStr(x FloatPtr, s string, base int, rnd RoundMode) int {
	//cstr := unsafe.Pointer(C.CString(s))
	cstr := C.CString(s)
	ret := C.mpfr_set_str(x, cstr, C.int(base), C.mpfr_rnd_t(rnd))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

// TODO OTHER VARIATIONS

func SetNan(x FloatPtr) {
	C.mpfr_set_nan(x)
}

func SetInf(x FloatPtr, sign int) {
	C.mpfr_set_inf(x, C.int(sign))
}

func SetZero(x FloatPtr, sign int) {
	C.mpfr_set_zero(x, C.int(sign))
}

func Swap(x FloatPtr, y FloatPtr) {
	C.mpfr_swap(x, y)
}

//
// 5.3 Combined Initialization and Assignment Functions
//

func InitSet(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set(rop, op, C.mpfr_rnd_t(rnd)))
}
func InitSetUi(z FloatPtr, op uint64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_ui(z, C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func InitSetSi(z FloatPtr, op int64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_si(z, C.long(op), C.mpfr_rnd_t(rnd)))
}
func InitSetD(z FloatPtr, val float64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_d(z, C.double(val), C.mpfr_rnd_t(rnd)))
}
func InitSetZ(z FloatPtr, val mpz.IntPtr, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_z(z, C.mpz_ptr(unsafe.Pointer(val)), C.mpfr_rnd_t(rnd)))
}
func InitSetQ(z FloatPtr, val mpq.RatPtr, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_q(z, C.mpq_ptr(unsafe.Pointer(val)), C.mpfr_rnd_t(rnd)))
}

// TODO: mpfr_set_ld "long double" is probably a 64-bit integer, but in some
// place it might be 80bit.

// MACRO:
// TODO: mpfr_init_set_z, init_set_q, set_f

func InitSetStr(x FloatPtr, s string, base int, rnd RoundMode) int {
	//cstr := unsafe.Pointer(C.CString(s))
	cstr := C.CString(s)
	ret := C.mpfr_init_set_str(x, cstr, C.int(base), C.mpfr_rnd_t(rnd))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

// 5.4 Conversion Functions
func GetFlt(op FloatPtr, rnd RoundMode) float32 {
	return float32(C.mpfr_get_flt(op, C.mpfr_rnd_t(rnd)))
}
func GetD(op FloatPtr, rnd RoundMode) float64 {
	return float64(C.mpfr_get_d(op, C.mpfr_rnd_t(rnd)))
}

func GetSi(op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_get_si(op, C.mpfr_rnd_t(rnd)))
}
func GetUi(op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_get_ui(op, C.mpfr_rnd_t(rnd)))
}

func GetZ(rop mpz.IntPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_get_z(C.mpz_ptr(unsafe.Pointer(rop)), op, C.mpfr_rnd_t(rnd)))
}

func GetStrNdigits(b int, prec int) int {
	return int(C.mpfr_get_str_ndigits(C.int(b), C.mpfr_prec_t(prec)))
}

// char *str, mpfr_exp_t *expptr, int base, size_t n, mpfr_t op, mpfr_rnd_t rnd)
func GetStr(base int, n int, op FloatPtr, rnd RoundMode) (string, int) {
	var exp int
	cIntPtr := (*C.mpfr_exp_t)(unsafe.Pointer(&exp))
	p := C.mpfr_get_str(nil, cIntPtr, C.int(base), C.size_t(n), op, C.mpfr_rnd_t(rnd))
	s := C.GoString(p)
	C.mpfr_free_str(p)
	return s, exp
}

// 5.5 Arithmetic Functions
func Add(rop FloatPtr, op1 FloatPtr, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_add(rop, op1, op2, C.mpfr_rnd_t(rnd)))
}
func AddUi(rop FloatPtr, op1 FloatPtr, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_add_ui(rop, op1, C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func AddSi(rop FloatPtr, op1 FloatPtr, op2 int, rnd RoundMode) int {
	return int(C.mpfr_add_si(rop, op1, C.long(op2), C.mpfr_rnd_t(rnd)))
}
func AddD(rop FloatPtr, op1 FloatPtr, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_add_d(rop, op1, C.double(op2), C.mpfr_rnd_t(rnd)))
}

func Sub(rop FloatPtr, op1 FloatPtr, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sub(rop, op1, op2, C.mpfr_rnd_t(rnd)))
}
func SubUi(rop FloatPtr, op1 FloatPtr, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_sub_ui(rop, op1, C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func SubSi(rop FloatPtr, op1 FloatPtr, op2 int, rnd RoundMode) int {
	return int(C.mpfr_sub_si(rop, op1, C.long(op2), C.mpfr_rnd_t(rnd)))
}
func SubD(rop FloatPtr, op1 FloatPtr, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_sub_d(rop, op1, C.double(op2), C.mpfr_rnd_t(rnd)))
}

func Mul(rop FloatPtr, op1 FloatPtr, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_mul(rop, op1, op2, C.mpfr_rnd_t(rnd)))
}

func MulUi(rop FloatPtr, op1 FloatPtr, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_mul_ui(rop, op1, C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func MulSi(rop FloatPtr, op1 FloatPtr, op2 int, rnd RoundMode) int {
	return int(C.mpfr_mul_si(rop, op1, C.long(op2), C.mpfr_rnd_t(rnd)))
}
func MulD(rop FloatPtr, op1 FloatPtr, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_mul_d(rop, op1, C.double(op2), C.mpfr_rnd_t(rnd)))
}
func Sqr(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sqr(rop, op, C.mpfr_rnd_t(rnd)))
}

func Div(rop FloatPtr, op1 FloatPtr, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_div(rop, op1, op2, C.mpfr_rnd_t(rnd)))
}

func DivUi(rop FloatPtr, op1 FloatPtr, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_div_ui(rop, op1, C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func DivSi(rop FloatPtr, op1 FloatPtr, op2 int, rnd RoundMode) int {
	return int(C.mpfr_div_si(rop, op1, C.long(op2), C.mpfr_rnd_t(rnd)))
}
func DivD(rop FloatPtr, op1 FloatPtr, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_div_d(rop, op1, C.double(op2), C.mpfr_rnd_t(rnd)))
}
func UiDiv(rop FloatPtr, op1 uint, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_ui_div(rop, C.ulong(op1), op2, C.mpfr_rnd_t(rnd)))
}
func SiDiv(rop FloatPtr, op1 int, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_si_div(rop, C.long(op1), op2, C.mpfr_rnd_t(rnd)))
}
func DDiv(rop FloatPtr, op1 float64, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_d_div(rop, C.double(op1), op2, C.mpfr_rnd_t(rnd)))
}
func Sqrt(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sqrt(rop, op, C.mpfr_rnd_t(rnd)))
}
func SqrtUi(rop FloatPtr, op uint, rnd RoundMode) int {
	return int(C.mpfr_sqrt_ui(rop, C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func RecSqrt(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rec_sqrt(rop, op, C.mpfr_rnd_t(rnd)))
}
func Cbrt(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_cbrt(rop, op, C.mpfr_rnd_t(rnd)))
}
func RootnUi(rop FloatPtr, op FloatPtr, n uint, rnd RoundMode) int {
	return int(C.mpfr_rootn_ui(rop, op, C.ulong(n), C.mpfr_rnd_t(rnd)))
}
func RootnSi(rop FloatPtr, op FloatPtr, n int, rnd RoundMode) int {
	return int(C.mpfr_rootn_si(rop, op, C.long(n), C.mpfr_rnd_t(rnd)))
}

/* DEPRECATED
func Root(rop FloatPtr, op FloatPtr, n uint, rnd RoundMode) int {
	return int(C.mpfr_rootn_si(rop, op, C.ulong(n), C.mpfr_rnd_t(rnd)))
}
*/

func Neg(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_neg(rop, op, C.mpfr_rnd_t(rnd)))
}

func Abs(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_abs(rop, op, C.mpfr_rnd_t(rnd)))
}

func Mul2ui(rop FloatPtr, op1 FloatPtr, n uint, rnd RoundMode) int {
	return int(C.mpfr_mul_2ui(rop, op1, C.ulong(n), C.mpfr_rnd_t(rnd)))
}

func Mul2si(rop FloatPtr, op1 FloatPtr, n int, rnd RoundMode) int {
	return int(C.mpfr_mul_2si(rop, op1, C.long(n), C.mpfr_rnd_t(rnd)))
}

func Div2ui(rop FloatPtr, op1 FloatPtr, n uint, rnd RoundMode) int {
	return int(C.mpfr_div_2ui(rop, op1, C.ulong(n), C.mpfr_rnd_t(rnd)))
}

func Div2si(rop FloatPtr, op1 FloatPtr, n int, rnd RoundMode) int {
	return int(C.mpfr_div_2si(rop, op1, C.long(n), C.mpfr_rnd_t(rnd)))
}

func FacUi(rop FloatPtr, n uint, rnd RoundMode) int {
	return int(C.mpfr_fac_ui(rop, C.ulong(n), C.mpfr_rnd_t(rnd)))
}

// TOOD other functions

//
// 5.6 Comparison Functions
//

func Cmp(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_cmp(op1, op2))
}
func Cmpabs(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_cmpabs(op1, op2))
}
func NanP(op FloatPtr) int {
	return int(C.mpfr_nan_p(op))
}
func InfP(op FloatPtr) int {
	return int(C.mpfr_inf_p(op))
}
func NumberP(op FloatPtr) int {
	return int(C.mpfr_number_p(op))
}
func ZeroP(op FloatPtr) int {
	return int(C.mpfr_zero_p(op))
}
func RegularP(op FloatPtr) int {
	return int(C.mpfr_regular_p(op))
}
func Sgn(op FloatPtr) int {
	return int(C.mpfr_sgn(op))
}
func GreaterP(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_greater_p(op1, op2))
}
func GreaterequalP(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_greaterequal_p(op1, op2))
}
func LessP(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_less_p(op1, op2))
}
func LessequalP(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_lessequal_p(op1, op2))
}
func EqualP(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_equal_p(op1, op2))
}
func LessgreaterP(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_lessgreater_p(op1, op2))
}
func UnorderedP(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_unordered_p(op1, op2))
}
func TotalOrderP(op1 FloatPtr, op2 FloatPtr) int {
	return int(C.mpfr_total_order_p(op1, op2))
}

//
// 5.7 Transcendental Functions
//

func Log(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log(rop, op, C.mpfr_rnd_t(rnd)))
}

func LogUi(rop FloatPtr, op uint64, rnd RoundMode) int {
	return int(C.mpfr_log_ui(rop, C.ulong(op), C.mpfr_rnd_t(rnd)))
}

func Log2(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log2(rop, op, C.mpfr_rnd_t(rnd)))
}

func Log10(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log10(rop, op, C.mpfr_rnd_t(rnd)))
}

func Log1p(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log1p(rop, op, C.mpfr_rnd_t(rnd)))
}

func Log2p1(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log2p1(rop, op, C.mpfr_rnd_t(rnd)))
}

func Log10p1(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log10p1(rop, op, C.mpfr_rnd_t(rnd)))
}

func Exp(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp(rop, op, C.mpfr_rnd_t(rnd)))
}

func Exp2(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp2(rop, op, C.mpfr_rnd_t(rnd)))
}

func Exp10(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp10(rop, op, C.mpfr_rnd_t(rnd)))
}

func Expm1(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_expm1(rop, op, C.mpfr_rnd_t(rnd)))
}

func Exp2m1(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp2m1(rop, op, C.mpfr_rnd_t(rnd)))
}

func Exp10m1(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp10m1(rop, op, C.mpfr_rnd_t(rnd)))
}

//
// Powers
//

func Pow(rop FloatPtr, op1 FloatPtr, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_pow(rop, op1, op2, C.mpfr_rnd_t(rnd)))
}

func Powr(rop FloatPtr, op1 FloatPtr, op2 FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_powr(rop, op1, op2, C.mpfr_rnd_t(rnd)))
}

// Trigonometry
func Cos(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_cos(rop, op, C.mpfr_rnd_t(rnd)))
}
func Sin(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sin(rop, op, C.mpfr_rnd_t(rnd)))
}
func Tan(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_tan(rop, op, C.mpfr_rnd_t(rnd)))
}

func Cospi(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_cospi(rop, op, C.mpfr_rnd_t(rnd)))
}
func Sinpi(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sinpi(rop, op, C.mpfr_rnd_t(rnd)))
}
func Tanpi(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_tanpi(rop, op, C.mpfr_rnd_t(rnd)))
}

func Sec(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sec(rop, op, C.mpfr_rnd_t(rnd)))
}
func Csc(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_csc(rop, op, C.mpfr_rnd_t(rnd)))
}
func Cot(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_cot(rop, op, C.mpfr_rnd_t(rnd)))
}

func Acos(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_acos(rop, op, C.mpfr_rnd_t(rnd)))
}
func Asin(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_asin(rop, op, C.mpfr_rnd_t(rnd)))
}
func Atan(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_atan(rop, op, C.mpfr_rnd_t(rnd)))
}

func Acospi(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_acospi(rop, op, C.mpfr_rnd_t(rnd)))
}
func Asinpi(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_asinpi(rop, op, C.mpfr_rnd_t(rnd)))
}
func Atanpi(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_atanpi(rop, op, C.mpfr_rnd_t(rnd)))
}

// Special constants
func ConstLog2(rop FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_const_log2(rop, C.mpfr_rnd_t(rnd)))
}
func ConstPi(rop FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_const_pi(rop, C.mpfr_rnd_t(rnd)))
}
func ConstEuler(rop FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_const_euler(rop, C.mpfr_rnd_t(rnd)))
}
func ConstCatalan(rop FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_const_catalan(rop, C.mpfr_rnd_t(rnd)))
}

//
// 5.8 Input and Output Functinos
//

func Dump(x FloatPtr) {
	C.mpfr_dump(x)
}

//
// 5.9 Formatted Output Functions
//

func Printf2(template string, x FloatPtr) int {
	return int(C.mpfr_printf2(C.CString(template), x))
}

func Sprintf3(template string, rnd RoundMode, x FloatPtr) string {
	p := C.mpfr_sprintf3(C.CString(template), C.mpfr_rnd_t(rnd), x)
	if p == nil {
		return ""
	}
	s := C.GoString(p)
	C.mpfr_free_str(p)
	return s
}

func Sprintf2(template string, x FloatPtr) string {
	p := C.mpfr_sprintf2(C.CString(template), x)
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

func Rint(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint(rop, op, C.mpfr_rnd_t(rnd)))
}
func Ceil(rop FloatPtr, op FloatPtr) int {
	return int(C.mpfr_ceil(rop, op))
}
func Floor(rop FloatPtr, op FloatPtr) int {
	return int(C.mpfr_floor(rop, op))
}
func Round(rop FloatPtr, op FloatPtr) int {
	return int(C.mpfr_round(rop, op))
}
func Roundeven(rop FloatPtr, op FloatPtr) int {
	return int(C.mpfr_roundeven(rop, op))
}
func Trunc(rop FloatPtr, op FloatPtr) int {
	return int(C.mpfr_trunc(rop, op))
}
func Frac(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_frac(rop, op, C.mpfr_rnd_t(rnd)))
}
func RintCeil(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_ceil(rop, op, C.mpfr_rnd_t(rnd)))
}
func RintFloor(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_floor(rop, op, C.mpfr_rnd_t(rnd)))
}
func RintRound(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_round(rop, op, C.mpfr_rnd_t(rnd)))
}
func RintRoundeven(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_roundeven(rop, op, C.mpfr_rnd_t(rnd)))
}
func RintTrunc(rop FloatPtr, op FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_trunc(rop, op, C.mpfr_rnd_t(rnd)))
}

// 5.11 Rounding-Related Functions
func MinPrec(x FloatPtr) uint {
	return uint(C.mpfr_min_prec(x))
}
func PrecRound(x FloatPtr, prec int, rnd RoundMode) int {
	return int(C.mpfr_prec_round(x, C.mpfr_prec_t(prec), C.mpfr_rnd_t(rnd)))
}

//
// 5.12 Miscellaneous Functions
//

func Signbit(op FloatPtr) int {
	return int(C.mpfr_signbit(op))
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
