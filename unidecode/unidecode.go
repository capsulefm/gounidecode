// Unidecode implements a unicode transliterator, which
// replaces non-ASCII characters with their ASCII
// counterparts

package unidecode

import (
	"unicode"
)

// Given an unicode encoded string, returns
// another string with non-ASCII characters replaced
// with their closest ASCII counterparts.
// e.g. Unicode("áéíóú") => "aeiou"

var converters = map[string]map[rune]rune{
	"fr": {
		'À': 'À',
		'à': 'à',
		'Â': 'Â',
		'â': 'â',
		'Æ': 'Æ',
		'æ': 'æ',
		'Ç': 'Ç',
		'ç': 'ç',
		'É': 'É',
		'é': 'é',
		'È': 'È',
		'è': 'è',
		'Ê': 'Ê',
		'ê': 'ê',
		'Ë': 'Ë',
		'ë': 'ë',
		'Î': 'Î',
		'î': 'î',
		'Ï': 'Ï',
		'ï': 'ï',
		'Ô': 'Ô',
		'ô': 'ô',
		'Œ': 'Œ',
		'œ': 'œ',
		'Ù': 'Ù',
		'ù': 'ù',
		'Û': 'Û',
		'û': 'û',
		'Ü': 'Ü',
		'ü': 'ü',
	},
	"no": {
		'æ': 'æ',
		'ø': 'ø',
		'å': 'å',
		'Æ': 'Æ',
		'Ø': 'Ø',
		'Å': 'Å',
		'ä': 'æ',
		'ö': 'ø',
		'Ä': 'Æ',
		'Ö': 'Ø',
	},
	"da": {
		'æ': 'æ',
		'ø': 'ø',
		'å': 'å',
		'Æ': 'Æ',
		'Ø': 'Ø',
		'Å': 'Å',
		'ä': 'æ',
		'ö': 'ø',
		'Ä': 'Æ',
		'Ö': 'Ø',
	},
	"sv": {
		'æ': 'ä',
		'ø': 'ö',
		'å': 'å',
		'Æ': 'Ä',
		'Ø': 'Ö',
		'Å': 'Å',
		'ä': 'ä',
		'ö': 'ö',
		'Ä': 'Ä',
		'Ö': 'Ö',
	},
	"de": {
		'Ä': 'Ä',
		'ä': 'ä',
		'Ö': 'Ö',
		'ö': 'ö',
		'Ü': 'Ü',
		'ü': 'ü',
		'ß': 'ß',
	},
	"en": {},
}

func Unidecode(s, toLangCode string) string {
	convert := converters[toLangCode]

	if convert == nil {
		return s
	}
	str := ""
	for _, c := range s {
		if c <= unicode.MaxASCII {
			str += string(c)
			continue
		}
		if convert != nil {
			n := convert[c]
			//			fmt.Println("Unidecode:", toLangCode, string(c), string(n))
			if n != 0 {
				str += string(n)
				continue
			}
		}
		if c > unicode.MaxRune {
			/* Ignore reserved chars */
			continue
		}
		if d, ok := transliterations[c]; ok {
			str += d
		}
	}
	return str
}
