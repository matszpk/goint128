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
    "testing"
)

type UInt128LocTC struct {
    lang string
    a UInt128
    expected string
}

func TestUInt128Locale(t *testing.T) {
    testCases := []UInt128LocTC {
        UInt128LocTC{ "af", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "am", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ar", UInt128{1234567890,0}, "١٬٢٣٤٬٥٦٧٬٨٩٠" },
        UInt128LocTC{ "az", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "bg", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "bn", UInt128{1234567890,0}, "১,২৩,৪৫,৬৭,৮৯০", },
        UInt128LocTC{ "ca", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "cs", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "da", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "de", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "el", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "en", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "es", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "et", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "fa", UInt128{1234567890,0}, "۱٬۲۳۴٬۵۶۷٬۸۹۰" },
        UInt128LocTC{ "fi", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "fil", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "fr", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "gu", UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "he", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "hi", UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "hr", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "hu", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "hy", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "id", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "is", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "it", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ja", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ka", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "kk", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "km", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "kn", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ko", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ky", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "lo", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "lt", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "lv", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "mk", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ml", UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "mn", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "mo", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "mr", UInt128{1234567890,0}, "१,२३,४५,६७,८९०" },
        UInt128LocTC{ "ms", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "mul", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "my", UInt128{1234567890,0}, "၁,၂၃၄,၅၆၇,၈၉၀" },
        UInt128LocTC{ "nb", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "ne", UInt128{1234567890,0}, "१,२३४,५६७,८९०" },
        UInt128LocTC{ "nl", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "no", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "pa", UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "pl", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pt", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ro", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "ru", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sh", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "si", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "sk", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sl", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "sq", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sr", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "sv", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "sw", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "ta", UInt128{1234567890,0}, "1,23,45,67,890" },
        UInt128LocTC{ "te", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "th", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tl", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tn", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "tr", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "uk", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "ur", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "uz", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "vi", UInt128{1234567890,0}, "1.234.567.890" },
        UInt128LocTC{ "zh", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "zu", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "en", UInt128{123456789,0}, "123,456,789" },
        UInt128LocTC{ "en", UInt128{23456789,0}, "23,456,789" },
        UInt128LocTC{ "en", UInt128{789,0}, "789" },
        UInt128LocTC{ "en", UInt128{89,0}, "89" },
        UInt128LocTC{ "ta", UInt128{234567890,0}, "23,45,67,890" },
        UInt128LocTC{ "ta", UInt128{34567890,0}, "3,45,67,890" },
        UInt128LocTC{ "ta", UInt128{4567890,0}, "45,67,890" },
        UInt128LocTC{ "ta", UInt128{890,0}, "890" },
        UInt128LocTC{ "ta", UInt128{89,0}, "89" },
        UInt128LocTC{ "ta", UInt128{9,0}, "9" },
        UInt128LocTC{ "", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "C", UInt128{1234567890,0}, "1,234,567,890" },
        UInt128LocTC{ "pl-PL", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pl_PL", UInt128{1234567890,0}, "1 234 567 890" },
        UInt128LocTC{ "pl_PL.UTF-8", UInt128{1234567890,0}, "1 234 567 890" },
    }
    for i, tc := range testCases {
        a := tc.a
        result := tc.a.Locale(tc.lang)
        if tc.expected!=result {
            t.Errorf("Result mismatch: %d: fmt(%v,%s)->%v!=%v",
                     i, tc.a, tc.lang, tc.expected, result)
        }
        if tc.a!=a {
            t.Errorf("Argument has been modified: %d %s: %v!=%v", i, tc.lang, a, tc.a)
        }
    }
}
