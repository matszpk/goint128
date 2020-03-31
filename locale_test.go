/*
 * locale_test.go - tests for int128 routines
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
    "strconv"
    "testing"
)

type UInt128LocTC struct {
    lang string
    sep1000 bool
    a UInt128
    expected string
}

func TestUInt128LocaleFormat(t *testing.T) {
    testCases := []UInt128LocTC {
        UInt128LocTC{ "af", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "am", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ar", true, UInt128{1234567890,0}, "١٬٢٣٤٬٥٦٧٬٨٩٠" },
        UInt128LocTC{ "az", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "bg", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "bn", true, UInt128{1234567890,0}, "১,২৩,৪৫,৬৭,৮৯০" },
        UInt128LocTC{ "ca", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "cs", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "da", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "de", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "el", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "en", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "es", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "et", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "fa", true, UInt128{1234567890,0}, "۱٬۲۳۴٬۵۶۷٬۸۹۰" },
        UInt128LocTC{ "fi", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "fil", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "fr", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "gu", true, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "he", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "hi", true, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "hr", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "hu", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "hy", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "id", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "is", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "it", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ja", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ka", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "kk", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "km", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "kn", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ko", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ky", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "lo", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "lt", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "lv", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "mk", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ml", true, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "mn", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "mo", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "mr", true, UInt128{1234567890,0}, "१,२३,४५,६७,८९०" },
        UInt128LocTC{ "ms", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "mul", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "my", true, UInt128{1234567890,0}, "၁,၂၃၄,၅၆၇,၈၉၀" },
        UInt128LocTC{ "nb", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "ne", true, UInt128{1234567890,0}, "१,२३४,५६७,८९०" },
        UInt128LocTC{ "nl", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "no", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "pa", true, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "pl", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pt", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ro", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ru", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sh", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "si", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "sk", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sl", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "sq", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sr", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "sv", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sw", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ta", true, UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "te", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "th", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tl", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tn", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tr", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "uk", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "ur", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "uz", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "vi", true, UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "zh", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "zu", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "en", true, UInt128{123456789,0}, "123,456,789" },
        UInt128LocTC{ "en", true, UInt128{23456789,0}, "23,456,789" },
        UInt128LocTC{ "en", true, UInt128{789,0}, "789" },
        UInt128LocTC{ "en", true, UInt128{89,0}, "89" },
        UInt128LocTC{ "ta", true, UInt128{234567890,0}, "23,45,67,890" },
        UInt128LocTC{ "ta", true, UInt128{34567890,0}, "3,45,67,890" },
        UInt128LocTC{ "ta", true, UInt128{4567890,0}, "45,67,890" },
        UInt128LocTC{ "ta", true, UInt128{890,0}, "890" },
        UInt128LocTC{ "ta", true, UInt128{89,0}, "89" },
        UInt128LocTC{ "ta", true, UInt128{9,0}, "9" },
        UInt128LocTC{ "", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "C", true, UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "pl-PL", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pl_PL", true, UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pl_PL.UTF-8", true, UInt128{1234567890,0}, "1 234 567 890" },
        // no separator 1000
        UInt128LocTC{ "af", false, UInt128{1234567890,0}, "1234567890" },
        UInt128LocTC{ "am", false, UInt128{1234567890,0}, "1234567890" },
        UInt128LocTC{ "ar", false, UInt128{1234567890,0}, "١٢٣٤٥٦٧٨٩٠" },
        UInt128LocTC{ "az", false, UInt128{1234567890,0}, "1234567890" },
        UInt128LocTC{ "bg", false, UInt128{1234567890,0}, "1234567890" },
        UInt128LocTC{ "bn", false, UInt128{1234567890,0}, "১২৩৪৫৬৭৮৯০" },
        UInt128LocTC{ "ca", false, UInt128{1234567890,0}, "1234567890" },
    }
    for i, tc := range testCases {
        a := tc.a
        result := tc.a.LocaleFormat(tc.lang, tc.sep1000)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: fmt(%v,%s)->%v!=%v",
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
    }
}
