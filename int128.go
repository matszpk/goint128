/*
 * int128.go - int128 routines
 *
 * goint128 - go int128 library
 * Copyright (C) 2020  Mateusz Szpakowski
 *
 * This library is free software; you can redistribute it and/or
 * modify it under the terms of the GNU Lesser General Public
 * License as published by the Free Software Foundation; either
 * version 2.1 of the License, or (at your option) any later version.
 *
 * This library is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the GNU
 * Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public
 * License along with this library; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA  02110-1301  USA
 */

// Package to operate on 128-bit integers
package goint128

import (
    "bytes"
    "encoding/binary"
    "errors"
    "math"
    "math/bits"
    "sort"
    "strconv"
)

type UInt128 [2]uint64

// add 128-bit unsigned integers
func (a UInt128) Add(b UInt128) UInt128 {
    var c UInt128
    var carry uint64
    c[0], carry = Add64(a[0], b[0], 0)
    c[1], _ = Add64(a[1], b[1], carry)
    return c
}

// add 128-bit unsigned integers with carry and return sum and output carry
func (a UInt128) AddC(b UInt128, oldCarry uint64) (UInt128, uint64) {
    var c UInt128
    var carry uint64
    c[0], carry = Add64(a[0], b[0], oldCarry)
    c[1], carry = Add64(a[1], b[1], carry)
    return c, carry
}

// add 128-bit unsigned integer and 64-bit unsigned integer
func (a UInt128) Add64(b uint64) UInt128 {
    var c UInt128
    var carry uint64
    c[0], carry = Add64(a[0], b, 0)
    c[1], _ = Add64(a[1], 0, carry)
    return c
}

// subtract 128-bit unsigned integers
func (a UInt128) Sub(b UInt128) UInt128 {
    var c UInt128
    var borrow uint64
    c[0], borrow = Sub64(a[0], b[0], 0)
    c[1], _ = Sub64(a[1], b[1], borrow)
    return c
}

// subtract 128-bit unsigned integers with borrow and return difference and borrow
func (a UInt128) SubB(b UInt128, oldBorrow uint64) (UInt128, uint64) {
    var c UInt128
    var borrow uint64
    c[0], borrow = Sub64(a[0], b[0], oldBorrow)
    c[1], borrow = Sub64(a[1], b[1], borrow)
    return c, borrow
}

// subtract 64-bit unsigned from 128-bit unsigned integer
func (a UInt128) Sub64(b uint64) UInt128 {
    var c UInt128
    var borrow uint64
    c[0], borrow = Sub64(a[0], b, 0)
    c[1], _ = Sub64(a[1], 0, borrow)
    return c
}

// compare 128-bit unsigned integer and return 0 if they equal,
// 1 if first is greater than second, or -1 if first is lesser than second
func (a UInt128) Cmp(b UInt128) int {
    if a[1]==b[1] {
        if a[0]==b[0] {
            return 0
        } else if a[0]>b[0] {
            return 1
        } else {
            return -1
        }
    } else if a[1]>b[1] {
        return 1
    } else { // a[1]<b[1]
        return -1
    }
}

// multiply 128-bit unsinged integers and return lower 128 bits value
func (a UInt128) Mul(b UInt128) UInt128 {
    var c UInt128
    c[1], c[0] = Mul64(a[0], b[0])
    c[1] += a[1]*b[0] + a[0]*b[1]
    return c
}

// multiply 128-bit unsigned integer and 64-bit unsigned integer and
// return lower 128 bits product
func (a UInt128) Mul64(b uint64) UInt128 {
    var c UInt128
    c[1], c[0] = Mul64(a[0], b)
    c[1] += a[1]*b
    return c
}

// multiply 128-bit unsigned integers and return high and lower product
func (a UInt128) MulFull(b UInt128) (UInt128, UInt128) {
    var clo, cm1, cm2, chi UInt128
    clo[1], clo[0] = Mul64(a[0], b[0])
    cm1[1], cm1[0] = Mul64(a[1], b[0])
    cm2[1], cm2[0] = Mul64(a[0], b[1])
    chi[1], chi[0] = Mul64(a[1], b[1])
    var carry uint64
    clo[1], carry = Add64(clo[1], cm1[0], 0)
    chi[0], carry = Add64(chi[0], cm1[1], carry)
    chi[1], _ = Add64(chi[1], 0, carry)
    clo[1], carry = Add64(clo[1], cm2[0], 0)
    chi[0], carry = Add64(chi[0], cm2[1], carry)
    chi[1], _ = Add64(chi[1], 0, carry)
    return chi, clo
}

// shift 128-bit unsigned integer left by b bits
func (a UInt128) Shl(b uint) UInt128 {
    if b==0 { return a }
    if b>=64 {
        return UInt128{ 0, a[0]<<(b-64) }
    }
    return UInt128{ a[0]<<b, (a[1]<<b) | (a[0]>>(64-b)) }
}

// shift 128-bit unsigned integer right by b bits
func (a UInt128) Shr(b uint) UInt128 {
    if b==0 { return a }
    if b>=64 {
        return UInt128{ a[1]>>(b-64), 0 }
    }
    return UInt128{ (a[0]>>b) | (a[1]<<(64-b)), a[1]>>b }
}

// divide 128-bit unsigned integer by 64-bit unsigned integer and
// return quotient and 64-bit remainder
func (a UInt128) Div64(b uint64) (UInt128, uint64) {
    var c UInt128
    if b==0 {
        panic("Divide by zero")
    }
    if a[1]>=b {
        // higher than 64-bit value - use slow division
        shift := int(bits.LeadingZeros64(b) - bits.LeadingZeros64(a[1]))
        var blo uint64 = 0 // low bits of b
        b <<= uint(shift)
        shift += 64
        tmp := a
        tmpa := a
        c[0], c[1] = 0, 0
        var borrow uint64
        for ; shift>=0; shift-- {
            // a - (b<<X)
            tmpa[0], borrow = Sub64(tmp[0], blo, 0)
            tmpa[1], borrow = Sub64(tmp[1], b, borrow)
            c[1] = (c[0]>>63) | (c[1]<<1) // shift
            c[0] <<= 1
            if borrow==0 {
                // tmpa>=(b<<X) - then add 1
                tmp = tmpa
                c[0] |= 1
            }
            blo = (blo>>1) | (b<<63)
            b >>= 1
        }
        return c, tmp[0]
    }
    var rem uint64
    c[1] = 0
    c[0], rem = Div64(a[1], a[0], b)
    return c, rem
}

// divide 256-bit (lo, hi) unsigned integer by 128-bit unsigned integer and return
// 128-bit quotient and remainder
func UInt128DivFull(hi, lo, b UInt128) (UInt128, UInt128) {
    if b[1]==0 && hi[0]==0 && hi[1]==0 {
        c, rem := lo.Div64(b[0])
        return c, UInt128{rem, 0}
    }
    if b[0]==0 && b[1]==0 {
        panic("Divide by zero")
    }
    var borrow uint64
    lza := 0
    if hi[0]==0 && hi[1]==0 {
        lza = 128
    } else if hi[1]!=0 {
        lza = bits.LeadingZeros64(hi[1])
    } else {
        lza = bits.LeadingZeros64(hi[0])+64
    }
    lzb := 0
    if b[1]!=0 {
        lzb = bits.LeadingZeros64(b[1])
    } else {
        lzb = bits.LeadingZeros64(b[0])+64
    }
    // check overflow
    if lza < lzb {
        panic("Divide overflow")
    } else if lza==lzb {
        _, borrow = Sub64(hi[0], b[0], 0)
        _, borrow = Sub64(hi[1], b[1], borrow)
        if borrow==0 { // hi>=b
            panic("Divide overflow")
        }
    }
    sh := uint(lza-lzb)
    pos := int(128-sh)
    // shift A (lo,hi) by shift (move to highest bit of b)
    var tlo, thi UInt128
    if sh!=128 {
        tlo = lo.Shl(sh)
        thi = hi.Shl(sh)
        if sh!=0 {
            tmp := lo.Shr(128-sh)
            thi[0] |= tmp[0]
            thi[1] |= tmp[1]
        }
    } else {
        thi = lo
        tlo[0], tlo[1] = 0, 0
    }
    // main loop
    var tmp UInt128
    c := UInt128{0,0}
    for ; pos>0; pos-- {
        tmp[0], borrow = Sub64(thi[0], b[0], 0)
        tmp[1], borrow = Sub64(thi[1], b[1], borrow)
        c[1] = (c[0]>>63) | (c[1]<<1) // shift
        c[0] <<= 1
        if borrow==0 {
            thi = tmp
            c[0] |= 1
        }
        // shift T (shifted copy of A)
        thi[1] = (thi[0]>>63) | (thi[1]<<1) // shift
        thi[0] = (tlo[1]>>63) | (thi[0]<<1)
        tlo[1] = (tlo[0]>>63) | (tlo[1]<<1)
        tlo[0] <<= 1
    }
    // last iteration
    tmp[0], borrow = Sub64(thi[0], b[0], 0)
    tmp[1], borrow = Sub64(thi[1], b[1], borrow)
    c[1] = (c[0]>>63) | (c[1]<<1) // shift
    c[0] <<= 1
    if borrow==0 {
        thi = tmp
        c[0] |= 1
    }
    return c, thi
}

var uint128_10powers []UInt128 = []UInt128{
    UInt128{1, 0},
    UInt128{10, 0},
    UInt128{100, 0},
    UInt128{1000, 0},
    UInt128{10000, 0},
    UInt128{100000, 0},
    UInt128{1000000, 0},
    UInt128{10000000, 0},
    UInt128{100000000, 0},
    UInt128{1000000000, 0},
    UInt128{10000000000, 0},
    UInt128{100000000000, 0},
    UInt128{1000000000000, 0},
    UInt128{10000000000000, 0},
    UInt128{100000000000000, 0},
    UInt128{1000000000000000, 0},
    UInt128{10000000000000000, 0},
    UInt128{100000000000000000, 0},
    UInt128{1000000000000000000, 0},
    UInt128{10000000000000000000, 0},
    UInt128{7766279631452241920, 5},
    UInt128{3875820019684212736, 54},
    UInt128{1864712049423024128, 542},
    UInt128{200376420520689664, 5421},
    UInt128{2003764205206896640, 54210},
    UInt128{1590897978359414784, 542101},
    UInt128{15908979783594147840, 5421010},
    UInt128{11515845246265065472, 54210108},
    UInt128{4477988020393345024, 542101086},
    UInt128{7886392056514347008, 5421010862},
    UInt128{5076944270305263616, 54210108624},
    UInt128{13875954555633532928, 542101086242},
    UInt128{9632337040368467968, 5421010862427},
    UInt128{4089650035136921600, 54210108624275},
    UInt128{4003012203950112768, 542101086242752},
    UInt128{3136633892082024448, 5421010862427522},
    UInt128{12919594847110692864, 54210108624275221},
    UInt128{68739955140067328,    542101086242752217},
    UInt128{687399551400673280,   5421010862427522170},
}

// format 128-bit unsigned integer to bytes
func (a UInt128) FormatBytes() []byte {
    if a[0]==0 && a[1]==0 { return []byte{ '0' } }
    var borrow uint64
    var chars [41]byte
    i := sort.Search(len(uint128_10powers), func(ii int) bool {
        _, borrow = Sub64(a[0], uint128_10powers[ii][0], 0)
        _, borrow = Sub64(a[1], uint128_10powers[ii][1], borrow)
        return borrow!=0 // a<uint128_10powers[ii]
    })-1
    end := i
    if i<19 {
        var tmpa, tmp, x, x1 uint64
        tmp = a[0]
        for ; i>=0; i-- {
            // calculate digit
            x = uint128_10powers[i][0]
            var digit byte = '0'
            {
                x1 = x<<3
                // check if 3 bit of digit - 8
                tmpa, borrow = Sub64(tmp, x1, 0)
                if borrow==0 {
                    digit += 8
                    tmp = tmpa
                }
                x1 = x<<2
                // check if 2 bit of digit - 4
                tmpa, borrow = Sub64(tmp, x1, 0)
                if borrow==0 {
                    digit += 4
                    tmp = tmpa
                }
            }
            // check if 1 bit of digit - 2
            x1 = x<<1
            tmpa, borrow = Sub64(tmp, x1, 0)
            if borrow==0 {
                digit += 2
                tmp = tmpa
            }
            // check if 0 bit of digit - 1
            tmpa, borrow = Sub64(tmp, x, 0)
            if borrow==0 {
                digit++
                tmp = tmpa
            }
            chars[40-i] = digit
        }
    } else {
        var tmpa, tmp, x, x1 UInt128
        var borrow uint64
        tmp = a
        for ; i>=0; i-- {
            // calculate digit
            x = uint128_10powers[i]
            var digit byte = '0'
            if i<38 {
                x1[1] = (x[1]<<3) | (x[0]>>61)
                x1[0] = x[0]<<3
                // check if 3 bit of digit - 8
                tmpa[0], borrow = Sub64(tmp[0], x1[0], 0)
                tmpa[1], borrow = Sub64(tmp[1], x1[1], borrow)
                if borrow==0 {
                    digit += 8
                    tmp = tmpa
                }
                x1[1] = (x[1]<<2) | (x[0]>>62)
                x1[0] = x[0]<<2
                // check if 2 bit of digit - 4
                tmpa[0], borrow = Sub64(tmp[0], x1[0], 0)
                tmpa[1], borrow = Sub64(tmp[1], x1[1], borrow)
                if borrow==0 {
                    digit += 4
                    tmp = tmpa
                }
            }
            // check if 1 bit of digit - 2
            x1[1] = (x[1]<<1) | (x[0]>>63)
            x1[0] = x[0]<<1
            tmpa[0], borrow = Sub64(tmp[0], x1[0], 0)
            tmpa[1], borrow = Sub64(tmp[1], x1[1], borrow)
            if borrow==0 {
                digit += 2
                tmp = tmpa
            }
            // check if 0 bit of digit - 1
            tmpa[0], borrow = Sub64(tmp[0], x[0], 0)
            tmpa[1], borrow = Sub64(tmp[1], x[1], borrow)
            if borrow==0 {
                digit++
                tmp = tmpa
            }
            chars[40-i] = digit
        }
    }
    return chars[40-end:]
}

// format 128-bit unsigned integer to string
func (a UInt128) Format() string {
    return string(a.FormatBytes())
}

// parse unsigned integer from string and return value and error (nil if no error)
func ParseUInt128(str string) (UInt128, error) {
    lastDigitValue := UInt128{ 11068046444225730969, 1844674407370955161 }
    slen := len(str)
    var out UInt128
    var carry uint64
    var i int
    for i=0; i<19 && i<slen && str[i]>='0' && str[i]<='9'; i++ {
        digit := byte(str[i])-'0'
        // multiply by 10
        out[0] *= 10
        // add digit
        out[0] += uint64(digit)
    }
    
    for ; i<slen && str[i]>='0' && str[i]<='9'; i++ {
        if out[1]>lastDigitValue[1] ||
            (out[1]==lastDigitValue[1] && out[0] > lastDigitValue[0]) {
            return UInt128{}, strconv.ErrRange
        }
        digit := byte(str[i])-'0'
        temp := out
        // multiply by 10
        out[1] = (temp[1]<<3) | (temp[0]>>61)
        out[0] = temp[0]<<3
        out[0], carry = Add64(out[0], temp[0]<<1, 0)
        out[1], _ = Add64(out[1], (temp[1]<<1) | (temp[0]>>63), carry)
        // add digit
        out[0], carry = Add64(out[0], uint64(digit), 0)
        out[1], carry = Add64(out[1], 0, carry)
        if carry!=0 {
            return UInt128{}, strconv.ErrRange
        }
    }
    if i==0 || i!=slen {
        return UInt128{}, strconv.ErrSyntax
    }
    return out, nil
}

func ParseUInt128Bytes(str []byte) (UInt128, error) {
    lastDigitValue := UInt128{ 11068046444225730969, 1844674407370955161 }
    slen := len(str)
    var out UInt128
    var carry uint64
    var i int
    for i=0; i<19 && i<slen && str[i]>='0' && str[i]<='9'; i++ {
        digit := byte(str[i])-'0'
        // multiply by 10
        out[0] *= 10
        // add digit
        out[0] += uint64(digit)
    }
    
    for ; i<slen && str[i]>='0' && str[i]<='9'; i++ {
        if out[1]>lastDigitValue[1] ||
            (out[1]==lastDigitValue[1] && out[0] > lastDigitValue[0]) {
            return UInt128{}, strconv.ErrRange
        }
        digit := byte(str[i])-'0'
        temp := out
        // multiply by 10
        out[1] = (temp[1]<<3) | (temp[0]>>61)
        out[0] = temp[0]<<3
        out[0], carry = Add64(out[0], temp[0]<<1, 0)
        out[1], _ = Add64(out[1], (temp[1]<<1) | (temp[0]>>63), carry)
        // add digit
        out[0], carry = Add64(out[0], uint64(digit), 0)
        out[1], carry = Add64(out[1], 0, carry)
        if carry!=0 {
            return UInt128{}, strconv.ErrRange
        }
    }
    if i==0 || i!=slen {
        return UInt128{}, strconv.ErrSyntax
    }
    return out, nil
}

// convert 128-unsigned integer to 64-bit float point value
func (a UInt128) ToFloat64() float64 {
    return float64(a[0]) + float64(a[1])*float64(18446744073709551616.0)
}

// convert 64-bit float point value to 128-bit unsigned integer
func Float64ToUInt128(a float64) (UInt128, error) {
    if math.IsNaN(a) || a >= 340282366920938463463374607431768211456.0 || a < 0.0 {
        return UInt128{}, strconv.ErrRange
    }
    am, ae := math.Frexp(a) // get binary mantisa and exponent
    if ae<0 { return UInt128{0,0}, nil }
    if ae<=64 { return UInt128{ uint64(a), 0 }, nil }
    ami := uint64(am * 18446744073709551616.0)
    ae -= 64
    // if last bit filled (shift is 64)
    if ae==64 { return UInt128{ 0, ami }, nil }
    return UInt128{ ami<<uint(ae), ami>>uint(64-ae) }, nil
}

// marshalling/unmarshaling

func (a UInt128) MarshalBinary() (data []byte, err error) {
    data2 := make([]byte, 16)
    binary.LittleEndian.PutUint64(data2[0:8], a[0])
    binary.LittleEndian.PutUint64(data2[8:16], a[1])
    return data2, nil
}

var ErrDataTooSmall error = errors.New("Data is too small")

func (a *UInt128) UnmarshalBinary(data []byte) error {
    if len(data) < 16 { return ErrDataTooSmall }
    a[0] = binary.LittleEndian.Uint64(data[0:8])
    a[1] = binary.LittleEndian.Uint64(data[8:16])
    return nil
}

func (a UInt128) MarshalText() (text []byte, err error) {
    return a.FormatBytes(), nil
}

func (a *UInt128) UnmarshalText(text []byte) error {
    var err error
    *a, err = ParseUInt128Bytes(text)
    return err
}


func (a UInt128) MarshalJSON() ([]byte, error) {
    if a[1]==0 {
        return []byte(a.Format()), nil
    }
    var sb bytes.Buffer
    sb.WriteRune('"')
    sb.WriteString(a.Format())
    sb.WriteRune('"')
    return sb.Bytes(), nil
}

func (a *UInt128) UnmarshalJSON(data []byte) error {
    dlen := len(data)
    var err error
    if dlen>=2 && (data[0]=='"'||data[0]=='\'') &&
                    (data[dlen-1]=='"'||data[dlen-1]=='\'') {
        *a, err = ParseUInt128(string(data[1:dlen-1]))
        return err
    }
    *a, err = ParseUInt128(string(data))
    return err
}
