package big

import (
	"github.com/client9/bignum/mpz"
	"runtime"
)

type Int struct {
	ptr *mpz.Int
}

func NewInt(x int64) *Int {
	n := new(mpz.Int)
	mpz.InitSetSi(n, int(x))
	z := &Int{
		ptr: n,
	}
	runtime.AddCleanup(z, mpz.Clear, n)
	return z
}

func (z *Int) init() {
	if z.ptr == nil {
		n := new(mpz.Int)
		mpz.InitSetSi(n, 0)
		z.ptr = n
		runtime.AddCleanup(z, mpz.Clear, n)
	}
}

func NewIntTmp(x int64) *Int {
	n := new(mpz.Int)
	mpz.InitSetSi(n, int(x))
	z := &Int{
		ptr: n,
	}
	//runtime.AddCleanup(z, mpz.Clear, n)
	return z
}
func (z *Int) Clear() {
	if z.ptr != nil {
		mpz.Clear(z.ptr)
		z.ptr = nil
	}
}

func (z *Int) Abs(x *Int) *Int {
	z.init()
	x.init()
	mpz.Abs(z.ptr, x.ptr)
	return z
}

func (z *Int) Add(x, y *Int) *Int {
	z.init()
	x.init()
	y.init()
	mpz.Add(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO ANDNOT
// TODO APPEND
// TODO APPENDTEXT
// TODO BINOMIAL
// TODO BIT
// TODO BITLEN
// TODO BITS
// TODO BYTES

func (z *Int) Cmp(y *Int) *Int {
	z.init()
	y.init()
	mpz.Cmp(z.ptr, y.ptr)
	return z
}

func (z *Int) CmpAbs(y *Int) *Int {
	z.init()
	y.init()
	mpz.Cmpabs(z.ptr, y.ptr)
	return z
}

// TODO DIV
// TODO DIVMOD

func (z *Int) Exp(x, y, m *Int) *Int {
	z.init()
	x.init()
	y.init()
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

func (z *Int) Gcd(x, y *Int) *Int {
	z.init()
	x.init()
	y.init()
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
	z.init()
	x.init()
	y.init()
	mpz.Mod(z.ptr, x.ptr, y.ptr)
	return z
}

func (z *Int) ModInverse(x, y *Int) *Int {
	z.init()
	x.init()
	y.init()
	mpz.Invert(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO MODSQRT

func (z *Int) Mul(x, y *Int) *Int {
	z.init()
	x.init()
	y.init()
	mpz.Mul(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO MULRANGE
// TODO NEG

func (z *Int) Neg(x *Int) *Int {
	z.init()
	x.init()
	mpz.Neg(z.ptr, x.ptr)
	return z
}

func (z *Int) Not(x *Int) *Int {
	z.init()
	x.init()
	mpz.Com(z.ptr, x.ptr)
	return z
}

func (z *Int) Or(x, y *Int) *Int {
	z.init()
	x.init()
	y.init()
	mpz.Ior(z.ptr, x.ptr, y.ptr)
	return z
}

func (z *Int) ProbablyPrime(n int) bool {
	z.init()
	return mpz.ProbabPrimeP(z.ptr, n) == 1
}

func (z *Int) Quo(x, y *Int) *Int {
	z.init()
	x.init()
	y.init()
	mpz.TdivQ(z.ptr, x.ptr, y.ptr)
	return z
}

func (z *Int) QuoRem(x, y, r *Int) (*Int, *Int) {
	z.init()
	x.init()
	y.init()
	r.init()
	mpz.TdivQr(z.ptr, r.ptr, x.ptr, y.ptr)
	return z, r
}

// TODO RAND

func (z *Int) Rem(x, y *Int) *Int {
	z.init()
	x.init()
	y.init()
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
		n := new(mpz.Int)
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
		n := new(mpz.Int)
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
	z.init()
	x.init()
	mpz.Sqrt(z.ptr, x.ptr)
	return z
}

func (z *Int) String() string {
	z.init()
	return mpz.GetStr(10, z.ptr)
}

func (z *Int) Sub(x, y *Int) *Int {
	z.init()
	x.init()
	y.init()
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
	z.init()
	x.init()
	y.init()
	mpz.Xor(z.ptr, x.ptr, y.ptr)
	return z
}
