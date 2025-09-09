package bigmath

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
*/
import "C"

import (
	"runtime"
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

type FloatPtr C.mpfr_t

type Float struct {
	Ptr  FloatPtr
	init bool
}

func floatFinalize(z *Float) {
	if z.init {
		runtime.SetFinalizer(z, nil)
		Clear(&z.Ptr)
		//C.mpfr_clear(&z.Ptr[0])
		z.init = false
	}
}

func (z *Float) Init() {
	if z.init {
		return
	}
	Init(&z.Ptr)
	//C.mpfr_init(&z.Ptr[0])
	runtime.SetFinalizer(z, floatFinalize)
	z.init = true
}
func (z *Float) SetD(x float64, rnd RoundMode) int {
	return int(C.mpfr_set_d(&z.Ptr[0], C.double(x), C.mpfr_rnd_t(rnd)))
}

//
// 5.1 Initialization Functions
//

func Init(z *FloatPtr) {
	C.mpfr_init(&z[0])
}
func Init2(z *FloatPtr, prec int) {
	C.mpfr_init2(&z[0], C.mpfr_prec_t(prec))
}

func Clear(z *FloatPtr) {
	C.mpfr_clear(&z[0])
}

func SetDefaultPrec(prec int) {
	C.mpfr_set_default_prec(C.mpfr_prec_t(prec))
}

func GetDefaultPrec() int {
	return int(C.mpfr_get_default_prec())
}

func SetPrec(z *FloatPtr, prec int) {
	C.mpfr_set_prec(&z[0], C.mpfr_prec_t(prec))
}

func GetPrec(z *FloatPtr) int {
	return int(C.mpfr_get_prec(&z[0]))
}

//
// 5.2 Assignment Functions
//

func Set(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_set(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func SetUi(z *FloatPtr, op uint64, rnd RoundMode) int {
	return int(C.mpfr_set_ui(&z[0], C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func SetSi(z *FloatPtr, op int64, rnd RoundMode) int {
	return int(C.mpfr_set_si(&z[0], C.long(op), C.mpfr_rnd_t(rnd)))
}
func SetFlt(z *FloatPtr, val float32, rnd RoundMode) int {
	return int(C.mpfr_set_flt(&z[0], C.float(val), C.mpfr_rnd_t(rnd)))
}
func SetD(z *FloatPtr, val float64, rnd RoundMode) int {
	return int(C.mpfr_set_d(&z[0], C.double(val), C.mpfr_rnd_t(rnd)))
}

// TODO OTHER VARIATIONS

func SetNan(x *FloatPtr) {
	C.mpfr_set_nan(&x[0])
}

func SetInf(x *FloatPtr, sign int) {
	C.mpfr_set_inf(&x[0], C.int(sign))
}

func SetZero(x *FloatPtr, sign int) {
	C.mpfr_set_zero(&x[0], C.int(sign))
}

func Swap(x *FloatPtr, y *FloatPtr) {
	C.mpfr_swap(&x[0], &y[0])
}

//
// 5.3 Combined Initialization and Assignment Functions
//

func InitSet(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func InitSetUi(z *FloatPtr, op uint64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_ui(&z[0], C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func InitSetSi(z *FloatPtr, op int64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_si(&z[0], C.long(op), C.mpfr_rnd_t(rnd)))
}
func InitSetD(z *FloatPtr, val float64, rnd RoundMode) int {
	return int(C.fn_mpfr_init_set_d(&z[0], C.double(val), C.mpfr_rnd_t(rnd)))
}

// TODO: mpfr_set_ld "long double" is probably a 64-bit integer, but in some
// place it might be 80bit.

// TODO: mpfr_set_z, set_q, set_f

func InitSetStr(x *FloatPtr, s string, base int, rnd RoundMode) int {
	//cstr := unsafe.Pointer(C.CString(s))
	cstr := C.CString(s)
	ret := C.mpfr_set_str(&x[0], cstr, C.int(base), C.mpfr_rnd_t(rnd))
	C.free(unsafe.Pointer(cstr))
	return int(ret)
}

// 5.4 Conversion Functions
func GetFlt(op *FloatPtr, rnd RoundMode) float32 {
	return float32(C.mpfr_get_flt(&op[0], C.mpfr_rnd_t(rnd)))
}
func GetD(op *FloatPtr, rnd RoundMode) float64 {
	return float64(C.mpfr_get_d(&op[0], C.mpfr_rnd_t(rnd)))
}

func GetSi(op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_get_si(&op[0], C.mpfr_rnd_t(rnd)))
}
func GetUi(op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_get_ui(&op[0], C.mpfr_rnd_t(rnd)))
}

func GetStrNdigits(b int, prec int) int {
	return int(C.mpfr_get_str_ndigits(C.int(b), C.mpfr_prec_t(prec)))
}

// char *str, mpfr_exp_t *expptr, int base, size_t n, mpfr_t op, mpfr_rnd_t rnd)
func GetStr(base int, n int, op *FloatPtr, rnd RoundMode) (string, int) {
	var exp int
	cIntPtr := (*C.mpfr_exp_t)(unsafe.Pointer(&exp))
	p := C.mpfr_get_str(nil, cIntPtr, C.int(base), C.size_t(n), &op[0], C.mpfr_rnd_t(rnd))
	s := C.GoString(p)
	C.mpfr_free_str(p)
	return s, exp
}

// 5.5 Arithmetic Functions
func Add(rop *FloatPtr, op1 *FloatPtr, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_add(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}
func AddUi(rop *FloatPtr, op1 *FloatPtr, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_add_ui(&rop[0], &op1[0], C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func AddSi(rop *FloatPtr, op1 *FloatPtr, op2 int, rnd RoundMode) int {
	return int(C.mpfr_add_si(&rop[0], &op1[0], C.long(op2), C.mpfr_rnd_t(rnd)))
}
func AddD(rop *FloatPtr, op1 *FloatPtr, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_add_d(&rop[0], &op1[0], C.double(op2), C.mpfr_rnd_t(rnd)))
}

func Sub(rop *FloatPtr, op1 *FloatPtr, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sub(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}
func SubUi(rop *FloatPtr, op1 *FloatPtr, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_sub_ui(&rop[0], &op1[0], C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func SubSi(rop *FloatPtr, op1 *FloatPtr, op2 int, rnd RoundMode) int {
	return int(C.mpfr_sub_si(&rop[0], &op1[0], C.long(op2), C.mpfr_rnd_t(rnd)))
}
func SubD(rop *FloatPtr, op1 *FloatPtr, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_sub_d(&rop[0], &op1[0], C.double(op2), C.mpfr_rnd_t(rnd)))
}

func Mul(rop *FloatPtr, op1 *FloatPtr, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_mul(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}

func MulUi(rop *FloatPtr, op1 *FloatPtr, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_mul_ui(&rop[0], &op1[0], C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func MulSi(rop *FloatPtr, op1 *FloatPtr, op2 int, rnd RoundMode) int {
	return int(C.mpfr_mul_si(&rop[0], &op1[0], C.long(op2), C.mpfr_rnd_t(rnd)))
}
func MulD(rop *FloatPtr, op1 *FloatPtr, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_mul_d(&rop[0], &op1[0], C.double(op2), C.mpfr_rnd_t(rnd)))
}
func Sqr(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sqr(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Div(rop *FloatPtr, op1 *FloatPtr, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_div(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}

func DivUi(rop *FloatPtr, op1 *FloatPtr, op2 uint, rnd RoundMode) int {
	return int(C.mpfr_div_ui(&rop[0], &op1[0], C.ulong(op2), C.mpfr_rnd_t(rnd)))
}
func DivSi(rop *FloatPtr, op1 *FloatPtr, op2 int, rnd RoundMode) int {
	return int(C.mpfr_div_si(&rop[0], &op1[0], C.long(op2), C.mpfr_rnd_t(rnd)))
}
func DivD(rop *FloatPtr, op1 *FloatPtr, op2 float64, rnd RoundMode) int {
	return int(C.mpfr_div_d(&rop[0], &op1[0], C.double(op2), C.mpfr_rnd_t(rnd)))
}
func UiDiv(rop *FloatPtr, op1 uint, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_ui_div(&rop[0], C.ulong(op1), &op2[0], C.mpfr_rnd_t(rnd)))
}
func SiDiv(rop *FloatPtr, op1 int, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_si_div(&rop[0], C.long(op1), &op2[0], C.mpfr_rnd_t(rnd)))
}
func DDiv(rop *FloatPtr, op1 float64, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_d_div(&rop[0], C.double(op1), &op2[0], C.mpfr_rnd_t(rnd)))
}
func Sqrt(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sqrt(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func SqrtUi(rop *FloatPtr, op uint, rnd RoundMode) int {
	return int(C.mpfr_sqrt_ui(&rop[0], C.ulong(op), C.mpfr_rnd_t(rnd)))
}
func RecSqrt(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rec_sqrt(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Cbrt(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_cbrt(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RootnUi(rop *FloatPtr, op *FloatPtr, n uint, rnd RoundMode) int {
	return int(C.mpfr_rootn_ui(&rop[0], &op[0], C.ulong(n), C.mpfr_rnd_t(rnd)))
}
func RootnSi(rop *FloatPtr, op *FloatPtr, n int, rnd RoundMode) int {
	return int(C.mpfr_rootn_si(&rop[0], &op[0], C.long(n), C.mpfr_rnd_t(rnd)))
}

/* DEPRECATED
func Root(rop *FloatPtr, op *FloatPtr, n uint, rnd RoundMode) int {
	return int(C.mpfr_rootn_si(&rop[0], &op[0], C.ulong(n), C.mpfr_rnd_t(rnd)))
}
*/

func Neg(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_neg(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Abs(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_abs(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

// TOOD other functions

//
// 5.6 Comparison Functions
//

func Cmp(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_cmp(&op1[0], &op2[0]))
}
func Cmpabs(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_cmpabs(&op1[0], &op2[0]))
}
func Sgn(op *FloatPtr) int {
	return int(C.mpfr_sgn(&op[0]))
}
func GreaterP(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_greater_p(&op1[0], &op2[0]))
}
func GreaterequalP(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_greaterequal_p(&op1[0], &op2[0]))
}
func LessP(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_less_p(&op1[0], &op2[0]))
}
func LessequalP(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_lessequal_p(&op1[0], &op2[0]))
}
func EqualP(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_equal_p(&op1[0], &op2[0]))
}
func LessgreaterP(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_lessgreater_p(&op1[0], &op2[0]))
}
func UnorderedP(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_unordered_p(&op1[0], &op2[0]))
}
func TotalOrderP(op1 *FloatPtr, op2 *FloatPtr) int {
	return int(C.mpfr_total_order_p(&op1[0], &op2[0]))
}

//
// 5.7 Transcendental Functions
//

func Log(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func LogUi(rop *FloatPtr, op uint64, rnd RoundMode) int {
	return int(C.mpfr_log_ui(&rop[0], C.ulong(op), C.mpfr_rnd_t(rnd)))
}

func Log2(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log2(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Log10(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log10(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Log1p(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log1p(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Log2p1(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log2p1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Log10p1(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_log10p1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp2(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp2(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp10(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp10(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Expm1(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_expm1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp2m1(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp2m1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Exp10m1(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_exp10m1(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

//
// Powers
//

func Pow(rop *FloatPtr, op1 *FloatPtr, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_pow(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}

func Powr(rop *FloatPtr, op1 *FloatPtr, op2 *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_powr(&rop[0], &op1[0], &op2[0], C.mpfr_rnd_t(rnd)))
}

// Trigonometry
func Cos(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_cos(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Sin(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sin(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Tan(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_tan(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Cospi(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_cospi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Sinpi(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sinpi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Tanpi(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_tanpi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Sec(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_sec(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Csc(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_csc(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Cot(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_cot(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Acos(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_acos(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Asin(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_asin(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Atan(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_atan(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

func Acospi(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_acospi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Asinpi(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_asinpi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Atanpi(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_atanpi(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

// Special constants
func ConstLog2(rop *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_const_log2(&rop[0], C.mpfr_rnd_t(rnd)))
}
func ConstPi(rop *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_const_pi(&rop[0], C.mpfr_rnd_t(rnd)))
}
func ConstEuler(rop *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_const_euler(&rop[0], C.mpfr_rnd_t(rnd)))
}
func ConstCatalan(rop *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_const_catalan(&rop[0], C.mpfr_rnd_t(rnd)))
}

//
// 5.8 Input and Output Functinos
//

func Dump(x *FloatPtr) {
	C.mpfr_dump(&x[0])
}

//
// 5.9 Formatted Output Functions
//

func Printf2(template string, x *FloatPtr) int {
	return int(C.mpfr_printf2(C.CString(template), &x[0]))
}

//
// 5.10 Integer and Remainder Related Functions
//

func Rint(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func Ceil(rop *FloatPtr, op *FloatPtr) int {
	return int(C.mpfr_ceil(&rop[0], &op[0]))
}
func Floor(rop *FloatPtr, op *FloatPtr) int {
	return int(C.mpfr_floor(&rop[0], &op[0]))
}
func Round(rop *FloatPtr, op *FloatPtr) int {
	return int(C.mpfr_round(&rop[0], &op[0]))
}
func Roundeven(rop *FloatPtr, op *FloatPtr) int {
	return int(C.mpfr_roundeven(&rop[0], &op[0]))
}
func Trunc(rop *FloatPtr, op *FloatPtr) int {
	return int(C.mpfr_trunc(&rop[0], &op[0]))
}

func RintCeil(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_ceil(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RintFloor(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_floor(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RintRound(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_round(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RintRoundeven(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_roundeven(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}
func RintTrunc(rop *FloatPtr, op *FloatPtr, rnd RoundMode) int {
	return int(C.mpfr_rint_trunc(&rop[0], &op[0], C.mpfr_rnd_t(rnd)))
}

// 5.12 Miscellaneous Functions
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
