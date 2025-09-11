package mpq

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMPQSize(t *testing.T) {
	var n Rat
	fmt.Println(unsafe.Sizeof(n))
}
func TestMPQSmoke(t *testing.T) {
	var r Rat
	Init(&r)
	SetUi(&r, 1, 223874823784)
	fmt.Println(GetStr(10, &r))
}
