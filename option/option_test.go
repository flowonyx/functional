package option_test

import (
	"fmt"
	"strconv"

	"github.com/flowonyx/functional"
	"github.com/flowonyx/functional/option"
)

func ExampleHandleOption() {
	input := option.Some(1)
	err := option.HandleOption(input, func(i int) error {
		fmt.Printf("%d", i)
		return nil
	}, func() error {
		fmt.Println("None")
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Output: 1
}

func ExampleHandleOptionIgnorNone() {
	input := option.None[int]()
	err := option.HandleOptionIgnoreNone(input, func(i int) error {
		fmt.Printf("%d", i)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Output:
}

func ExampleMapOption() {
	f := func(i int) string { return "i:" + strconv.Itoa(i) }
	fo := option.Map(f)
	input := []option.Option[int]{option.None[int](), option.Some(1)}
	r := functional.Map(fo, input)
	fmt.Println(r)
	// Output: [None Some(i:1)]
}
