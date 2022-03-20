package strings

import (
	"fmt"
	"strconv"
	"time"
	"unicode"

	"github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/option"
	"golang.org/x/exp/constraints"
)

// Common date formats.
const (
	YYYYMMDD         = "2006-01-02"
	YYYYMMDDNoDashes = "20060102"
	YYYYMD           = "2006-1-2"
	MMDDYYYY         = "01-02-2006"
	MMDDYYYYNoDashes = "01022006"
	MDYYYY           = "1-2-2006"
	MDYYYYName       = "Jan 2, 2006"
	MDYYYYFullName   = "January 2, 2006"

	HHmmss = ""
)

var dateFormats = []string{
	YYYYMMDD,
	YYYYMMDDNoDashes,
	YYYYMD,
	MMDDYYYY,
	MMDDYYYYNoDashes,
	MDYYYY,
	MDYYYYName,
	MDYYYYFullName,
}

// FromBool converts a bool into either "True" or "False".
func FromBool[TBool ~bool](b TBool) string {
	return strconv.FormatBool(bool(b))
}

// ToBool converts a string or rune to bool:
// "true", "t", or "1" case insensitive converts to true
// "false", "f", or "0" case insensitive converts to false.
func ToBool[TString StringOrRune](s TString) (bool, error) {
	return strconv.ParseBool(string(s))
}

// ToBoolOpt converts a string or rune to an optional bool
// following the rules of ToBool. If the string will not convert to
// a bool, then it returns None.
func ToBoolOpt[TString StringOrRune](s TString) option.Option[bool] {
	b, err := ToBool(s)
	if err != nil {
		return option.None[bool]()
	}
	return option.Some(b)
}

// FromInt converts an integer value to a string.
func FromInt[TInt constraints.Integer](i TInt) string {
	if i > 0 {
		return strconv.FormatUint(uint64(i), 10)
	}
	return strconv.FormatInt(int64(i), 10)
}

// ToInt converts from a string to an integer.
func ToInt[TString StringOrRune](s TString) (int, error) {
	return strconv.Atoi(string(s))
}

// ToInt8 converts from a string to an integer.
func ToInt8[TString StringOrRune](s TString) (int8, error) {
	i, err := strconv.ParseInt(string(s), 10, 8)
	return int8(i), err
}

// ToInt16 converts from a string to an integer.
func ToInt16[TString StringOrRune](s TString) (int16, error) {
	i, err := strconv.ParseInt(string(s), 10, 16)
	return int16(i), err
}

// ToInt32 converts from a string to an integer.
func ToInt32[TString StringOrRune](s TString) (int32, error) {
	i, err := strconv.ParseInt(string(s), 10, 32)
	return int32(i), err
}

// ToInt64 converts from a string to an integer.
func ToInt64[TString StringOrRune](s TString) (int64, error) {
	i, err := strconv.ParseInt(string(s), 10, 64)
	return int64(i), err
}

// ToIntOpt converts from a string to to an optional integer but
// on failure, returns None.
func ToIntOpt[TString StringOrRune](s TString) option.Option[int] {
	i, err := ToInt(s)
	if err != nil {
		return option.None[int]()
	}
	return option.Some(i)
}

// FromFloat converts a float to a string.
func FromFloat[TFloat constraints.Float](f TFloat) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

// ToFloat converts a string to a float64.
// If it fails, it returns a *NumError error from the
// strconv package.
func ToFloat[TString StringOrRune](s TString) (float64, error) {
	return strconv.ParseFloat(string(s), 64)
}

// ToFloatOpt converts from a string to to an optional float64 but
// on failure, returns None.
func ToFloatOpt[TString StringOrRune](s TString) option.Option[float64] {
	f, err := ToFloat(s)
	if err != nil {
		return option.None[float64]()
	}
	return option.Some(f)
}

// ToDate accepts a date as a string, with an optional format to use in parsing it.
// If no format is supplied, it uses a predefined list and tries them until it finds one
// that succeeds. The predefined formats are only for dates. They do not parse times.
func ToDate[TString ~string](s TString, format ...string) (time.Time, error) {
	f := findDateFormat(s, format...)
	if f.IsNone() {
		return time.Now(), fmt.Errorf("%s is not a valid date according to the given format", s)
	}
	return time.Parse(f.Value(), string(s))
}

func findDateFormat[TString ~string](s TString, formats ...string) option.Option[string] {
	if len(formats) == 0 {
		formats = dateFormats
	}
	return list.TryFind(func(f string) bool {
		_, err := time.Parse(f, string(s))
		return err == nil
	}, formats...)
}

// CamelCaseToUnderscore converts camel case strings to its equivalent as underscore separated.
func CamelCaseToUnderscore[TString ~string](camel TString) TString {
	return Collecti(func(i int, r rune) TString {
		if unicode.IsLower(r) {
			return TString(r)
		}
		r = unicode.ToLower(r)
		if i == 0 {
			return TString(r)
		}
		return "_" + TString(r)
	}, camel)
}

// NormalizeNewLine replaces any "non-normalized" newlines ("\r\n", '\r') with '\n'.
func NormalizeNewLine(s string) string {
	return ReplaceAll(ReplaceAll(s, "\r\n", '\n'), '\r', '\n')
}
