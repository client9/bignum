package big

import (
	"runtime"
	"unsafe"

	"github.com/client9/bignum/mpz"
)

func newIntPtr() mpz.IntPtr {
	return mpz.IntPtr(unsafe.Pointer(new(mpz.Int)))
}

type Int struct {
	ptr mpz.IntPtr
}

func NewInt(x int64) *Int {
	n := newIntPtr()
	mpz.InitSetSi(n, int(x))
	z := &Int{
		ptr: n,
	}
	runtime.AddCleanup(z, mpz.Clear, n)
	return z
}

func (z *Int) init() {
	n := newIntPtr()
	mpz.InitSetSi(n, 0)
	z.ptr = n
	runtime.AddCleanup(z, mpz.Clear, n)
}

func NewIntTmp(x int64) *Int {
	n := newIntPtr()
	mpz.InitSetSi(n, int(x))
	z := &Int{
		ptr: n,
	}
	return z
}
func (z *Int) Clear() {
	if z.ptr != nil {
		mpz.Clear(z.ptr)
		z.ptr = nil
	}
}

func (z *Int) Abs(x *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	mpz.Abs(z.ptr, x.ptr)
	return z
}

func (z *Int) Add(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.Add(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO ANDNOT -- write C so we don't make two round trips.
// TODO APPEND
// TODO APPENDTEXT
// TODO BINOMIAL -- there is a MPZ function for this
// TODO BIT

func (z *Int) BitLen() int {
	if z.ptr == nil {
		return 0
	}
	return mpz.Sizeinbase(z.ptr, 2)
}

// TODO BITS
// TODO BYTES

func (z *Int) Cmp(y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if y.ptr == nil {
		y.init()
	}

	mpz.Cmp(z.ptr, y.ptr)
	return z
}

func (z *Int) CmpAbs(y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if y.ptr == nil {
		y.init()
	}

	mpz.Cmpabs(z.ptr, y.ptr)
	return z
}

func (z *Int) Div(x, y *Int) *Int {
	sgn := y.Sign()
	if sgn == 0 {
		panic("division by zero")
	}
	if z.ptr == nil {
		z.init()
	}

	// TODO: if x or m is unini we can skip steps
	if x.ptr == nil {
		x.init()
	}

	if sgn == 1 {
		mpz.FdivR(z.ptr, x.ptr, y.ptr)
	} else {
		mpz.CdivR(z.ptr, x.ptr, y.ptr)
	}
	return z
}

func (z *Int) DivMod(x, y, m *Int) (*Int, *Int) {
	// Sign handles uninitialized values so this is ok
	sgn := y.Sign()
	if sgn == 0 {
		panic("division by zero")
	}

	if z.ptr == nil {
		z.init()
	}

	// TODO: if x or m is unini we can skip steps
	if x.ptr == nil {
		x.init()
	}
	if m.ptr == nil {
		m.init()
	}

	if sgn == 1 {
		mpz.FdivQr(z.ptr, m.ptr, x.ptr, y.ptr)
	} else {
		mpz.CdivQr(z.ptr, m.ptr, x.ptr, y.ptr)
	}
	return z, m
}

func (z *Int) Exp(x, y, m *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	//
	// it's ok for m to be nil or uninitialized
	//
	if y.Sign() <= 0 {
		mpz.SetUi(z.ptr, 1)
		return z
	}

	if m == nil || m.Sign() == 0 {
		mpz.PowUi(z.ptr, x.ptr, mpz.GetUi(y.ptr))
	} else {
		mpz.Powm(z.ptr, x.ptr, y.ptr, m.ptr)
	}
	return z
}

// TODO FILLBYTES

func (z *Int) Float64() float64 {
	if z.ptr == nil {
		return 0.0
	}
	return mpz.GetD(z.ptr)
}

// TODO FORMAT

func (z *Int) GCD(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.Gcd(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO GOBDECODE
// TODO GOBENCODE

func (z *Int) Int64() int64 {
	if z.ptr == nil {
		return 0
	}
	return int64(mpz.GetSi(z.ptr))
}

func (z *Int) IsInt64() bool {
	if z.ptr == nil {
		return true
	}
	return mpz.FitsSlongP(z.ptr) == 1
}

func (z *Int) IsUint64() bool {
	if z.ptr == nil {
		return true
	}
	return mpz.FitsUlongP(z.ptr) == 1
}

// TODO LSH
// TODO MARSHALJSON
// TODO MARSHALTEXT

func (z *Int) Mod(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.Mod(z.ptr, x.ptr, y.ptr)
	return z
}

func (z *Int) ModInverse(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.Invert(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO MODSQRT

// TODO: if any are uninitialized, we can return 0
func (z *Int) Mul(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.Mul(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO MULRANGE

// TODO: Whats Neg(0), if 0, then we can inline
func (z *Int) Neg(x *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	mpz.Neg(z.ptr, x.ptr)
	return z
}

func (z *Int) Not(x *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	mpz.Com(z.ptr, x.ptr)
	return z
}

func (z *Int) Or(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.Ior(z.ptr, x.ptr, y.ptr)
	return z
}

func (z *Int) ProbablyPrime(n int) bool {
	if z.ptr == nil {
		return false
	}
	return mpz.ProbabPrimeP(z.ptr, n) == 1
}

func (z *Int) Quo(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.TdivQ(z.ptr, x.ptr, y.ptr)
	return z
}

func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int) {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	if r.ptr == nil {
		r.init()
	}
	mpz.TdivQr(z.ptr, r.ptr, x.ptr, y.ptr)
	return z, r
}

// TODO RAND

func (z *Int) Rem(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.TdivR(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO RSH
// TODO SCAN
// TODO SET
// TODO SETBIT
// TODO SETBITS
// TODO SETBYTES

func (z *Int) SetInt64(x int64) {
	if z.ptr == nil {
		n := newIntPtr()
		mpz.InitSetSi(n, int(x))
		z.ptr = n
		runtime.AddCleanup(z, mpz.Clear, n)
	} else {
		mpz.SetSi(z.ptr, int(x))
	}
}

// TODO SETSTRING

func (z *Int) SetUint64(x uint64) {
	if z.ptr == nil {
		n := newIntPtr()
		mpz.InitSetUi(n, uint(x))
		z.ptr = n
		runtime.AddCleanup(z, mpz.Clear, n)
	} else {
		mpz.SetUi(z.ptr, uint(x))
	}
}

func (z *Int) Sign() int {
	if z.ptr == nil {
		return 0
	}
	return mpz.Sgn(z.ptr)
}

func (z *Int) Sqrt(x *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	mpz.Sqrt(z.ptr, x.ptr)
	return z
}

func (z *Int) String() string {
	if z.ptr == nil {
		return ""
	}
	return mpz.GetStr(10, z.ptr)
}

func (z *Int) Sub(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.Sub(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO TEXT
// TODO TRAILINGZEROBITS

func (z *Int) Uint64() uint64 {
	if z.ptr == nil {
		return 0
	}
	return uint64(mpz.GetUi(z.ptr))
}

// TODO UNMARSHALJSON
// TODO UNMARSHALTEXT

func (z *Int) Xor(x, y *Int) *Int {
	if z.ptr == nil {
		z.init()
	}
	if x.ptr == nil {
		x.init()
	}
	if y.ptr == nil {
		y.init()
	}
	mpz.Xor(z.ptr, x.ptr, y.ptr)
	return z
}
