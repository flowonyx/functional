package list_test

import (
	"fmt"
	"strconv"

	"github.com/flowonyx/functional/list"
	"github.com/flowonyx/functional/option"
)

func ExamplePick() {
	i, err := list.Pick(func(s string) option.Option[int] {
		if i, err := strconv.Atoi(s); err == nil {
			return option.Some(i)
		}
		return option.None[int]()
	}, []string{"zero", "1", "2", "three", "4"})
	if err != nil {
		panic(err)
	}
	fmt.Println(i)
	// Output: 1
}

func ExampleTryPick() {
	i := list.TryPick(func(s string) option.Option[int] {
		if i, err := strconv.Atoi(s); err == nil {
			return option.Some(i)
		}
		return option.None[int]()
	}, []string{"zero", "1", "2", "three", "4"})

	fmt.Println(i.Value())
	// Output: 1
}

func ExampleChoose() {
	i := list.Choose(func(s string) option.Option[int] {
		if i, err := strconv.Atoi(s); err == nil {
			return option.Some(i)
		}
		return option.None[int]()
	}, []string{"zero", "1", "2", "three", "4"})

	fmt.Println(i)
	// Output: [1 2 4]
}
