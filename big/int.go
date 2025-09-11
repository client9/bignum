package big

import (
	"github.com/client9/bignum/mpz"
	"runtime"
)

// Zero value is *not* ready to use.
// It must be created by a New function or a Set function.
// Otherwise it will panic on nil pointer.
type Int struct {
	ptr *mpz.Int
}

func NewInt(x int64) *Int {
	var n mpz.Int
	mpz.InitSetSi(&n, int(x))
	z := &Int{
		ptr: &n,
	}
	runtime.AddCleanup(z, mpz.Clear, &n)
	return z
}
func NewIntTmp(x int64) *Int {
	var n mpz.Int
	mpz.InitSetSi(&n, int(x))
	z := &Int{
		ptr: &n,
	}
	//runtime.AddCleanup(z, mpz.Clear, &n)
	return z
}
func (z *Int) Clear() {
	mpz.Clear(z.ptr)
	z.ptr = nil
}

func (z *Int) Abs(x *Int) *Int {
	mpz.Abs(z.ptr, x.ptr)
	return z
}
func (z *Int) Add(x, y *Int) *Int {
	mpz.Add(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO AND
// TODO ANDNOT
// TODO APPEND
// TODO APPENDTEXT
// TODO BINOMIAL
// TODO BIT
// TODO BITLEN
// TODO BITS
// TODO BYTES

func (z *Int) Cmp(y *Int) *Int {
	mpz.Cmp(z.ptr, y.ptr)
	return z
}
func (z *Int) CmpAbs(y *Int) *Int {
	mpz.Cmpabs(z.ptr, y.ptr)
	return z
}

// TODO DIV
// TODO DIVMOD

func (z *Int) Exp(x, y, m *Int) *Int {
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
// TODO FLOAT64
// TODO FORMAT
// TODO GCD
// TODO GOBDECODE
// TODO GOBENCODE
// TODO INT64
// TODO ISINT64
// TODO ISUINT64
// TODO LSH
// TODO MARSHALJSON
// TODO MARSHALTEXT
// TODO MOD
// TODO MODINVERSE
// TODO MODSQRT

func (z *Int) Mul(x, y *Int) *Int {
	mpz.Mul(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO MULRANGE
// TODO NEG
// TODO NOT
// TODO OR

func (z *Int) ProbablyPrime(n int) bool {
	return mpz.ProbabPrimeP(z.ptr, n) == 1
}

// TODO QUO
// TODO QUOREM
// TODO RAND
// TODO REM
// TODO RSH
// TODO SCAN
// TODO SET
// TODO SETBIT
// TODO SETBITS
// TODO SETBYTES
// TODO SETINT64
// TODO SETSTRING
// TODO SETUINT64

func (z *Int) Sign() int {
	return mpz.Sgn(z.ptr)
}

// TODO SQRT

func (z *Int) String() string {
	return mpz.GetStr(10, z.ptr)
}
func (z *Int) Sub(x, y *Int) *Int {
	mpz.Sub(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO TEXT

// TODO TRAILINGZEROBITS
// TODO UINT64
// TODO UNMARSHALJSON
// TODO UNMARSHALTEXT
// TODO XOR
