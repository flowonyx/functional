package result_test

import (
	"fmt"

	"github.com/flowonyx/functional/result"
)

func ExampleHandleResult() {
	res := result.OK(1)
	err := result.HandleResult(res, func(i int) error {
		fmt.Println(i)
		return nil
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
	if err != nil {
		panic(err)
	}
	// Output: 1
}

func ExampleBind() {
	input := result.OK(2)
	res := result.Bind(func(t int) result.Result[int] { return result.OK(t * 2) }, input)
	if res.IsError() {
		panic(res.Err())
	}
	if !res.IsOK() {
		panic("not ok")
	}
	fmt.Println(res.Value())
	// Output: 4
}

func Example2Bind() {
	input := result.OK(2)
	res := result.Bind(func(t int) result.Result[int] { return result.Error[int](fmt.Errorf("error: %d", t)) }, input)
	res = result.Bind(func(t int) result.Result[int] { return result.OK(t) }, res)
	if res.IsOK() {
		panic("should not be ok")
	}
	if !res.IsError() {
		panic("should be error")
	}
	fmt.Println(res.Err())
	// Output: error: 2
}

func ExampleMapResult() {
	input := result.OK(1)
	res := result.Map(func(i int) string { return fmt.Sprint("input: ", i) }, input)
	fmt.Println(res.Value())
	// Output: input: 1
}
