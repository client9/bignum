package mpz

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestMPZSize(t *testing.T) {
	var n Int
	fmt.Println(unsafe.Sizeof(n))
}
