/*
 * locale_test.go - tests for int128 routines
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

package goint128

import (
    "strconv"
    "testing"
)

type UInt128LocTC struct {
    lang string
    noSep1000 bool
    a UInt128
    expected string
}

func TestUInt128LocaleFormat(t *testing.T) {
    testCases := []UInt128LocTC {
        UInt128LocTC{ "af", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "am", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ar", false, UInt128{1234567890,0}, "١٬٢٣٤٬٥٦٧٬٨٩٠" },
        UInt128LocTC{ "az", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "bg", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "bn", false, UInt128{1234567890,0}, "১,২৩,৪৫,৬৭,৮৯০" },
        UInt128LocTC{ "ca", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "cs", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "da", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "de", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "el", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "en", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "es", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "et", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "fa", false, UInt128{1234567890,0}, "۱٬۲۳۴٬۵۶۷٬۸۹۰" },
        UInt128LocTC{ "fi", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "fil", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "fr", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "gu", false, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "he", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "hi", false, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "hr", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "hu", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "hy", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "id", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "is", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "it", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ja", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ka", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "kk", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "km", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "kn", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ko", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ky", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "lo", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "lt", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "lv", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "mk", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ml", false, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "mn", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "mo", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "mr", false, UInt128{1234567890,0}, "१,२३,४५,६७,८९०" },
        UInt128LocTC{ "ms", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "mul", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "my", false, UInt128{1234567890,0}, "၁,၂၃၄,၅၆၇,၈၉၀" },
        UInt128LocTC{ "nb", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "ne", false, UInt128{1234567890,0}, "१,२३४,५६७,८९०" },
        UInt128LocTC{ "nl", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "no", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "pa", false, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "pl", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pt", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ro", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ru", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sh", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "si", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "sk", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sl", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "sq", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sr", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "sv", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sw", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ta", false, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "te", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "th", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tl", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tn", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tr", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "uk", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "ur", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "uz", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "vi", false, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "zh", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "zu", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "en", false, UInt128{123456789,0}, "123,456,789" },
        UInt128LocTC{ "en", false, UInt128{23456789,0}, "23,456,789" },
        UInt128LocTC{ "en", false, UInt128{789,0}, "789" },
        UInt128LocTC{ "en", false, UInt128{89,0}, "89" },
        UInt128LocTC{ "ta", false, UInt128{234567890,0}, "23,45,67,890" },
        UInt128LocTC{ "ta", false, UInt128{34567890,0}, "3,45,67,890" },
        UInt128LocTC{ "ta", false, UInt128{4567890,0}, "45,67,890" },
        UInt128LocTC{ "ta", false, UInt128{890,0}, "890" },
        UInt128LocTC{ "ta", false, UInt128{89,0}, "89" },
        UInt128LocTC{ "ta", false, UInt128{9,0}, "9" },
        UInt128LocTC{ "", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "C", false, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "pl-PL", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pl_PL", false, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pl_PL.UTF-8", false, UInt128{1234567890,0}, "1 234 567 890" },
        // no separator 1000
        UInt128LocTC{ "af", true, UInt128{1234567890,0}, "1234567890" },
        UInt128LocTC{ "am", true, UInt128{1234567890,0}, "1234567890" },
        UInt128LocTC{ "ar", true, UInt128{1234567890,0}, "١٢٣٤٥٦٧٨٩٠" },
        UInt128LocTC{ "az", true, UInt128{1234567890,0}, "1234567890" },
        UInt128LocTC{ "bg", true, UInt128{1234567890,0}, "1234567890" },
        UInt128LocTC{ "bn", true, UInt128{1234567890,0}, "১২৩৪৫৬৭৮৯০" },
        UInt128LocTC{ "ca", true, UInt128{1234567890,0}, "1234567890" },
    }
    for i, tc := range testCases {
        a := tc.a
        result := tc.a.LocaleFormat(tc.lang, tc.noSep1000)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: fmt(%v,%s)->%v!=%v",
                     i, tc.a, tc.lang, tc.expected, result)
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d %s: %v!=%v", i, tc.lang, a, tc.a)
        }
        resultBytes := tc.a.LocaleFormatBytes(tc.lang, tc.noSep1000)
        if tc.expected!=string(resultBytes) {
            t.Errorf("Result mismatch: %d: fmtBytes(%v,%s)->%v!=%v",
                     i, tc.a, tc.lang, tc.expected, result)
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d %s: %v!=%v", i, tc.lang, a, tc.a)
        }
    }
}

type UInt128LocParseTC struct {
    lang string
    str string
    expected UInt128
    expError error
}

func TestUInt128LocaleParse(t *testing.T) {
    testCases := []UInt128LocParseTC {
        UInt128LocParseTC{ "en", "", UInt128{}, strconv.ErrSyntax },
        UInt128LocParseTC{ "en", "1,234,567,890", UInt128{1234567890,0}, nil },
        UInt128LocParseTC{ "en", "1234,567,890", UInt128{1234567890,0}, nil },
        UInt128LocParseTC{ "de", "1.234.567.890", UInt128{1234567890,0}, nil },
        UInt128LocParseTC{ "pl", "1 234 567 890", UInt128{1234567890,0}, nil },
        UInt128LocParseTC{ "pl", "1 234 567 890", UInt128{1234567890,0}, nil },
        UInt128LocParseTC{ "bn", "১,২৩,৪৫,৬৭,৮৯০", UInt128{1234567890,0}, nil },
        UInt128LocParseTC{ "bn", "1,234,567,890", UInt128{1234567890,0}, nil },
        UInt128LocParseTC{ "bn", "১,২৩,৪৫x৬৭,৮৯০", UInt128{}, strconv.ErrSyntax },
    }
    for i, tc := range testCases {
        result, err := LocaleParseUInt128(tc.lang, tc.str)
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: parse(%v,%v)->%v,%v!=%v,%v",
                     i, tc.lang, tc.str, tc.expected, tc.expError, result, err)
        }
        result, err = LocaleParseUInt128Bytes(tc.lang, []byte(tc.str))
        if tc.expected!=result || tc.expError!=err {
            t.Errorf("Result mismatch: %d: parseBytes(%v,%v)->%v,%v!=%v,%v",
                     i, tc.lang, tc.str, tc.expected, tc.expError, result, err)
        }
    }
}
