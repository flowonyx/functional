// Package strings provides methods that work on strings or runes--but also types descended from them.
// Where possible, the functions take either a string or a rune.
// Most, if not all, functions from the standard "strings" package are replicated here, but with generic inputs.
package strings

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/option"

	"golang.org/x/exp/constraints"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// aliases so we do not need to separately import the standard "strings" package
// if we are already using this package.

type Builder = strings.Builder
type Reader = strings.Reader
type Replacer = strings.Replacer

// StringOrRune is a type constraint that includes
// any type descended from string or from rune.
type StringOrRune interface {
	~string | ~rune
}

// Clone returns a fresh copy of s. It guarantees to make a copy of s into a new allocation,
// which can be important when retaining only a small substring of a much larger string.
// Using Clone can help such programs use less memory. Of course, since using Clone makes a copy,
// overuse of Clone can make programs use more memory. Clone should typically be used only rarely,
// and only when profiling indicates that it is needed. For strings of length zero the
// string "" will be returned and no allocation is made.
func Clone[TString ~string](s TString) TString {
	return TString(strings.Clone(string(s)))
}

// Collect accepts a function which maps each rune in the given string to a string,
// then concatenates them together into one string. If str is a descendant of string
// the actual type of str will be returned.
func Collect[TString1, TString2 ~string, TRune ~rune](mapping func(TRune) TString2, str TString1) TString1 {
	output := &strings.Builder{}
	for _, r := range str {
		output.WriteString(string(mapping(TRune(r))))
	}
	return TString1(output.String())
}

// Collecti accepts a function which maps each rune in the given string to a string,
// then concatenates them together into one string. If str is a descendant of string
// the actual type of str will be returned.
// The difference between Collecti and Collect is that the mapping function receives the index of the rune.
func Collecti[TString1, TString2 ~string, TRune ~rune](mapping func(int, TRune) TString2, str TString1) TString1 {
	output := &strings.Builder{}
	for i, r := range str {
		output.WriteString(string(mapping(i, TRune(r))))
	}
	return TString1(output.String())
}

// Compare returns an integer comparing two strings lexicographically. The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
// Compare is included only for symmetry with package bytes. It is usually clearer and always faster to use the built-in string comparison operators ==, <, >, and so on.
func Compare[TString1, TString2 StringOrRune](a TString1, b TString2) int {
	return strings.Compare(string(a), string(b))
}

// Concat concatenates a list of strings or runes into one string.
func Concat[T StringOrRune](s []T) string {
	v := &strings.Builder{}
	for _, i := range s {
		v.WriteString(string(i))
	}
	return v.String()
}

// Contains reports whether substr is within s.
func Contains[TString1, TString2 StringOrRune](s TString1, substr TString2) bool {
	return strings.Contains(string(s), string(substr))
}

// ContainsAny reports whether any Unicode code points in chars are within s.
func ContainsAny[TString1, TString2 StringOrRune](s TString1, chars TString2) bool {
	return strings.ContainsAny(string(s), string(chars))
}

// ContainsRune reports whether the Unicode code point r is within s.
func ContainsRune[TString ~string, TRune ~rune](s TString, r TRune) bool {
	return strings.ContainsRune(string(s), rune(r))
}

// Count counts the number of non-overlapping instances of substr in s.
// If substr is an empty string, Count returns 1 + the number of Unicode code points in s.
func Count[TString1, TString2 StringOrRune](s TString1, substr TString2) int {
	return strings.Count(string(s), string(substr))
}

// Cut slices s around the first instance of sep, returning the text before and after sep.
// The found result reports whether sep appears in s. If sep does not appear in s, cut returns s, "", false.
func Cut[TString ~string, TSep StringOrRune](s TString, sep TSep) (before, after TString, found bool) {
	b, a, f := strings.Cut(string(s), string(sep))
	return TString(b), TString(a), f
}

// EqualFold reports whether s and t, interpreted as UTF-8 strings,
// are equal under Unicode case-folding, which is a more general form of case-insensitivity.
func EqualFold[TString1, TString2 StringOrRune](s TString1, t TString2) bool {
	return strings.EqualFold(string(s), string(t))
}

// Exists checks for the existence of a rune within str that matches the predicate.
func Exists[TString ~string, TRune ~rune](predicate func(TRune) bool, str TString) bool {
	for _, r := range str {
		if predicate(TRune(r)) {
			return true
		}
	}
	return false
}

// Fields splits the string s around each instance of one or more consecutive white space characters,
// as defined by unicode.IsSpace, returning a slice of substrings of s or an empty slice if s contains only white space.
func Fields[TString ~string](s TString) []TString {
	return list.Map(
		func(s string) TString {
			return TString(s)
		}, strings.Fields(string(s)))
}

// FieldsFunc splits the string s at each run of Unicode code points c satisfying f(c) and returns an array of slices of s.
// If all code points in s satisfy f(c) or the string is empty, an empty slice is returned.
// FieldsFunc makes no guarantees about the order in which it calls f(c) and assumes that f always returns the same value for a given c.
func FieldsFunc[TString ~string, TRune ~rune](dividerPredicate func(TRune) bool, s TString) []TString {
	f := func(r rune) bool { return dividerPredicate(TRune(r)) }
	return list.Map(
		func(s string) TString {
			return TString(s)
		}, strings.FieldsFunc(string(s), f))
}

// Filter returns a string built from str that only contains runes matching the predicate.
func Filter[TString ~string, TRune ~rune](predicate func(TRune) bool, str TString) TString {
	output := &strings.Builder{}
	for _, r := range str {
		if predicate(TRune(r)) {
			output.WriteRune(r)
		}
	}
	return TString(output.String())
}

// Fold implements Unicode case folding. Case folding does not normalize the input and may not preserve a normal form.
// Use the collate or search package for more convenient and linguistically sound comparisons.
// Use golang.org/x/text/secure/precis for string comparisons where security aspects are a concern.
func Fold[TString ~string](s TString) TString {
	return TString(cases.Fold().String(string(s)))
}

// ForAll checks whether every rune in str matches the predicate.
func ForAll[TString ~string, TRune ~rune](predicate func(TRune) bool, str TString) bool {
	for _, r := range str {
		if !predicate(TRune(r)) {
			return false
		}
	}
	return true
}

// FromRunes simply creates a string from the given runes. It is faster to just call string(input) instead,
// but sometimes it is convenient to have a function that can be passed around.
func FromRunes[TRune ~rune](input []TRune) string {
	r := list.Map(func(r TRune) rune { return rune(r) }, input)
	return string(r)
}

// HasPrefix tests whether the string s begins with prefix.
func HasPrefix[TString, TPrefix StringOrRune](s TString, prefix TPrefix) bool {
	return strings.HasPrefix(string(s), string(prefix))
}

// HasSuffix tests whether the string s ends with suffix.
func HasSuffix[TString, TSuffix StringOrRune](s TString, prefix TSuffix) bool {
	return strings.HasSuffix(string(s), string(prefix))
}

// Index returns the index of the first instance of substr in s, or -1 if substr is not present in s.
func Index[TString1, TString2 StringOrRune](s TString1, substr TString2) int {
	return strings.Index(string(s), string(substr))
}

// IndexAny returns the index of the first instance of any Unicode code point from chars in s, or -1 if no Unicode code point from chars is present in s.
func IndexAny[TString1, TString2 StringOrRune](s TString1, chars TString2) int {
	return strings.IndexAny(string(s), string(chars))
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexByte[TString1 StringOrRune, TByte ~byte](s TString1, c TByte) int {
	return strings.IndexByte(string(s), byte(c))
}

// IndexByte returns the index of the first instance of c in s, or -1 if c is not present in s.
func IndexFunc[TString1 StringOrRune](s TString1, f func(rune) bool) int {
	return strings.IndexFunc(string(s), f)
}

// IndexRune returns the index of the first instance of the Unicode code point r, or -1 if rune is not present in s.
// If r is utf8.RuneError, it returns the first instance of any invalid UTF-8 byte sequence.
func IndexRune[TString1 StringOrRune, TRune ~rune](s TString1, r TRune) int {
	return strings.IndexRune(string(s), rune(r))
}

// InitString creates a string of length count where each rune is initialized
// by the initializer function.
func InitString[TString ~string, TInt1, TInt2 constraints.Integer](count TInt1, initializer func(TInt2) TString) TString {
	if count <= 0 {
		return TString("")
	}
	output := &strings.Builder{}
	list.DoRangeTo(func(i TInt2) { output.WriteString(string(initializer(i))) }, TInt2(count-1))
	return TString(output.String())
}

// Iter performs an action for each rune in str, passing the rune to the action function.
func Iter[TString ~string, TRune ~rune](action func(TRune), str TString) {
	for _, r := range str {
		action(TRune(r))
	}
}

// Iteri performs an action for each rune in str, passing the index and the rune to the action function.
func Iteri[TString ~string](action func(int, rune), str TString) {
	for i, r := range str {
		action(i, r)
	}
}

// Join joins elems of any type into a string of sep separated values.
// It uses fmt.Sprint to represent each elem as a string.
// If you are passing runes or strings, it is slightly faster to call
// JoinRunes or JoinStrings instead.
func Join[TSep StringOrRune, T any](elems []T, sep TSep) string {
	output := &strings.Builder{}
	for i, v := range elems {
		output.WriteString(printv(v))
		if i < list.LastIndexOf(elems) {
			output.WriteString(string(sep))
		}
	}
	return output.String()
}

// JoinRunes joins runes into a string of sep seprated values.
// If you are trying to join them without a separator, use Concat
// or simply string(elems) instead.
func JoinRunes[TSep StringOrRune, TRune ~rune](elems []TRune, sep TSep) string {
	output := &strings.Builder{}
	for i, v := range elems {
		output.WriteRune(rune(v))
		if i < list.LastIndexOf(elems) {
			output.WriteString(string(sep))
		}
	}
	return output.String()
}

// JoinStrings joins strings into a string of sep seprated values.
// If you are trying to join them without a separator, use Concat instead.
func JoinStrings[TSep StringOrRune, TString ~string](elems []TString, sep TSep) string {
	return strings.Join(ToStringSlice(elems), string(sep))
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndex[TString1, TString2 StringOrRune](s TString1, substr TString2) int {
	return strings.LastIndex(string(s), string(substr))
}

// LastIndex returns the index of the last instance of substr in s, or -1 if substr is not present in s.
func LastIndexAny[TString1, TString2 StringOrRune](s TString1, chars TString2) int {
	return strings.LastIndexAny(string(s), string(chars))
}

// LastIndexByte returns the index of the last instance of c in s, or -1 if c is not present in s.
func LastIndexByte[TString1 StringOrRune, TByte ~byte](s TString1, c TByte) int {
	return strings.LastIndexByte(string(s), byte(c))
}

// LastIndexFunc returns the index into s of the last Unicode code point satisfying f(c), or -1 if none do.
func LastIndexFunc[TString1 StringOrRune](s TString1, f func(rune) bool) int {
	return strings.LastIndexFunc(string(s), f)
}

// LastIndexRune returns the index of the last instance of r in s, or -1 if r is not present in s.
func LastIndexRune[TString1 ~string, TRune ~rune](s TString1, r TRune) int {
	index := IndexRune(s, r)
	i := index
	for i > 0 && i < len(s)-1 {
		i = IndexRune(s[i+1:], r)
		if i < 0 {
			return index
		}
		index += i + 1
		i = index
	}
	return index
}

// Lines splits s on newline boundaries into a slice of strings.
// The results do not include the newlines.
func Lines[TString ~string](s TString) []TString {
	return Split(s, []string{"\r\n", "\r", "\n"}...)
}

// GetLine gets the value of the line in s indicated by index.
// An error is returned if the index is out of range.
func GetLine[TString ~string, TInt constraints.Integer](s TString, index TInt) (TString, error) {
	lines := Lines(s)
	if list.LastIndexOf(lines) < int(index) {
		return "", fmt.Errorf("%w: requested line out of range: %d", errors.BadArgumentErr, index)
	}
	return lines[index], nil
}

// GetLineOpt gets the value of the line in s indicated by index.
// None is returned if the index is out of range.
func GetLineOpt[TString ~string, TInt constraints.Integer](s TString, index TInt) option.Option[TString] {
	l, err := GetLine(s, index)
	if err != nil {
		return option.None[TString]()
	}
	return option.Some(l)
}

// Lower returns s with all lowercase letters based on the English language.
// If you need to use another language, call LowerSpecialCase directly.
func Lower[TString ~string](s TString) TString {
	return LowerSpecialCase(s, language.English)
}

// LowerSpecialCase uses language specific rules for returning s in lower case.
func LowerSpecialCase[TString ~string](s TString, language language.Tag) TString {
	return TString(cases.Lower(language).String(string(s)))
}

// Map returns a copy of the string s with all its characters modified according to the mapping function.
// If mapping returns a negative value, the character is dropped from the string with no replacement.
func Map[TString ~string, TRune ~rune](mapping func(TRune) TRune, s TString) TString {
	f := func(r rune) rune { return rune(mapping(TRune(r))) }
	return TString(strings.Map(f, string(s)))
}

// Mapi does the same as Map except that the mapping function is supplied with the index of the
// rune within the string along with the rune.
// If mapping returns a negative value, the character is dropped from the string with no replacement.
func Mapi[TString ~string, TRune ~rune](mapping func(int, TRune) TRune, s TString) TString {
	output := &strings.Builder{}
	Iteri(
		func(i int, r rune) {
			res := mapping(i, TRune(r))
			if res < 0 {
				return
			}
			output.WriteRune(rune(res))
		}, s)
	return TString(output.String())
}

// Range creates a slice of runes that between the start and end runes (inclusive).
// Range('A', 'C') would create []rune{'A', 'B', 'C'}.
func Range[TRune ~rune](start, end TRune) []TRune {
	r := []TRune{}
	for i := int(start); i <= int(end); i++ {
		r = append(r, TRune(i))
	}
	return r
}

// Repeat returns a new string consisting of count copies of the string s.
// It panics if count is negative or if the result of (len(s) * count) overflows.
func Repeat[TString StringOrRune, TCount constraints.Integer](s TString, count TCount) string {
	return strings.Repeat(string(s), int(count))
}

// Replace returns a copy of the string s with the first n non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
// If n < 0, there is no limit on the number of replacements.
func Replace[TString1 ~string, TString2, TString3 StringOrRune, TInt constraints.Integer](s TString1, old TString2, new TString3, n TInt) TString1 {
	return TString1(strings.Replace(string(s), string(old), string(new), int(n)))
}

// ReplaceAll returns a copy of the string s with all non-overlapping instances of old replaced by new.
// If old is empty, it matches at the beginning of the string and after each UTF-8 sequence, yielding up to k+1 replacements for a k-rune string.
func ReplaceAll[TString1 ~string, TString2, TString3 StringOrRune](s TString1, old TString2, new TString3) TString1 {
	return TString1(strings.ReplaceAll(string(s), string(old), string(new)))
}

// NewReplacer returns a new Replacer from a list of old, new string pairs.
// Replacements are performed in the order they appear in the target string, without overlapping matches.
// The old string comparisons are done in argument order.
// NewReplacer panics if given an odd number of arguments.
func NewReplacer[TString StringOrRune](oldnew ...TString) *Replacer {
	return strings.NewReplacer(ToStringSlice(oldnew)...)
}

// Split slices s into all substrings separated by sep and returns a slice of the substrings between those separators.
// If s does not contain sep and sep is not empty, Split returns a slice of length 1 whose only element is s.
// If sep is empty, Split splits after each UTF-8 sequence. If both s and sep are empty, Split returns an empty slice.
// It is equivalent to SplitN with a count of -1.
// To split around the first instance of a separator, see Cut.
func Split[TString ~string, TSep StringOrRune](s TString, seps ...TSep) []TString {
	return split(strings.Split, s, seps...)
}

// SplitAfter slices s into all substrings after each instance of sep and returns a slice of those substrings.
// If s does not contain sep and sep is not empty, SplitAfter returns a slice of length 1 whose only element is s.
// If sep is empty, SplitAfter splits after each UTF-8 sequence. If both s and sep are empty, SplitAfter returns an empty slice.
// It is equivalent to SplitAfterN with a count of -1.
func SplitAfter[TString ~string, TSep StringOrRune](s TString, seps ...TSep) []TString {
	return split(strings.SplitAfter, s, seps...)
}

// SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.
// The count determines the number of substrings to return:
// n > 0: at most n substrings; the last substring will be the unsplit remainder.
// n == 0: the result is nil (zero substrings)
// n < 0: all substrings
// Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for SplitAfter.
func SplitAfterN[TString ~string, TSep StringOrRune, TInt constraints.Integer](n TInt, s TString, seps ...TSep) []TString {
	return splitN(strings.SplitAfterN, n, s, seps...)
}

// SplitAfterN slices s into substrings after each instance of sep and returns a slice of those substrings.
// The count determines the number of substrings to return:
// n > 0: at most n substrings; the last substring will be the unsplit remainder.
// n == 0: the result is nil (zero substrings)
// n < 0: all substrings
// Edge cases for s and sep (for example, empty strings) are handled as described in the documentation for SplitAfter.
func SplitN[TString ~string, TSep StringOrRune, TInt constraints.Integer](n TInt, s TString, seps ...TSep) []TString {
	return splitN(strings.SplitN, n, s, seps...)
}

// Title returns a string with English title casing. It uses an approximation of the default Unicode Word Break algorithm.
// If you want to have the title casing specific to another language, use TitleSpecial instead.
func Title[TString ~string](s TString) TString {
	return TitleSpecial(s, language.English)
}

// TitleSpecial returns a language-specific title casing of the given string. It uses an approximation of the default Unicode Word Break algorithm.
func TitleSpecial[TString ~string](s TString, language language.Tag) TString {
	return TString(cases.Title(language).String(string(s)))
}

// ToStringSlice converts a slice of ~string or ~rune to a slice of string.
func ToStringSlice[TString StringOrRune](input []TString) []string {
	return list.Map(func(s TString) string { return string(s) }, input)
}

// Upper returns a string with English upper casing. It uses an approximation of the default Unicode Word Break algorithm.
// If you want to have the upper casing specific to another language, use UpperSpecial instead.
func Upper[TString ~string](s TString) TString {
	return UpperSpecial(s, language.English)
}

// UpperSpecial returns a string with language-specific upper casing of the given string. It uses an approximation of the default Unicode Word Break algorithm.
func UpperSpecial[TString ~string](s TString, language language.Tag) TString {
	return TString(cases.Upper(language).String(string(s)))
}

// ToValidUTF8 returns a copy of the string s with each run of invalid UTF-8 byte sequences replaced by the replacement string, which may be empty.
func ToValueUTF8[TString1, TString2 ~string](s TString1, replacement TString2) TString1 {
	return TString1(strings.ToValidUTF8(string(s), string(replacement)))
}

// Trim returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
func Trim[TString1 ~string, TString2 StringOrRune](s TString1, cutset TString2) TString1 {
	return TString1(strings.Trim(string(s), string(cutset)))
}

// TrimFunc returns a slice of the string s with all leading and trailing Unicode code points c satisfying f(c) removed.
func TrimFunc[TString ~string, TRune ~rune](s TString, f func(TRune) bool) TString {
	tf := func(r rune) bool { return f(TRune(r)) }
	return TString(strings.TrimFunc(string(s), tf))
}

// TrimLeft returns a slice of the string s with all leading Unicode code points contained in cutset removed.
// To remove a prefix, use TrimPrefix instead.
func TrimLeft[TString1 ~string, TString2 StringOrRune](s TString1, cutset TString2) TString1 {
	return TString1(strings.TrimLeft(string(s), string(cutset)))
}

// TrimLeftFunc returns a slice of the string s with all leading Unicode code points c satisfying f(c) removed.
func TrimLeftFunc[TString ~string, TRune ~rune](s TString, f func(TRune) bool) TString {
	tf := func(r rune) bool { return f(TRune(r)) }
	return TString(strings.TrimLeftFunc(string(s), tf))
}

// TrimPrefix returns s without the provided leading prefix string. If s doesn't start with prefix, s is returned unchanged.
func TrimPrefix[TString1 ~string, TString2 StringOrRune](s TString1, prefix TString2) TString1 {
	return TString1(strings.TrimPrefix(string(s), string(prefix)))
}

// TrimRight returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.
// To remove a suffix, use TrimSuffix instead.
func TrimRight[TString1 ~string, TString2 StringOrRune](s TString1, cutset TString2) TString1 {
	return TString1(strings.TrimRight(string(s), string(cutset)))
}

// TrimRightFunc returns a slice of the string s with all trailing Unicode code points c satisfying f(c) removed.
func TrimRightFunc[TString ~string, TRune ~rune](s TString, f func(TRune) bool) TString {
	tf := func(r rune) bool { return f(TRune(r)) }
	return TString(strings.TrimRightFunc(string(s), tf))
}

// TrimSpace returns a slice of the string s, with all leading and trailing white space removed, as defined by Unicode.
func TrimSpace[TString1 ~string](s TString1) TString1 {
	return TString1(strings.TrimSpace(string(s)))
}

// TrimSuffix returns s without the provided trailing suffix string. If s doesn't end with suffix, s is returned unchanged.
func TrimSuffix[TString1 ~string, TString2 StringOrRune](s TString1, suffix TString2) TString1 {
	return TString1(strings.TrimSuffix(string(s), string(suffix)))
}

// Quote returns a double-quoted Go string literal representing s.
// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for control characters and non-printable characters as defined by IsPrint.
func Quote[TString ~string](s TString) TString {
	return TString(strconv.Quote(string(s)))
}

// QuoteRune returns a single-quoted Go character literal representing the rune.
// The returned string uses Go escape sequences (\t, \n, \xFF, \u0100) for control characters and non-printable characters as defined by IsPrint.
func QuoteRune[TRune ~rune](r TRune) string {
	return strconv.QuoteRune(rune(r))
}

// Unquote interprets s as a single-quoted, double-quoted, or backquoted Go string literal, returning the string value that s quotes.
// (If s is single-quoted, it would be a Go character literal; Unquote returns the corresponding one-character string.)
func Unquote[TString ~string](s TString) TString {
	r, err := strconv.Unquote(string(s))
	if err != nil {
		return s
	}
	return TString(r)
}

func split[TString1 ~string, TString2 StringOrRune](splitter func(string, string) []string, s TString1, seps ...TString2) []TString1 {
	seps = list.SortByDescending(func(s TString2) int { return len(string(s)) }, seps)
	lastSep, err := list.Last(seps)
	if err != nil {
		return []TString1{s}
	}
	for _, sep := range seps[:len(seps)-1] {
		s = ReplaceAll(s, sep, lastSep)
	}
	return list.Map(func(s string) TString1 { return TString1(s) }, splitter(string(s), string(lastSep)))
}

func splitN[TString1 ~string, TString2 StringOrRune, TInt constraints.Integer](splitter func(string, string, int) []string, n TInt, s TString1, seps ...TString2) []TString1 {
	seps = list.SortByDescending(func(s TString2) int { return len(string(s)) }, seps)
	lastSep, err := list.Last(seps)
	if err != nil {
		return []TString1{s}
	}
	for _, sep := range seps[:len(seps)-1] {
		s = ReplaceAll(s, sep, lastSep)
	}
	return list.Map(func(s string) TString1 { return TString1(s) }, splitter(string(s), string(lastSep), int(n)))
}

func printv(v any) string {
	switch o := v.(type) {
	case string:
		return o
	case rune:
		return string(o)
	default:
		return fmt.Sprint(o)
	}
}
