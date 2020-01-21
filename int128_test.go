/*
 * int128_test.go - tests for int128 routines
 *
 * goint128 - go int128 library
 * Copyright (C) 2019  Mateusz Szpakowski
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

package goint128

import (
    "bytes"
    "encoding/json"
    "fmt"
    "math"
    "strconv"
    "testing"
)

func getPanicInt2(f func(), paniced *bool, panicStr *string) {
    defer func() {
        if x:=recover(); x!=nil {
            *paniced = true
            *panicStr = fmt.Sprint(x)
        }
    }()
    f() // call
}

func getPanic2(f func()) (bool, string) {
    paniced := false
    panicStr := ""
    getPanicInt2(f, &paniced, &panicStr)
    return paniced, panicStr
}

type UInt128TC struct {
    a, b UInt128
    expected UInt128
}

func TestUInt128Add(t *testing.T) {
    testCases := []UInt128TC {
        UInt128TC{ UInt128{ 2454, 3421 }, UInt128{ 78731, 831 },
                UInt128{ 81185, 4252 } },
        UInt128TC{ UInt128{ 0xffffffffffff1001, 0x2442 }, UInt128{ 0xf003, 0xa8bc },
                UInt128{ 0x4, 0xccff } },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Add(tc.b)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: %v+%v->%v!=%v",
                     i, tc.a, tc.b, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

func TestUInt128Sub(t *testing.T) {
    testCases := []UInt128TC {
        UInt128TC{ UInt128{ 81185, 4252 }, UInt128{ 2454, 3421 },
                UInt128{ 78731, 831 } },
        UInt128TC{ UInt128{ 0x4, 0xccff }, UInt128{ 0xffffffffffff1001, 0x2442 },
                UInt128{ 0xf003, 0xa8bc } },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Sub(tc.b)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: %v-%v->%v!=%v",
                     i, tc.a, tc.b, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UInt128_64TC struct {
    a UInt128
    b uint64
    expected UInt128
}

func TestUInt128Add64(t *testing.T) {
    testCases := []UInt128_64TC {
        UInt128_64TC{ UInt128{ 3454, 3421 }, 78731, UInt128{ 82185, 3421 } },
        UInt128_64TC{ UInt128{ 0xffffffffffff1001, 0x2446 }, 0xf003,
                UInt128{ 0x4, 0x2447 } },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Add64(tc.b)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: %v+%v->%v!=%v",
                     i, tc.a, tc.b, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

func TestUInt128Sub64(t *testing.T) {
    testCases := []UInt128_64TC {
        UInt128_64TC{ UInt128{ 81185, 9165 }, 2454, UInt128{ 78731, 9165 } },
        UInt128_64TC{ UInt128{ 0x5, 0xccff }, 0xffffffffffff2001,
                UInt128{ 0xe004, 0xccfe } },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Sub64(tc.b)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: %v-%v->%v!=%v",
                     i, tc.a, tc.b, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UInt128CTC struct {
    a, b UInt128
    c uint64
    expected UInt128
    expC uint64
}

func TestUInt128AddC(t *testing.T) {
    testCases := []UInt128CTC {
        UInt128CTC { UInt128{ 8481, 7754 }, UInt128{ 1121, 5531 }, 0,
            UInt128{ 9602, 13285 }, 0 },
        UInt128CTC { UInt128{ 8481, 7754 }, UInt128{ 1121, 5531 }, 1,
            UInt128{ 9603, 13285 }, 0 },
        UInt128CTC { UInt128{ 0xfffffffffffffffe, 7754 }, UInt128{ 1, 5531 }, 1,
            UInt128{ 0, 13286 }, 0 },
        UInt128CTC { UInt128{ 0xfffffffffffffffd, 7754 }, UInt128{ 1, 5531 }, 1,
            UInt128{ 0xffffffffffffffff, 13285 }, 0 },
        UInt128CTC { UInt128{ 0xffffffffffffff22, 0xfffffffffffffffe },
            UInt128{ 0xde, 1 }, 0, UInt128{ 0, 0 }, 1 },
        UInt128CTC { UInt128{ 0xffffffffffffff25, 0xfffffffffffffffe },
            UInt128{ 0xde, 2 }, 0, UInt128{ 3, 1 }, 1 },
        UInt128CTC { UInt128{ 0xffffffffffffff25, 0xfffffffffffffffe },
            UInt128{ 0xd1, 3 }, 0, UInt128{ 0xfffffffffffffff6, 1 }, 1 },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result, resultC := tc.a.AddC(tc.b, tc.c)
        if tc.expected!=result || tc.expC!=resultC {
            t.Errorf("Result mismatch: %d: addc(%v,%v,%v)->%v,%v!=%v,%v",
                     i, tc.a, tc.b, tc.c, tc.expected, tc.expC, result, resultC)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

func TestUInt128SubB(t *testing.T) {
    testCases := []UInt128CTC {
        UInt128CTC{ UInt128{ 81183, 4252 }, UInt128{ 2454, 3421 }, 0,
                UInt128{ 78729, 831 }, 0 },
        UInt128CTC{ UInt128{ 81185, 4252 }, UInt128{ 2454, 3421 }, 1,
                UInt128{ 78730, 831 }, 0 },
        UInt128CTC{ UInt128{ 0x4, 0xccff }, UInt128{ 0xffffffffffff1001, 0x2442 }, 1,
                UInt128{ 0xf002, 0xa8bc }, 0 },
        UInt128CTC{ UInt128{ 81185, 4252 }, UInt128{ 81183, 4253 }, 0,
                UInt128{ 2 , 0xffffffffffffffff }, 1 },
        UInt128CTC{ UInt128{ 81185, 4252 }, UInt128{ 81187, 4253 }, 0,
                UInt128{ 0xfffffffffffffffe, 0xfffffffffffffffe }, 1 },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result, resultC := tc.a.SubB(tc.b, tc.c)
        if tc.expected!=result || tc.expC!=resultC {
            t.Errorf("Result mismatch: %d: subb(%v,%v,%v)->%v,%v!=%v,%v",
                     i, tc.a, tc.b, tc.c, tc.expected, tc.expC, result, resultC)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UInt128CmpTC struct {
    a, b UInt128
    expected int
}

func TestUInt128Cmp(t *testing.T) {
    testCases := []UInt128CmpTC {
        UInt128CmpTC{ UInt128{ 3421, 2454 }, UInt128{ 831, 78731 }, -1 },
        UInt128CmpTC{ UInt128{ 6743, 6841 }, UInt128{ 7731121, 1212 }, 1 },
        UInt128CmpTC{ UInt128{ 1821, 33411 }, UInt128{ 589759892, 33411 }, -1 },
        UInt128CmpTC{ UInt128{ 5788219381, 33411 }, UInt128{ 954891, 33411 }, 1 },
        UInt128CmpTC{ UInt128{ 1231, 33411 }, UInt128{ 1231, 33411 }, 0 },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Cmp(tc.b)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: cmp(%v,%v)->%v!=%v",
                     i, tc.a, tc.b, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

func TestUInt128Mul(t *testing.T) {
    testCases := []UInt128TC {
        UInt128TC { UInt128{ 0xc9baa109a40baa11, 0x384b9a928941ac3 },
            UInt128{ 0x1839b9af9dc021, 0x49310ace3a1a15 },
            UInt128{ 0x6ac740f8d07aac31, 0x2fe36adfd8d92a0e } },
        UInt128TC { UInt128{ 0xffffffffffffffff, 0xffffffffffffffff },
            UInt128{ 0xfffffffffffffffd, 0xffffffffffffffff },
            UInt128{ 3, 0 } },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Mul(tc.b)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: %v*%v->%v!=%v",
                     i, tc.a, tc.b, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UInt128MulFTC struct {
    a, b UInt128
    expectedLo, expectedHi UInt128
}

func TestUInt128MulFull(t *testing.T) {
    testCases := []UInt128MulFTC {
        UInt128MulFTC { UInt128{ 0xa0a59e0cd5640249, 0x5ff18c5e354dd456 },
            UInt128{ 0x4ddec0edfcc8c414, 0xadf9e6b9046f6ea3 },
            UInt128{ 0xf5a23257e29811b4, 0x89c07fdabef4588c },
            UInt128{ 0xd8d0c5c68299cf33, 0x4133e4458cfc0e8e } },
        UInt128MulFTC { UInt128{ 0xffffffffffffffff, 0xffffffffffffffff },
            UInt128{ 0xfffffffffffffffd, 0xffffffffffffffff },
            UInt128{ 3, 0 },
            UInt128{ 0xfffffffffffffffc, 0xffffffffffffffff } },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result, resultLo := tc.a.MulFull(tc.b)
        if tc.expectedHi!=result || tc.expectedLo!=resultLo {
            t.Errorf("Result mismatch: %d: mulfull(%v,%v)->%v,%v!=%v,%v",
                     i, tc.a, tc.b, tc.expectedLo, tc.expectedHi, resultLo, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UInt128ShTC struct {
    a UInt128
    b uint
    expected UInt128
}

func TestUInt128Shl(t *testing.T) {
    testCases := []UInt128ShTC {
        UInt128ShTC { UInt128{ 0x62b71430f1765e8f, 0xb5ed145b3920ca5a }, 3,
            UInt128{ 0x15b8a1878bb2f478, 0xaf68a2d9c90652d3 } },
        UInt128ShTC { UInt128{ 0x62b71430f1765e8f, 0xb5ed145b3920ca5a }, 11,
            UInt128{ 0xb8a1878bb2f47800, 0x68a2d9c90652d315 } },
        UInt128ShTC { UInt128{ 0x62b7ac5532325e8f, 0xc5ed145b3920ca5a }, 0,
            UInt128{ 0x62b7ac5532325e8f, 0xc5ed145b3920ca5a } },
        UInt128ShTC { UInt128{ 0xf621e52aaa8b880c, 0xb4283ce0fd8464e2 }, 73,
            UInt128{ 0, 0x43ca555517101800 } },
        UInt128ShTC { UInt128{ 0xf621e52aaa8b880c, 0xb4283ce0fd8464e2 }, 64,
            UInt128{ 0, 0xf621e52aaa8b880c } },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Shl(tc.b)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d:%v<<%v->%v!=%v",
                     i, tc.a, tc.b, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

func TestUInt128Shr(t *testing.T) {
    testCases := []UInt128ShTC {
        UInt128ShTC { UInt128{ 0xeebbd1b847efeefa, 0x1f27b7128996878e }, 11,
            UInt128{ 0xf1ddd77a3708fdfd, 0x3e4f6e25132d0 } },
        UInt128ShTC { UInt128{ 0xecabd1b847efe63a, 0x1f27b7523196878f }, 0,
            UInt128{ 0xecabd1b847efe63a, 0x1f27b7523196878f } },
        UInt128ShTC { UInt128{ 0xf4f393b4762c797a, 0x51c18de532f49530 }, 82,
            UInt128{ 0x147063794cbd, 0 } },
        UInt128ShTC { UInt128{ 0xadd45555288f694c, 0x2b2e0d6f95ff2df1 }, 64,
            UInt128{ 0x2b2e0d6f95ff2df1, 0 } },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result := tc.a.Shr(tc.b)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: %v>>%v->%v!=%v",
                     i, tc.a, tc.b, tc.expected, result)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UInt128DivTC struct {
    a UInt128
    b uint64
    expected UInt128
    expRem uint64
}

func TestUInt128Div(t *testing.T) {
    testCases := []UInt128DivTC {
        UInt128DivTC { UInt128{ 0x0d362b7e0421d339, 0xbb09d477baa0 },
            0x6afcb5c6af1e507b, UInt128{ 492083670228144, 0 }, 0x13f254e3d9ce0aa9 },
        UInt128DivTC { UInt128{ 0x0bc4f2ea7ec06c3f, 0x7bdcd02be78fe },
            0x3e2dc3dd417, UInt128{ 0xf6491fcb9513612d, 0x1fd }, 0x25139d06d34 },
        UInt128DivTC { UInt128{ 0, 1<<55 }, 1<<55, UInt128{ 0, 1 }, 0 },
        UInt128DivTC { UInt128{ 0, 1<<62 }, 1<<55, UInt128{ 0, 128 }, 0 },
        UInt128DivTC { UInt128{ 0x2e9700d1e595b258, 0x34a67968e5a },
            0xc23b96121, UInt128{ 0x64b6c9b6ee122e0c, 0x45 }, 0x9671f36cc },
        UInt128DivTC { UInt128{ 55, 0 }, 7, UInt128{ 7, 0 }, 6 },
        // no remainder
        UInt128DivTC { UInt128{ 0x0f2b92f72757046a, 0x15b807b7564a },
            0x26b380a13ca, UInt128{ 0xfaa679c50cd8d211, 0x8 }, 0 },
    }
    for i, tc := range testCases {
        a, b := tc.a, tc.b
        result, rem := tc.a.Div64(tc.b)
        if tc.expected!=result || tc.expRem!=rem {
            t.Errorf("Result mismatch: %d: %v/%v->%v,%v!=%v,%v",
                     i, tc.a, tc.b, tc.expected, tc.expRem, result, rem)
        }
        if tc.a!=a || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v!=%v,%v",
                     i, a, b, tc.a, tc.b)
        }
    }
}

type UInt128DivFTC struct {
    alo, ahi UInt128
    b UInt128
    expected, expRem UInt128
}

func TestUInt128DivFull(t *testing.T) {
    testCases := []UInt128DivFTC {
        UInt128DivFTC{ UInt128{ 0xa168b431ea4cbf25, 0xeeaf8afeafe15bf3 }, // alo
            UInt128{ 0x79da7cfc64734fc8, 0x1ae093566b591f }, // ahi
            UInt128{ 0x64611073ad67885c, 0x159b7addc721d10f }, // b
            UInt128{ 0x71e0ef1e6710ea31, 0x13e6fd8cef95977 }, // quo
            UInt128{ 0x38db746f8d178d89, 0x1011ed7a4d743993 } }, // rem
        UInt128DivFTC{ UInt128{ 0xc9a7d6e2cc4a9fe1, 0x7c5f7c4fe1dd3975 }, // alo
            UInt128{ 0x78c86ab5339b57fc, 0xaa9ea603a6ff1 }, // ahi
            UInt128{ 0x8d4959f4e6d39704, 0x17b4ad5d2b7537 }, // b
            UInt128{ 0x106a3b20e0f77e82, 0x73288478235baedf }, // quo
            UInt128{ 0x74f4f81f3ba7f7d9, 0x24c773dc419e1 } }, // rem
        UInt128DivFTC{ UInt128{ 0xad1b0bef418b04f3, 0xad386b96ec18a75d }, // alo
            UInt128{ 0x3c179a833f04, 0 }, // ahi
            UInt128{ 0x448ab60d06e16d71, 0x21277fb3c975915 }, // b
            UInt128{ 0x1d00017916c509, 0 }, // quo
            UInt128{ 0x4c150da3b9b036fa, 0x1030338b9fb3651 } }, // rem
        UInt128DivFTC{ UInt128{ 0xfe846594f784bcc1, 0xf3abd28b98484862 }, // alo
            UInt128{ 0xd3e91d7d4a, 0 }, // ahi
            UInt128{ 0x1725a5b765d6df45, 0x251135 }, // b
            UInt128{ 0x978eaa37efa35277, 0x5b788 }, // quo
            UInt128{ 0x546262a1392fd9ae, 0x17d60b } }, // rem
        UInt128DivFTC{ UInt128{ 0x17575839531cc261, 0x876500912715e24f }, // alo
            UInt128{ 0x3832d66fa89b0, 0 }, // ahi
            UInt128{ 0xbb13d0419ee95154, 0x1ef6c6ca9f102 }, // b
            UInt128{ 0xd0a1ad051eb58b86, 1 }, // quo
            UInt128{ 0x4d87a7751d6f9469, 0xf20352ff6a13 } }, // rem
        UInt128DivFTC{ UInt128{ 0xb04916027d7360fd, 0xbdd6fc093b36eef0 }, // alo
            UInt128{ 0x1e2eb3c64254a, 0 }, // ahi
            UInt128{ 0x8076a6a122255eb2, 0x93b26bfc783ba6 }, // b
            UInt128{ 0x34508756e87ad76, 0 }, // quo
            UInt128{ 0xcd547bea135d70f1, 0x8f7c9163375c51 } }, // rem
        // lower a and b
        UInt128DivFTC { UInt128{ 0x0d362b7e0421d339, 0xbb09d477baa0 }, UInt128{}, // a
            UInt128{ 0x6afcb5c6af1e507b, 0 }, // b
            UInt128{ 492083670228144, 0 }, // quo
            UInt128 {0x13f254e3d9ce0aa9, 0 } }, // rem
        UInt128DivFTC { UInt128{ 0x0bc4f2ea7ec06c3f, 0x7bdcd02be78fe }, UInt128{}, // a
            UInt128{ 0x3e2dc3dd417, 0 }, // b
            UInt128{ 0xf6491fcb9513612d, 0x1fd }, // quo
            UInt128{ 0x25139d06d34 } }, // rem
        UInt128DivFTC { UInt128{ 58, 0 }, UInt128{}, UInt128{ 7, 0 },
            UInt128{ 8, 0 }, UInt128{ 2, 0 } },
        // no remainder
        UInt128DivFTC { UInt128{ 0xf023facc617c5db4, 0xe5a87c07bf5a5a69 }, // alo
            UInt128{ 0xaf5996526c0426de, 0x5f468b14014b }, // ahi
            UInt128{ 0x9523b1e7742f2017, 0x1b2e5c6b574ad598 }, // b
            UInt128{ 0xeacaf09f790c4c6c, 0x38155b1981fb0 }, UInt128{} }, //quo,rem
        // full remainder and max quotient
        UInt128DivFTC { UInt128{ 0xffffffffffffffff, 0xffffffffffffffff }, // alo
            UInt128{ 0x54cd83b46f259de8, 0x213a9ec7 }, // ahi
            UInt128{ 0x54cd83b46f259de9, 0x213a9ec7 }, // b
            UInt128{ 0xffffffffffffffff, 0xffffffffffffffff }, // quo
            UInt128{ 0x54cd83b46f259de8, 0x213a9ec7 } }, // rem
        // smaller a
        UInt128DivFTC { UInt128{ 0xc1e79b199458a88a, 0x38f41ebf9d94b }, UInt128{}, // a
            UInt128{ 0x9cc0cb116cd60d5e, 0x1051f1062 }, // b
            UInt128{ 0x37d62, 0 }, UInt128{ 0x3bbe15c73dc6a48e, 0x99ed96bf } }, // quo,rem
        UInt128DivFTC {
            UInt128{ 0xcc8a934a9b390141, 0xd8a91058bc8f94ae }, UInt128{}, // a
            UInt128{ 0xcc8a934a9b39013f, 0xd8a91058bc8f94ae }, // b
            UInt128{ 1, 0 }, UInt128{ 2, 0 } }, // quo,rem
        UInt128DivFTC {
            UInt128{ 0xcc8a934a9b390141, 0xd8a91058bc8f94ae }, UInt128{}, // a
            UInt128{ 0xcc8a934a9b390144, 0xd8a91058bc8f94ae }, // b
            UInt128{}, UInt128{ 0xcc8a934a9b390141, 0xd8a91058bc8f94ae } }, // quo,rem
    }
    for i, tc := range testCases {
        alo, ahi, b := tc.alo, tc.ahi, tc.b
        result, rem := UInt128DivFull(tc.ahi, tc.alo, tc.b)
        if tc.expected!=result || tc.expRem!=rem {
            t.Errorf("Result mismatch: %d: (%v,%v)/%v->%v,%v!=%v,%v",
                     i, tc.alo, tc.ahi, tc.b, tc.expected, tc.expRem, result, rem)
        }
        if tc.alo!=alo || tc.ahi!=ahi || tc.b!=b {
            t.Errorf("Argument has been modified: %d: %v,%v,%v!=%v,%v,%v",
                     i, alo, ahi, b, tc.alo, tc.ahi, tc.b)
        }
    }
    
    paniced, panicStr := getPanic2(func() {
        UInt128DivFull(UInt128{ 0x54cd83b46f259de9, 0x213a9ec7 }, UInt128{},
                       UInt128{ 0x54cd83b46f259de9, 0x213a9ec7 })
    })
    if !paniced || panicStr!="Divide overflow" {
        t.Errorf("Unexpected panic: %v,%v", paniced, panicStr)
    }
    paniced, panicStr = getPanic2(func() {
        UInt128DivFull(UInt128{ 0x54cd834632566de9, 0x213a9ec7545 }, UInt128{},
                       UInt128{ 0x54cd83111f259663, 0x213a9ec7 })
    })
    if !paniced || panicStr!="Divide overflow" {
        t.Errorf("Unexpected panic: %v,%v", paniced, panicStr)
    }
    paniced, panicStr = getPanic2(func() {
        UInt128DivFull(UInt128{ 0x54cd834632566de9, 0x213a9ec7545 }, UInt128{},
                       UInt128{})
    })
    if !paniced || panicStr!="Divide by zero" {
        t.Errorf("Unexpected panic: %v,%v", paniced, panicStr)
    }
}

type UInt128FmtTC struct {
    a UInt128
    expected string
}

func TestUInt128Format(t *testing.T) {
    testCases := []UInt128FmtTC {
        UInt128FmtTC { UInt128{ 0x5f75348b0131b3af, 0xb3af0f },
            "217224419425143693331510191" },
        UInt128FmtTC { UInt128{ 834899285198348317, 0 }, "834899285198348317" },
        UInt128FmtTC { UInt128{ 0xffffffffffffffff, 0xffffffffffffffff },
            "340282366920938463463374607431768211455" },
        UInt128FmtTC { UInt128{3875820019684212736, 54}, "1000000000000000000000" },
        UInt128FmtTC { UInt128{3875820019684212737, 54}, "1000000000000000000001" },
        UInt128FmtTC { UInt128{3875820019684212735, 54}, "999999999999999999999" },
        UInt128FmtTC { UInt128{ 0, 0 }, "0" },
        UInt128FmtTC { UInt128{ 1, 0 }, "1" },
        UInt128FmtTC { UInt128{ 9, 0 }, "9" },
        UInt128FmtTC { UInt128{ 0x7b6cf6eef49b5f2d, 0x3a6e20ada49e19 },
            "303386871892539280136352169180487469" },
        UInt128FmtTC { UInt128{ 2856728644567338944, 0 }, "2856728644567338944" },
        UInt128FmtTC { UInt128{ 0, 1 }, "18446744073709551616" },
        UInt128FmtTC { UInt128{ 2, 1 }, "18446744073709551618" },
        UInt128FmtTC { UInt128{ 4, 1 }, "18446744073709551620" },
        UInt128FmtTC { UInt128{ 0xfffffffffffffffe, 0 }, "18446744073709551614" },
        UInt128FmtTC { UInt128{ 130994, 0 }, "130994" },
        UInt128FmtTC { UInt128{ 0x8c261ad7409395f0, 0x47a96e },
            "86633852368880475724551664" },
        UInt128FmtTC { UInt128{ 0x1c9e66c000000000, 0xe1b1e5f90f944d6e },
            "300000000000000000000000000000000000000" },
        UInt128FmtTC { UInt128{ 0x1c9e66bfffffffff, 0xe1b1e5f90f944d6e },
            "299999999999999999999999999999999999999" },
        UInt128FmtTC { UInt128{1000000000000000, 0}, "1000000000000000" },
    }
    for i, tc := range testCases {
        a := tc.a
        result := tc.a.Format()
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: fmt(%v)->%v!=%v",
                     i, tc.a, tc.expected, result)
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d: %v!=%v", i, a, tc.a)
        }
    }
}

type UInt128ParseTC struct {
    str string
    expected UInt128
    expError error
}

func TestUInt128Parse(t *testing.T) {
    testCases := []UInt128ParseTC {
        UInt128ParseTC { "217224419425143693331510191",
            UInt128{ 0x5f75348b0131b3af, 0xb3af0f }, nil },
        UInt128ParseTC { "834899285198348317", UInt128{ 834899285198348317, 0 }, nil },
        UInt128ParseTC { "340282366920938463463374607431768211455",
            UInt128{ 0xffffffffffffffff, 0xffffffffffffffff }, nil },
        UInt128ParseTC { "1000000000000000000000",
            UInt128{3875820019684212736, 54}, nil },
        UInt128ParseTC { "1000000000000000000001",
            UInt128{3875820019684212737, 54}, nil },
        UInt128ParseTC { "999999999999999999999",
            UInt128{3875820019684212735, 54}, nil },
        UInt128ParseTC { "0", UInt128{ 0, 0 }, nil },
        UInt128ParseTC { "1", UInt128{ 1, 0 }, nil },
        UInt128ParseTC { "9", UInt128{ 9, 0 }, nil },
        UInt128ParseTC { "303386871892539280136352169180487469",
            UInt128{ 0x7b6cf6eef49b5f2d, 0x3a6e20ada49e19 }, nil },
        UInt128ParseTC { "2856728644567338944", UInt128{ 2856728644567338944, 0 }, nil },
        UInt128ParseTC { "18446744073709551616", UInt128{ 0, 1 }, nil },
        UInt128ParseTC { "18446744073709551618", UInt128{ 2, 1 }, nil },
        UInt128ParseTC { "18446744073709551614", UInt128{ 0xfffffffffffffffe, 0 }, nil },
        UInt128ParseTC { "130994", UInt128{ 130994, 0 }, nil },
        UInt128ParseTC { "300000000000000000000000000000000000000",
            UInt128{ 0x1c9e66c000000000, 0xe1b1e5f90f944d6e }, nil },
        UInt128ParseTC { "299999999999999999999999999999999999999",
            UInt128{ 0x1c9e66bfffffffff, 0xe1b1e5f90f944d6e }, nil },
        UInt128ParseTC { "1000000000000000", UInt128{1000000000000000, 0}, nil },
        UInt128ParseTC { "340282366920938463463374607431768211456",
            UInt128{}, strconv.ErrRange },
        UInt128ParseTC { "555892892181893785497579348923892811124",
            UInt128{}, strconv.ErrRange },
        UInt128ParseTC { "1558928921818937854975793489238928111248",
            UInt128{}, strconv.ErrRange },
        UInt128ParseTC { "", UInt128{}, strconv.ErrSyntax },
        UInt128ParseTC { "342xx", UInt128{}, strconv.ErrSyntax },
    }
    for i, tc := range testCases {
        result, err := ParseUInt128(tc.str)
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: parse(%v)->%v,%v!=%v,%v",
                     i, tc.str, tc.expected, tc.expError, result, err)
        }
    }
}

type UInt128ToFloat64TC struct {
    value UInt128
    expected float64
}

func TestUInt128ToFloat64(t *testing.T) {
    testCases := []UInt128ToFloat64TC{
        UInt128ToFloat64TC{ UInt128{ 0, 0 }, 0.0 },
        UInt128ToFloat64TC{ UInt128{ 1, 0 }, 1.0 },
        UInt128ToFloat64TC{ UInt128{ 54930201, 0 }, 54930201.0 },
        UInt128ToFloat64TC{ UInt128{ 85959028918918968, 0 }, 85959028918918968.0 },
        UInt128ToFloat64TC{ UInt128{ 16346246572275455745, 10277688839402 },
                    189589895689685989335661129029377.0 },
        UInt128ToFloat64TC{ UInt128{ 0xffffffffffffffff, 0xffffffffffffffff },
                340282366920938463463374607431768211455.0 },
    }
    for i, tc := range testCases {
        result := tc.value.ToFloat64()
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: tofloat64(%v)->%v!=%v",
                     i, tc.value, tc.expected, result)
        }
    }
}

type Float64ToUInt128TC struct {
    value float64
    expected UInt128
    expError error
}

func TestFloat64ToUInt128(t *testing.T) {
    testCases := []Float64ToUInt128TC{
        Float64ToUInt128TC{ 0.0, UInt128{ 0, 0 }, nil },
        Float64ToUInt128TC{ 1.0, UInt128{ 1, 0 }, nil },
        Float64ToUInt128TC{ 1.7, UInt128{ 1, 0 }, nil },
        Float64ToUInt128TC{ 145645677.18, UInt128{ 145645677, 0 }, nil },
        Float64ToUInt128TC{ 3145645677.778, UInt128{ 3145645677, 0 }, nil },
        Float64ToUInt128TC{ 187923786919586921.0,
            UInt128{ 187923786919586912, 0 }, nil },
        Float64ToUInt128TC{ 11792378691958692154.0,
            UInt128{ 11792378691958691840, 0 }, nil },
        Float64ToUInt128TC{ 26858969188828978177.0,
            UInt128{ 8412225115119427584, 1 }, nil },
        Float64ToUInt128TC{ 75901828515489894894398.0,
            UInt128{ 11923396248800854016, 4114 }, nil },
        Float64ToUInt128TC{ 93895689491189486962895905237.0,
            UInt128{ 12333477015461036032, 5090095526 }, nil },
        Float64ToUInt128TC{ 8892238191913586823938589549295667.0,
            UInt128{ 4611686018427387904, 482049198296564 }, nil },
        Float64ToUInt128TC{ 2938968690290190390494904909429029029.0,
            UInt128{ 0, 159321811943975104 }, nil },  // ?
        Float64ToUInt128TC{ 219849568662195967795923292939493492190.0,
            UInt128{ 0, 11918068998177697792 }, nil },  // ?
        Float64ToUInt128TC{ 340282366920938463463374607431768211456.0,
            UInt128{}, strconv.ErrRange },
        Float64ToUInt128TC{ -0.5, UInt128{}, strconv.ErrRange },
        Float64ToUInt128TC{ math.Inf(1), UInt128{}, strconv.ErrRange },
        Float64ToUInt128TC{ math.NaN(), UInt128{}, strconv.ErrRange },
    }
    for i, tc := range testCases {
        result, err := Float64ToUInt128(tc.value)
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: touint128(%v)->%v,%v!=%v,%v",
                     i, tc.value, tc.expected, tc.expError, result, err)
        }
    }
}

type UInt128MarshalBinTC struct {
    value UInt128
    expected []byte
}

func TestUInt128MarshalBinary(t *testing.T) {
    testCases := []UInt128MarshalBinTC{
        UInt128MarshalBinTC{ UInt128{ 0xccaa010203040506, 0xbbaca34c0a04521 },
                []byte{ 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0xaa, 0xcc,
                    0x21, 0x45, 0xa0, 0xc0, 0x34, 0xca, 0xba, 0xb } },
    }
    for i, tc := range testCases {
        result, err := tc.value.MarshalBinary()
        if err!=nil {
            t.Errorf("MarshalBinary returns error: %v", err)
        }
        if !bytes.Equal(tc.expected, result) {
            t.Errorf("Result mismatch: %d: marshalbin(%v)->%v!=%v",
                     i, tc.value, tc.expected, result)
        }
    }
}

type UInt128UnmarshalBinTC struct {
    data []byte
    expected UInt128
    expError error
}

func TestUInt128UnmarshalBinary(t *testing.T) {
    testCases := []UInt128UnmarshalBinTC{
        UInt128UnmarshalBinTC{ 
            []byte{ 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0xaa, 0xcc,
                    0x21, 0x45, 0xa0, 0xc0, 0x34, 0xca, 0xba, 0xb },
            UInt128{ 0xccaa010203040506, 0xbbaca34c0a04521 }, nil },
        UInt128UnmarshalBinTC{ 
            []byte{ 0x06, 0x05, 0x04, 0x03, 0x02, 0x01, 0xaa, 0xcc,
                    0x21, 0x45, 0xa0, 0xc0, 0x34, 0xca, 0xba },
            UInt128{}, ErrDataTooSmall },
    }
    for i, tc := range testCases {
        var v UInt128
        err := v.UnmarshalBinary(tc.data)
        if tc.expected!=v || tc.expError!=err {
            t.Errorf("Result mismatch: %d: unmarshalbin(%v)->%v,%v!=%v,%v",
                     i, tc.data, tc.expected, tc.expError, v, err)
        }
    }
}

func TestUInt128MarshalText(t *testing.T) {
    testCases := []UInt128MarshalBinTC{
        UInt128MarshalBinTC{ UInt128{ 34954975929367788, 0 },
                []byte("34954975929367788") },
        UInt128MarshalBinTC{ UInt128{ 1492718235287466483, 42196924 },
                []byte("778395859218490582901895667") },
    }
    for i, tc := range testCases {
        result, err := tc.value.MarshalText()
        if err!=nil {
            t.Errorf("MarshalText returns error: %v", err)
        }
        if !bytes.Equal(tc.expected, result) {
            t.Errorf("Result mismatch: %d: marshaltext(%v)->%v!=%v",
                     i, tc.value, tc.expected, result)
        }
    }
}

func TestUInt128UnmarshalText(t *testing.T) {
    testCases := []UInt128UnmarshalBinTC{
        UInt128UnmarshalBinTC{ []byte("34954975929367788"),
                UInt128{ 34954975929367788, 0 }, nil },
        UInt128UnmarshalBinTC{ []byte("778395859218490582901895667"),
                UInt128{ 1492718235287466483, 42196924 }, nil },
        UInt128UnmarshalBinTC{ []byte("778395859218490582901895667xxx"),
                UInt128{}, strconv.ErrSyntax },
    }
    for i, tc := range testCases {
        var v UInt128
        err := v.UnmarshalText(tc.data)
        if tc.expected!=v || tc.expError!=err {
            t.Errorf("Result mismatch: %d: unmarshaltext(%v)->%v,%v!=%v,%v",
                     i, tc.data, tc.expected, tc.expError, v, err)
        }
    }
}

func TestUInt128MarshalJSON(t *testing.T) {
    testCases := []UInt128MarshalBinTC{
        UInt128MarshalBinTC{ UInt128{ 34954975929367788, 0 },
                []byte("34954975929367788") },
        UInt128MarshalBinTC{ UInt128{ 1492718235287466483, 42196924 },
                []byte("\"778395859218490582901895667\"") },
    }
    for i, tc := range testCases {
        result, err := tc.value.MarshalJSON()
        if err!=nil {
            t.Errorf("MarshalJSON returns error: %v", err)
        }
        if !bytes.Equal(tc.expected, result) {
            t.Errorf("Result mismatch: %d: marshaljson(%v)->%v!=%v",
                     i, tc.value, tc.expected, result)
        }
    }
}

func TestUInt128UnmarshalJSON(t *testing.T) {
    testCases := []UInt128UnmarshalBinTC{
        UInt128UnmarshalBinTC{ []byte("34954975929367788"),
                UInt128{ 34954975929367788, 0 }, nil },
        UInt128UnmarshalBinTC{ []byte("\"34954975929367788\""),
                UInt128{ 34954975929367788, 0 }, nil },
        UInt128UnmarshalBinTC{ []byte("'34954975929367788'"),
                UInt128{ 34954975929367788, 0 }, nil },
        UInt128UnmarshalBinTC{ []byte("\"778395859218490582901895667\""),
                UInt128{ 1492718235287466483, 42196924 }, nil },
        UInt128UnmarshalBinTC{ []byte("'778395859218490582901895667'"),
                UInt128{ 1492718235287466483, 42196924 }, nil },
        UInt128UnmarshalBinTC{ []byte("778395859218490582901895667xxx"),
                UInt128{}, strconv.ErrSyntax },
    }
    for i, tc := range testCases {
        var v UInt128
        err := v.UnmarshalJSON(tc.data)
        if tc.expected!=v || tc.expError!=err {
            t.Errorf("Result mismatch: %d: unmarshaljson(%v)->%v,%v!=%v,%v",
                     i, tc.data, tc.expected, tc.expError, v, err)
        }
    }
}

type SampleStruct struct {
    A, B UInt128
}

func TestUInt128JSONHandling(t *testing.T) {
    const sampleText = `{ "A": 134554, "B": "234499215868989382112354567" }`
    var out SampleStruct
    var err error
    if err = json.Unmarshal([]byte(sampleText), &out); err!=nil {
        t.Errorf("Unmarshal returns error: %v", err)
    }
    expected := SampleStruct{ UInt128{134554, 0}, UInt128{17793088829901545735, 12712227} }
    if out!=expected {
        t.Errorf("Result mismatch: %v", out)
    }
    var b []byte
    b, err = json.Marshal(out)
    if string(b)!=`{"A":134554,"B":"234499215868989382112354567"}` {
        t.Errorf("Result mismatch: %v", string(b))
    }
}
