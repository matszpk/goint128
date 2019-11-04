// code

// +build go1.12

package goint128

import "math/bits"

func Add64(x, y, carry uint64) (sum, carryOut uint64) {
    return bits.Add64(x, y, carry)
}

func Sub64(x, y, carry uint64) (sum, carryOut uint64) {
    return bits.Sub64(x, y, carry)
}

func Mul64(x, y uint64) (hi, lo uint64) {
    return bits.Mul64(x, y)
}

func Div64(hi, lo, y uint64) (quo, rem uint64) {
    return bits.Div64(hi, lo, y)
}
