package big

import (
	"github.com/client9/bignum/mpz"
	"runtime"
)

// Zero value is *not* ready to use.  It must be created by a New function or a Set function.
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
func (z *Int) Cmp(y *Int) *Int {
	mpz.Cmp(z.ptr, y.ptr)
	return z
}
func (z *Int) CmpAbs(y *Int) *Int {
	mpz.Cmpabs(z.ptr, y.ptr)
	return z
}

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
func (z *Int) Mul(x, y *Int) *Int {
	mpz.Mul(z.ptr, x.ptr, y.ptr)
	return z
}
func (z *Int) ProbablyPrime(n int) bool {
	return mpz.ProbabPrimeP(z.ptr, n) == 1
}
func (z *Int) Sign() int {
	return mpz.Sgn(z.ptr)
}
func (z *Int) String() string {
	return mpz.GetStr(10, z.ptr)
}
func (z *Int) Sub(x, y *Int) *Int {
	mpz.Sub(z.ptr, x.ptr, y.ptr)
	return z
}
