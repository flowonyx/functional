package result_test

import (
	"fmt"
	"strconv"

	"github.com/flowonyx/functional/errors"
	"github.com/flowonyx/functional/result"
)

func ExampleResult_IsSuccess() {
	r := result.Success[int, error](1)
	r2 := result.Failure[int](errors.BadArgumentErr)
	fmt.Println(r.IsSuccess(), r2.IsSuccess())
	// Output: true false
}

func ExampleResult_IsFailure() {
	r := result.Success[int, error](1)
	r2 := result.Failure[int](errors.BadArgumentErr)
	fmt.Println(r.IsFailure(), r2.IsFailure())
	// Output: false true
}

func ExampleResult_FailureValue() {
	r := result.Failure[int]("bad argument")
	r2 := result.Success[int, error](1)
	fmt.Println(strconv.Quote(r.FailureValue()), r2.FailureValue())
	// Output: "bad argument" <nil>
}

func ExampleHandleResult() {
	res := result.Success[int, error](1)
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
	input := result.Success[int, error](2)
	res := result.Bind(func(t int) result.Result[int, error] { return result.Success[int, error](t * 2) }, input)
	if res.IsFailure() {
		panic(res.FailureValue())
	}
	if !res.IsSuccess() {
		panic("not ok")
	}
	fmt.Println(res.SuccessValue())
	// Output: 4
}

func ExampleBind_second() {
	input := result.Success[int, error](2)
	res := result.Bind(func(t int) result.Result[int, error] { return result.Failure[int](fmt.Errorf("error: %d", t)) }, input)
	res = result.Bind(func(t int) result.Result[int, error] { return result.Success[int, error](t) }, res)
	if res.IsSuccess() {
		panic("should not be ok")
	}
	if !res.IsFailure() {
		panic("should be error")
	}
	fmt.Println(res.FailureValue())
	// Output: error: 2
}

func ExampleMap() {
	input := result.Success[int, string](1)
	res := result.Map(func(i int) string { return fmt.Sprint("input: ", i) }, input)
	fmt.Println(res.SuccessValue())
	// Output: input: 1
}

func ExampleMapError() {
	mapping := func(err error) error {
		return fmt.Errorf("%w: wrapped", err)
	}
	input := result.Success[int, error](1)
	input2 := result.Failure[int](fmt.Errorf("inner error"))
	res := result.MapError(mapping, input)
	res2 := result.MapError(mapping, input2)
	fmt.Println(res.String(), res2.String())
	// Output: 1 inner error: wrapped
}

func ExampleDefaultValue() {
	input := result.Success[int, error](1)
	input2 := result.Failure[int](fmt.Errorf("error"))
	r := result.DefaultValue(2, input)
	r2 := result.DefaultValue(2, input2)
	fmt.Println(r, r2)
	// Output: 1 2
}

func ExampleDefaultWith() {
	input := result.Success[int, string](1)
	input2 := result.Failure[int]("error")
	r := result.DefaultWith(func() int { return 2 }, input)
	r2 := result.DefaultWith(func() int { return 2 }, input2)
	fmt.Println(r, r2)
	// Output: 1 2
}

func ExampleContains() {
	input := result.Success[int, string](1)
	input2 := result.Failure[int]("error")
	fmt.Println(result.Contains(1, input), result.Contains(1, input2))
	// Output: true false
}

func ExampleCount() {
	input := result.Success[int, string](1)
	input2 := result.Failure[int]("error")
	fmt.Println(result.Count(input), result.Count(input2))
	// Output: 1 0
}

func ExampleExists() {
	input := result.Success[int, string](1)
	input2 := result.Failure[int]("error")
	r := result.Exists(func(value int) bool { return value > 2 }, input)
	r2 := result.Exists(func(value int) bool { return value < 2 }, input)
	r3 := result.Exists(func(value int) bool { return value < 2 }, input2)
	fmt.Println(r, r2, r3)
	// Output: false true false
}

func ExampleFold() {
	input := result.Success[int, string](5)
	r := result.Fold(func(state string, value int) string {
		return state + strconv.Itoa(value)
	}, "state:", input)
	fmt.Println(r)
	// Output: state:5
}

func ExampleFoldBack() {
	input := result.Success[int, string](5)
	r := result.FoldBack(func(value int, state string) string {
		return state + strconv.Itoa(value)
	}, input, "state:")
	fmt.Println(r)
	// Output: state:5
}

func ExampleForAll() {
	input1 := result.Success[int, string](1)
	input2 := result.Failure[int]("error")
	r1 := result.ForAll(func(v int) bool { return v < 2 }, input1)
	r2 := result.ForAll(func(v int) bool { return v > 2 }, input1)
	r3 := result.ForAll(func(v int) bool { return v < 2 }, input2)
	fmt.Println(r1, r2, r3)
	// Output: true false true
}
