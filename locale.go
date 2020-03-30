/*
 * locale.go - locale routines
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

// Package to operate 128-bit integer
package goint128

import "strconv"

// locale formatting info
type LocFmt struct {
    Comma, Sep1000 rune
    Sep100and1000 bool
    Digits []rune
}

var normalDigits []rune = []rune("0123456789")
var arDigits []rune = []rune("٠١٢٣٤٥٦٧٨٩")
var faDigits []rune = []rune("۰۱۲۳۴۵۶۷۸۹")
var bnDigits []rune = []rune("০১২৩৪৫৬৭৮৯")
var mrDigits []rune = []rune("०१२३४५६७८९")
var myDigits []rune = []rune("၀၁၂၃၄၅၆၇၈၉")

var defaultLocaleFormat LocFmt = LocFmt{ '.', ',', false, normalDigits }

var localeFormats map[string]LocFmt = map[string]LocFmt {
    "af": LocFmt{ ',', ' ', false, normalDigits },
    "am": LocFmt{ '.', ',', false, normalDigits },
    "ar": LocFmt{ '٫', '٬', false, arDigits },
    "az": LocFmt{ ',', '.', false, normalDigits },
    "bg": LocFmt{ ',', ' ', false, normalDigits },
    "bn": LocFmt{ '.', ',', true, bnDigits },
    "ca": LocFmt{ ',', '.', false, normalDigits },
    "cs": LocFmt{ ',', ' ', false, normalDigits },
    "da": LocFmt{ ',', '.', false, normalDigits },
    "de": LocFmt{ ',', '.', false, normalDigits },
    "el": LocFmt{ ',', '.', false, normalDigits },
    "en": LocFmt{ '.', ',', false, normalDigits },
    "es": LocFmt{ ',', '.', false, normalDigits },
    "et": LocFmt{ ',', ' ', false, normalDigits },
    "fa": LocFmt{ '٫', '٬', false, faDigits },
    "fi": LocFmt{ ',', ' ', false, normalDigits },
    "fil": LocFmt{ '.', ',', false, normalDigits },
    "fr": LocFmt{ ',', ' ', false, normalDigits },
    "gu": LocFmt{ '.', ',', true, normalDigits },
    "he": LocFmt{ '.', ',', false, normalDigits },
    "hi": LocFmt{ '.', ',', true, normalDigits },
    "hr": LocFmt{ ',', '.', false, normalDigits },
    "hu": LocFmt{ ',', ' ', false, normalDigits },
    "hy": LocFmt{ ',', ' ', false, normalDigits },
    "id": LocFmt{ ',', '.', false, normalDigits },
    "is": LocFmt{ ',', '.', false, normalDigits },
    "it": LocFmt{ ',', '.', false, normalDigits },
    "ja": LocFmt{ '.', ',', false, normalDigits },
    "ka": LocFmt{ ',', ' ', false, normalDigits },
    "kk": LocFmt{ ',', ' ', false, normalDigits },
    "km": LocFmt{ ',', '.', false, normalDigits },
    "kn": LocFmt{ '.', ',', false, normalDigits },
    "ko": LocFmt{ '.', ',', false, normalDigits },
    "ky": LocFmt{ ',', ' ', false, normalDigits },
    "lo": LocFmt{ ',', '.', false, normalDigits },
    "lt": LocFmt{ ',', ' ', false, normalDigits },
    "lv": LocFmt{ ',', ' ', false, normalDigits },
    "mk": LocFmt{ ',', '.', false, normalDigits },
    "ml": LocFmt{ '.', ',', true, normalDigits },
    "mn": LocFmt{ '.', ',', false, normalDigits },
    "mo": LocFmt{ ',', '.', false, normalDigits },
    "mr": LocFmt{ '.', ',', true, mrDigits },
    "ms": LocFmt{ '.', ',', false, normalDigits },
    "mul": LocFmt{ '.', ',', false, normalDigits },
    "my": LocFmt{ '.', ',', false, myDigits },
    "nb": LocFmt{ ',', ' ', false, normalDigits },
    "ne": LocFmt{ '.', ',', false, mrDigits },
    "nl": LocFmt{ ',', '.', false, normalDigits },
    "no": LocFmt{ '.', ',', false, normalDigits },
    "pa": LocFmt{ '.', ',', true, normalDigits },
    "pl": LocFmt{ ',', ' ', false, normalDigits },
    "pt": LocFmt{ ',', '.', false, normalDigits },
    "ro": LocFmt{ ',', '.', false, normalDigits },
    "ru": LocFmt{ ',', ' ', false, normalDigits },
    "sh": LocFmt{ ',', '.', false, normalDigits },
    "si": LocFmt{ '.', ',', false, normalDigits },
    "sk": LocFmt{ ',', ' ', false, normalDigits },
    "sl": LocFmt{ ',', '.', false, normalDigits },
    "sq": LocFmt{ ',', ' ', false, normalDigits },
    "sr": LocFmt{ ',', '.', false, normalDigits },
    "sv": LocFmt{ ',', ' ', false, normalDigits },
    "sw": LocFmt{ '.', ',', false, normalDigits },
    "ta": LocFmt{ '.', ',', true, normalDigits },
    "te": LocFmt{ '.', ',', false, normalDigits },
    "th": LocFmt{ '.', ',', false, normalDigits },
    "tl": LocFmt{ '.', ',', false, normalDigits },
    "tn": LocFmt{ '.', ',', false, normalDigits },
    "tr": LocFmt{ ',', '.', false, normalDigits },
    "uk": LocFmt{ ',', ' ', false, normalDigits },
    "ur": LocFmt{ '.', ',', false, normalDigits },
    "uz": LocFmt{ ',', ' ', false, normalDigits },
    "vi": LocFmt{ ',', '.', false, normalDigits },
    "zh": LocFmt{ '.', ',', false, normalDigits },
    "zu": LocFmt{ '.', ',', false, normalDigits },
}

// get locale formating info
func GetLocFmt(lang string) *LocFmt {
    outLang := lang
    langSlen := len(lang)
    if langSlen>=3 && (lang[2]=='_' || lang[2]=='-') {
        outLang = lang[0:2]
    } else if langSlen>=4 && (lang[3]=='_' || lang[3]=='-') {
        outLang = lang[0:3]
    }
    l, ok := localeFormats[outLang]
    if !ok { l = defaultLocaleFormat }
    return &l
}

// format 128-bit unsigned integer including locale
func (a UInt128) LocaleFormat(lang string) string {
    l := GetLocFmt(lang)
    s := a.Format()
    os := make([]rune, 0, len(s))
    slen := len(s)
    ti := slen
    i := slen
    if !l.Sep100and1000 {
        ti = (slen)%3
        if ti==0 { ti=3 }
    }
    for _, r := range s {
        if r>='0' && r<='9' {
            os = append(os, l.Digits[r-'0'])
        }
        if i!=1 {
            if !l.Sep100and1000 || ti<=3 {
                ti--
                if ti==0 {
                    os = append(os, l.Sep1000)
                    ti = 3
                }
            } else {
                ti--
                if (ti-3)&1==0 {
                    os = append(os, l.Sep1000)
                }
            }
        }
        i--
    }
    return string(os)
}

// parse unsigned integer from string and return value and error (nil if no error)
func LocaleParseUInt128(lang, str string) (UInt128, error) {
    l := GetLocFmt(lang)
    // check whether localized number
    if len(str)==0 { return UInt128{}, strconv.ErrSyntax }
    
    os := make([]rune, 0, len(str))
    for _, r := range str {
        if r>='0' && r<='9' {
            os = append(os, r)
        } else if r!=l.Sep1000 {
            // if non-standard digit
            dig:=0
            found := false
            for ; dig<=9; dig++ {
                if l.Digits[dig]==r {
                    found = true
                    break
                }
            }
            if !found { return UInt128{}, strconv.ErrSyntax }
            os = append(os, '0'+rune(dig))
        }
        // otherwise skip sep1000
    }
    return ParseUInt128(string(os))
}
