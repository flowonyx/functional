package strings

import (
	"strings"

	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/list"
)

func Collect(mapping functional.Projection[rune, string], str string) string {
	output := &strings.Builder{}
	for _, r := range str {
		output.WriteString(mapping(r))
	}
	return output.String()
}

func Exists(predicate functional.Predicate[rune], str string) bool {
	if i := strings.IndexFunc(str, predicate); i >= 0 {
		return true
	}
	return false
}

func Filter(predicate functional.Predicate[rune], str string) string {
	output := &strings.Builder{}
	for _, r := range str {
		if predicate(r) {
			output.WriteRune(r)
		}
	}
	return output.String()
}

func ForAll(predicate functional.Predicate[rune], str string) bool {
	for _, r := range str {
		if !predicate(r) {
			return false
		}
	}
	return true
}

func InitString(count int, initializer func(int) string) string {
	output := &strings.Builder{}
	list.DoRangeTo(func(i int) { output.WriteString(initializer(i)) }, count-1)
	return output.String()
}

func Iter(action func(rune), str string) {
	for _, r := range str {
		action(r)
	}
}

func Iteri(action func(int, rune), str string) {
	for i, r := range str {
		action(i, r)
	}
}

func Map(mapping functional.Projection[rune, rune], str string) string {
	return strings.Map(mapping, str)
}

func MapTo[T any](mapping functional.Projection[rune, T], str string) []T {
	output := make([]T, len(str))
	for i, r := range str {
		output[i] = mapping(r)
	}
	return output
}

func Mapi(mapping func(int, rune) rune, str string) string {
	output := &strings.Builder{}
	Iteri(func(i int, r rune) { output.WriteRune(mapping(i, r)) }, str)
	return output.String()
}

func HasPrefix[T ~string | ~rune](s string, prefix T) bool {
	return strings.HasPrefix(s, string(prefix))
}

func HasSuffix[T ~string | ~rune](s string, suffix T) bool {
	return strings.HasSuffix(s, string(suffix))
}

func Range(start, end rune) []rune {
	r := []rune{}
	for i := int(start); i <= int(end); i++ {
		r = append(r, rune(i))
	}
	return r
}

func FromRunes(input []rune) string {
	return InitString(len(input), func(i int) string { return string(input[i]) })
}
