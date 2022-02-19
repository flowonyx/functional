package string

import (
	"strings"

	"github.com/flowonyx/functional"
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
	functional.DoRangeTo(func(i int) { output.WriteString(initializer(i)) }, count)
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

func Mapi(mapping func(int, rune) rune, str string) string {
	output := &strings.Builder{}
	Iteri(func(i int, r rune) { output.WriteRune(mapping(i, r)) }, str)
	return output.String()
}
