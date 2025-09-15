package big

import (
	"bytes"
	"fmt"
	std "math/big"
	"testing"
)

type Integer interface {
	Add(*Integer, *Integer) *Integer
}

func testAddition(i0, i1 string) bool {
	var x, y, z std.Int
	var a, b, c Int

	_, ok1 := x.SetString(i0, 10)
	_, ok2 := y.SetString(i1, 10)

	_, ok3 := a.SetString(i0, 10)
	_, ok4 := b.SetString(i1, 10)

	// parsing didn't match
	if ok1 != ok3 || ok2 != ok4 {
		return false
	}

	// if bad input, skip remaining part of test
	if !ok1 || !ok2 {
		return true
	}
	z.Add(&x, &y)

	a.SetString(i0, 10)
	b.SetString(i1, 10)
	c.Add(&a, &b)
	if z.String() != c.String() {
		fmt.Println(z.String(), c.String())
	}
	return z.String() == c.String()
}

var cases = []string{
	"",
	"+",
	"-",
	"0",
	"-0",
	"1",
	"-1",
	"2",
	"-2",
	"0xDEADBEEFdeadbeef",
	"0XDEADBEEFdeadbeef",
	"0o6662536",
	"0O6762354",
	"01234567",
	"012345678",
	"-0827717676",
	"2147483647",
	"2147483648",
	"-2147483647",
	"-2147483648",
	"9223372036854775807",
	"9223372036854775808",
	"-9223372036854775807",
	"-9223372036854775808",
	"18446744073709551615",
	"18446744073709551616",
	"-18446744073709551615",
	"-18446744073709551616",

	"100000000000000000000000000000000000000000000000000000000000000",
}

func FuzzImportExport(f *testing.F) {

	base := []int{0, 2, 8, 10, 16}

	for _, b := range base {
		for _, tc := range cases {
			f.Add(tc, b)
		}
	}
	f.Fuzz(func(t *testing.T, s string, b int) {
		testSetString(t, s, b)
	})
}

func TestSetString(t *testing.T) {
	var x Int
	_, ok := x.SetString("0xDEADBEEFdeadbeef", 0)
	if !ok {
		t.Errorf("Set String failed on hex")
	}
}

func testSetString(t *testing.T, s string, b int) {

	var xStd std.Int
	var xGmp Int

	// fuzzers will send in negative numbers which will panic std math/big
	// ignore here
	if b < 0 || b == 1 || b > 62 {
		return
	}
	// import as string
	_, okStd := xStd.SetString(s, b)
	_, okGmp := xGmp.SetString(s, b)
	if okStd != okGmp {
		t.Errorf("string parsing results are different for %q base %d, std=%v, gmp=%v", s, b, okStd, okGmp)
		return
	}

	// if invalid input skip rest of test
	//  std behavior is undefined anyways
	if !okStd {
		return
	}

	if xGmp.String() != xStd.String() {
		t.Errorf("Expected same string output for %q base %d  got std=%q, gmp=%q", s, b, xStd.String(), xGmp.String())
		return
	}

	// export as bytes
	xbGmp := xGmp.Bytes()
	xbStd := xStd.Bytes()
	if !bytes.Equal(xbGmp, xbStd) {
		t.Errorf("Expected bytes output to be the same: %s", s)
		return
	}

	// clear
	xGmp.SetInt64(0)
	xStd.SetInt64(0)

	// import as bytes
	xGmp.SetBytes(xbGmp)
	xStd.SetBytes(xbStd)

	// export as strings
	if xGmp.String() != xStd.String() {
		t.Errorf("Expected same: %s", s)
	}
}

func TestInt(t *testing.T) {
	a := NewInt(1000)
	b := NewInt(10)
	a.Add(a, b)
	b.Add(a, b)
	fmt.Println(a.String())
}

// should not explode
func TestIntZeroValue(t *testing.T) {
	var a, b Int
	a.Add(&a, &b)
}
