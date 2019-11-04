// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE_BSD file.

// code from golang 1.12 for compatibility

// +build !go1.12

package goint128

import "math/bits"
import "errors"


var divideError = errors.New("Divide by zero")
var overflowError = errors.New("Number overflow")

// Add64 returns the sum with carry of x, y and carry: sum = x + y + carry.
// The carry input must be 0 or 1; otherwise the behavior is undefined.
// The carryOut output is guaranteed to be 0 or 1.
func Add64(x, y, carry uint64) (sum, carryOut uint64) {
    yc := y + carry
    sum = x + yc
    if sum < x || yc < y {
        carryOut = 1
    }
    return
}

// Sub64 returns the difference of x, y and borrow: diff = x - y - borrow.
// The borrow input must be 0 or 1; otherwise the behavior is undefined.
// The borrowOut output is guaranteed to be 0 or 1.
func Sub64(x, y, borrow uint64) (diff, borrowOut uint64) {
    yb := y + borrow
    diff = x - yb
    if diff > x || yb < y {
            borrowOut = 1
    }
    return
}

// Mul64 returns the 128-bit product of x and y: (hi, lo) = x * y
// with the product bits' upper half returned in hi and the lower
// half returned in lo.
func Mul64(x, y uint64) (hi, lo uint64) {
    const mask32 = 1<<32 - 1
    x0 := x & mask32
    x1 := x >> 32
    y0 := y & mask32
    y1 := y >> 32
    w0 := x0 * y0
    t := x1*y0 + w0>>32
    w1 := t & mask32
    w2 := t >> 32
    w1 += x0 * y1
    hi = x1*y1 + w2 + w1>>32
    lo = x * y
    return
}

// Div64 returns the quotient and remainder of (hi, lo) divided by y:
// quo = (hi, lo)/y, rem = (hi, lo)%y with the dividend bits' upper
// half in parameter hi and the lower half in parameter lo.
// Div64 panics for y == 0 (division by zero) or y <= hi (quotient overflow).
func Div64(hi, lo, y uint64) (quo, rem uint64) {
    const (
        two32  = 1 << 32
        mask32 = two32 - 1
    )
    if y == 0 {
        panic(divideError)
    }
    if y <= hi {
        panic(overflowError)
    }

    s := uint(bits.LeadingZeros64(y))
    y <<= s

    yn1 := y >> 32
    yn0 := y & mask32
    un32 := hi<<s | lo>>(64-s)
    un10 := lo << s
    un1 := un10 >> 32
    un0 := un10 & mask32
    q1 := un32 / yn1
    rhat := un32 - q1*yn1

    for q1 >= two32 || q1*yn0 > two32*rhat+un1 {
        q1--
        rhat += yn1
        if rhat >= two32 {
            break
        }
    }

    un21 := un32*two32 + un1 - q1*y
    q0 := un21 / yn1
    rhat = un21 - q0*yn1

    for q0 >= two32 || q0*yn0 > two32*rhat+un0 {
        q0--
        rhat += yn1
        if rhat >= two32 {
            break
        }
    }

    return q1*two32 + q0, (un21*two32 + un0 - q0*y) >> s
}
