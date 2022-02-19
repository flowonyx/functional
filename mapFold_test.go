package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleMapFold() {
	type Charge struct {
		Type  ChargeType
		Value int
	}
	input := []Charge{{Type: In, Value: 1}, {Type: Out, Value: 2}, {Type: In, Value: 3}}
	r, s := functional.MapFold(func(acc int, charge Charge) (Charge, int) {
		val := charge.Value
		charge.Value *= 2
		if charge.Type == In {
			return charge, acc + val
		} else {
			return charge, acc - val
		}
	}, 0, input)
	fmt.Println(r, s)
	// Output: [{In 2} {Out 4} {In 6}] 2
}

func ExampleMapFoldBack() {
	type Charge struct {
		Type  ChargeType
		Value int
	}
	input := []Charge{{Type: In, Value: 1}, {Type: Out, Value: 2}, {Type: In, Value: 3}}
	r, s := functional.MapFoldBack(func(acc int, charge Charge) (Charge, int) {
		val := charge.Value
		charge.Value *= 2
		if charge.Type == In {
			return charge, acc + val
		} else {
			return charge, acc - val
		}
	}, 0, input)
	fmt.Println(r, s)
	// Output: [{In 6} {Out 4} {In 2}] 2
}
