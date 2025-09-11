package big

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestSizeOfInt(t *testing.T) {
	var n Int
	fmt.Println("Int", unsafe.Sizeof(n))
}
func TestSizeOfRat(t *testing.T) {
	var n Rat
	fmt.Println("Rat", unsafe.Sizeof(n))
}
func TestSizeOfFloat(t *testing.T) {
	var n Float
	fmt.Println("Float", unsafe.Sizeof(n))
}
