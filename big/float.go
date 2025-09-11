package big

import (
	stdlib "math/big"
	"runtime"

	"github.com/client9/bignum/mpfr"
)

// Zero value is *not* ready to use.
// It must be created by a New function or a Set function.
// Otherwise it will panic on nil pointer.
type Float struct {
	ptr *mpfr.Float

	// TODO: storing prec and mode natively takes 16 bytes on 64-bit platforms
	//   * No need for 64-bit values
	//   * prec should be a uint32
	//   * mode could also be a uint32
	// Total: 8 bytes
	prec uint
	mode mpfr.RoundMode
}

func NewFloat(x float64) *Float {
	var n mpfr.Float
	mpfr.InitSetD(&n, x, mpfr.RNDN)
	z := &Float{
		ptr:  &n,
		prec: 0,
		mode: 0,
	}
	runtime.AddCleanup(z, mpfr.Clear, &n)
	return z
}

// TODO PARSEFLOAT

func (z *Float) Abs(x *Float) *Float {
	mpfr.Abs(z.ptr, x.ptr, z.mode)
	return z
}

// TODO ACC

func (z *Float) Add(x, y *Float) *Float {
	if z.prec == 0 {
		z.prec = max(x.prec, y.prec)
	}
	mpfr.Add(z.ptr, x.ptr, y.ptr, z.mode)
	return z
}

// TODO APPEND
// TODO APPENDTEXT

// NONSTANDARD
// TODO make sure mpfr is ok with clearing Nul pointer.
func (z *Float) Clear() {
	mpfr.Clear(z.ptr)
	z.ptr = nil
	z.prec = 0
	z.mode = 0
}

func (x *Float) Cmp(y *Float) int {
	return mpfr.Cmp(x.ptr, y.ptr)
}

// TODO COPY

func (z *Float) Float32() float32 {
	return mpfr.GetFlt(z.ptr, z.mode)
}
func (z *Float) Float64() float64 {
	return mpfr.GetD(z.ptr, z.mode)
}

// TODO FORMAT
// TODO GOBDECODE
// TODO GOBENCODE

func (z *Float) Int() int {
	return int(mpfr.GetSi(z.ptr, z.mode))
}
func (z *Float) Int64() int64 {
	return int64(mpfr.GetSi(z.ptr, z.mode))
}

// TODO ISINF
// TODO ISINT
// TODO MANTEXP
// TODO MARSHALTEXT
// TODO MINPREC

func (z *Float) Mode() stdlib.RoundingMode {
	return exportRoundingMode(z.mode)
}
func (z *Float) Mul(x, y *Float) *Float {
	if z.prec == 0 {
		z.prec = max(x.prec, y.prec)
	}
	mpfr.Mul(z.ptr, x.ptr, y.ptr, z.mode)
	return z
}

func (z *Float) Neg(x *Float) *Float {
	mpfr.Neg(z.ptr, x.ptr, z.mode)
	return z
}

// TODO PARSE

func (z *Float) Prec(prec uint) uint {
	return z.prec
}

func (z *Float) Quo(x, y *Float) *Float {
	if z.prec == 0 {
		z.prec = max(x.prec, y.prec)
	}
	mpfr.Div(z.ptr, x.ptr, y.ptr, z.mode)
	return z
}

// TODO RAT
// TODO SCAN
// TODO SET
// TODO SETFLOAT64
// TODO SETINF
// TODO SETINT
// TODO SETINT64
// TODO SETMANTEXP

func (z *Float) SetMode(mode stdlib.RoundingMode) {
	z.mode = importRoundingMode(mode)
}

func (z *Float) SetPrec(prec uint) {
	mpfr.SetPrec(z.ptr, int(prec))
	z.prec = prec
}

// TODO SETRAT
// TODO SETSTRING
// TODO SETUINT64

func (x *Float) Sign() int {
	return mpfr.Sgn(x.ptr)
}

// TODO SIGNBIT

func (z *Float) Sqrt(x *Float) *Float {
	if z.prec == 0 {
		z.prec = x.prec
	}
	mpfr.Sqrt(z.ptr, x.ptr, z.mode)
	return z
}

func (z *Float) String() string {
	// matches Go
	return mpfr.Sprintf3("%.10R*g", z.mode, z.ptr)
}

func (z *Float) Sub(x, y *Float) *Float {
	if z.prec == 0 {
		z.prec = max(x.prec, y.prec)
	}
	mpfr.Sub(z.ptr, x.ptr, y.ptr, z.mode)
	return z
}

// TODO TEXT
// TODO UINT64
// TODO UNMARSHALTEXT

func max(a, b uint) uint {
	if a >= b {
		return a
	}
	return b
}

func importRoundingMode(r stdlib.RoundingMode) mpfr.RoundMode {
	switch r {
	case stdlib.ToNearestEven:
		return mpfr.RNDN
	case stdlib.ToNearestAway:
		panic("ToNearestAway RoundingMode not supported")
	case stdlib.ToZero:
		return mpfr.RNDZ
	case stdlib.AwayFromZero:
		return mpfr.RNDA
	case stdlib.ToNegativeInf:
		return mpfr.RNDD
	case stdlib.ToPositiveInf:
		return mpfr.RNDU
	default:
		panic("unknown rounding mode")
	}
}

func exportRoundingMode(r mpfr.RoundMode) stdlib.RoundingMode {
	switch r {
	case mpfr.RNDN:
		return stdlib.ToNearestEven
	case mpfr.RNDZ:
		return stdlib.ToZero
	case mpfr.RNDA:
		return stdlib.AwayFromZero
	case mpfr.RNDD:
		return stdlib.ToNegativeInf
	case mpfr.RNDU:
		return stdlib.ToPositiveInf
	default:
		panic("unsupported MPFR rounding mode")
	}
}
