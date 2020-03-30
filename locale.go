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

//import "strconv"

type locFmt struct {
    comma, sep1000 rune
    sep100and1000 bool
    digits []rune
}

var normalDigits []rune = []rune("0123456789")
var arDigits []rune = []rune("٠١٢٣٤٥٦٧٨٩")
var faDigits []rune = []rune("۰۱۲۳۴۵۶۷۸۹")
var bnDigits []rune = []rune("০১২৩৪৫৬৭৮৯")
var mrDigits []rune = []rune("०१२३४५६७८९")
var myDigits []rune = []rune("၀၁၂၃၄၅၆၇၈၉")

var defaultLocaleFormat locFmt = locFmt{ '.', ',', false, normalDigits }

var localeFormats map[string]locFmt = map[string]locFmt {
    "af": locFmt{ ',', ' ', false, normalDigits },
    "am": locFmt{ '.', ',', false, normalDigits },
    "ar": locFmt{ '٫', '٬', false, arDigits },
    "az": locFmt{ ',', '.', false, normalDigits },
    "bg": locFmt{ ',', ' ', false, normalDigits },
    "bn": locFmt{ '.', ',', true, bnDigits },
    "ca": locFmt{ ',', '.', false, normalDigits },
    "cs": locFmt{ ',', ' ', false, normalDigits },
    "da": locFmt{ ',', '.', false, normalDigits },
    "de": locFmt{ ',', '.', false, normalDigits },
    "el": locFmt{ ',', '.', false, normalDigits },
    "en": locFmt{ '.', ',', false, normalDigits },
    "es": locFmt{ ',', '.', false, normalDigits },
    "et": locFmt{ ',', ' ', false, normalDigits },
    "fa": locFmt{ '٫', '٬', false, faDigits },
    "fi": locFmt{ ',', ' ', false, normalDigits },
    "fil": locFmt{ '.', ',', false, normalDigits },
    "fr": locFmt{ ',', ' ', false, normalDigits },
    "gu": locFmt{ '.', ',', true, normalDigits },
    "he": locFmt{ '.', ',', false, normalDigits },
    "hi": locFmt{ '.', ',', true, normalDigits },
    "hr": locFmt{ ',', '.', false, normalDigits },
    "hu": locFmt{ ',', ' ', false, normalDigits },
    "hy": locFmt{ ',', ' ', false, normalDigits },
    "id": locFmt{ ',', '.', false, normalDigits },
    "is": locFmt{ ',', '.', false, normalDigits },
    "it": locFmt{ ',', '.', false, normalDigits },
    "ja": locFmt{ '.', ',', false, normalDigits },
    "ka": locFmt{ ',', ' ', false, normalDigits },
    "kk": locFmt{ ',', ' ', false, normalDigits },
    "km": locFmt{ ',', '.', false, normalDigits },
    "kn": locFmt{ '.', ',', false, normalDigits },
    "ko": locFmt{ '.', ',', false, normalDigits },
    "ky": locFmt{ ',', ' ', false, normalDigits },
    "lo": locFmt{ ',', '.', false, normalDigits },
    "lt": locFmt{ ',', ' ', false, normalDigits },
    "lv": locFmt{ ',', ' ', false, normalDigits },
    "mk": locFmt{ ',', '.', false, normalDigits },
    "ml": locFmt{ '.', ',', true, normalDigits },
    "mn": locFmt{ '.', ',', false, normalDigits },
    "mo": locFmt{ ',', '.', false, normalDigits },
    "mr": locFmt{ '.', ',', true, mrDigits },
    "ms": locFmt{ '.', ',', false, normalDigits },
    "mul": locFmt{ '.', ',', false, normalDigits },
    "my": locFmt{ '.', ',', false, myDigits },
    "nb": locFmt{ ',', ' ', false, normalDigits },
    "ne": locFmt{ '.', ',', false, mrDigits },
    "nl": locFmt{ ',', '.', false, normalDigits },
    "no": locFmt{ '.', ',', false, normalDigits },
    "pa": locFmt{ '.', ',', true, normalDigits },
    "pl": locFmt{ ',', ' ', false, normalDigits },
    "pt": locFmt{ ',', '.', false, normalDigits },
    "ro": locFmt{ ',', '.', false, normalDigits },
    "ru": locFmt{ ',', ' ', false, normalDigits },
    "sh": locFmt{ ',', '.', false, normalDigits },
    "si": locFmt{ '.', ',', false, normalDigits },
    "sk": locFmt{ ',', ' ', false, normalDigits },
    "sl": locFmt{ ',', '.', false, normalDigits },
    "sq": locFmt{ ',', ' ', false, normalDigits },
    "sr": locFmt{ ',', '.', false, normalDigits },
    "sv": locFmt{ ',', ' ', false, normalDigits },
    "sw": locFmt{ '.', ',', false, normalDigits },
    "ta": locFmt{ '.', ',', true, normalDigits },
    "te": locFmt{ '.', ',', false, normalDigits },
    "th": locFmt{ '.', ',', false, normalDigits },
    "tl": locFmt{ '.', ',', false, normalDigits },
    "tn": locFmt{ '.', ',', false, normalDigits },
    "tr": locFmt{ ',', '.', false, normalDigits },
    "uk": locFmt{ ',', ' ', false, normalDigits },
    "ur": locFmt{ '.', ',', false, normalDigits },
    "uz": locFmt{ ',', ' ', false, normalDigits },
    "vi": locFmt{ ',', '.', false, normalDigits },
    "zh": locFmt{ '.', ',', false, normalDigits },
    "zu": locFmt{ '.', ',', false, normalDigits },
}

func getLocFmt(lang string) *locFmt {
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
    l := getLocFmt(lang)
    s := a.Format()
    os := make([]rune, 0, len(s))
    slen := len(s)
    ti := slen
    i := slen
    if !l.sep100and1000 {
        ti = (slen)%3
        if ti==0 { ti=3 }
    }
    for _, r := range s {
        if r>='0' && r<='9' {
            os = append(os, l.digits[r-'0'])
        }
        if i!=1 {
            if !l.sep100and1000 || ti<=3 {
                ti--
                if ti==0 {
                    os = append(os, l.sep1000)
                    ti = 3
                }
            } else {
                ti--
                if (ti-3)&1==0 {
                    os = append(os, l.sep1000)
                }
            }
        }
        i--
    }
    return string(os)
}

// parse unsigned integer from string and return value and error (nil if no error)
/*func LocaleParseUInt128(lang, str string) (UInt128, error) {
    l := getLocFmt(lang)
    // check whether localized number
    if len(str)==0 { return UInt128{}, strconv.ErrSyntax }
    
    //if str[0]>='0' || str
    
    return UInt128{}, strconv.ErrSyntax
}*/
