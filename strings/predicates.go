package strings

import (
	"net/mail"

	"github.com/flowonyx/functional/list"
)

// IsDate checks if the given string is a date according to either the given format(s)
// or the default formats if none are provided.
func IsDate[TString ~string](s TString, formats ...string) bool {
	return findDateFormat(s, formats...).IsSome()
}

// IsEmail checks if the given string can be parsed as an email address.
func IsEmail[TString ~string](s TString) bool {
	if _, err := mail.ParseAddress(string(s)); err != nil {
		return false
	}
	return true
}

// RuneIsAnyOf checks if the the given rune r is in the given list of runes.
func RuneIsAnyOf[TRune, TRuneList ~rune](runes []TRuneList, r TRune) bool {
	return list.Contains(TRuneList(r), runes)
}

// RuneIsASCIILower checks if the given rune is an ASCII lower case character.
func RuneIsASCIILower[TRune ~rune](r TRune) bool {
	return r >= 'a' && r <= 'z'
}

// RuneIsASCIILower checks if the given rune is an ASCII upper case character.
func RuneIsASCIIUpper[TRune ~rune](r TRune) bool {
	return r >= 'A' && r <= 'Z'
}

// RuneIsASCIILower checks if the given rune is an ASCII character (either upper or lower case).
func RuneIsASCIILetter[TRune ~rune](r TRune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}

// RuneIsHex checks if the given rune is a valid hexadecimal numeral (0-9, a-f, A-F).
func RuneIsHex[TRune ~rune](r TRune) bool {
	return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')
}

// RuneIsOctal checks if the given rune is a valid octal numeral (0-7).
func RuneIsOctal[TRune ~rune](r TRune) bool {
	return r >= '0' && r <= '7'
}

// RuneIsNewLine checks if the given rune is a valid newline character ('\n' or '\r').
func RuneIsNewLine[TRune ~rune](r TRune) bool {
	return RuneIsAnyOf([]rune{'\n', '\r'}, r)
}

// RuneIsNoneOf checks whether the given rune r is in the given list of runes and returns false if it is present.
func RuneIsNoneOf[TRune, TRuneList ~rune](runes []TRuneList, r TRune) bool {
	return !list.Contains(TRuneList(r), runes)
}
