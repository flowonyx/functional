package option_test

import (
	"fmt"

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
