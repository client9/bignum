package big

import (
	"github.com/client9/bignum/mpq"
	"runtime"
)

type Rat struct {
	ptr *mpq.Rat
}

func NewRat(a, b int64) *Rat {
	var n mpq.Rat

	mpq.Init(&n)
	mpq.SetSi(&n, int(a), uint(b))

	z := &Rat{
		ptr: &n,
	}
	runtime.AddCleanup(z, mpq.Clear, &n)
	return z
}
func NewRatTmp(a, b int64) *Rat {
	var n mpq.Rat
	mpq.Init(&n)
	mpq.SetSi(&n, int(a), uint(b))
	z := &Rat{
		ptr: &n,
	}
	//runtime.AddCleanup(z, mpq.Clear, &n)
	return z
}
func (z *Rat) init() {
	if z.ptr == nil {
		n := new(mpq.Rat)
		mpq.Init(n)
		z.ptr = n
		runtime.AddCleanup(z, mpq.Clear, n)
	}
}
func (z *Rat) Clear() {
	if z.ptr != nil {
		mpq.Clear(z.ptr)
		z.ptr = nil
	}
}

func (z *Rat) Abs(x *Rat) *Rat {
	z.init()
	x.init()
	mpq.Abs(z.ptr, x.ptr)
	return z
}
func (z *Rat) Add(x, y *Rat) *Rat {
	z.init()
	x.init()
	y.init()
	mpq.Add(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO AppendText

func (z *Rat) Cmp(y *Rat) *Rat {
	z.init()
	y.init()
	mpq.Cmp(z.ptr, y.ptr)
	return z
}

// TODO DENOM
// TODO Float32
// TODO Float64
// TODO FloatPrec
// TODO FloatString
// TODO GobDecode
// TODO GobEncode
// TODO IsInt
// TODO MarshalText

func (z *Rat) Mul(x, y *Rat) *Rat {
	z.init()
	x.init()
	y.init()
	mpq.Mul(z.ptr, x.ptr, y.ptr)
	return z
}

// TOOD NEG
// TODO NUM
// TODO QUO
// TOOD RatString
// TODO SCAN
// TODO SET
// TODO SETFLOAT64
// TODO SETFRAC
// TODO SETFRAC64
//    RETURNS PURE INT if denom == 1
// TODO SETINT
// TODO SETINT64
// TODO SETSTRING
// TODO SETUINT64

func (z *Rat) Sign() int {
	z.init()
	return mpq.Sgn(z.ptr)
}
func (z *Rat) String() string {
	z.init()
	return mpq.GetStr(10, z.ptr)
}
func (z *Rat) Sub(x, y *Rat) *Rat {
	z.init()
	x.init()
	y.init()
	mpq.Sub(z.ptr, x.ptr, y.ptr)
	return z
}

// TODO UNMARSHALTEXT
