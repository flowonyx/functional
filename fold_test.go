package functional_test

import (
	"fmt"
	"strconv"

	. "github.com/flowonyx/functional"
	"github.com/flowonyx/functional/option"
)

type ChargeType string
type CoinToss int

const (
	In  = "In"
	Out = "Out"
)

const (
	Heads = CoinToss(iota)
	Tails
)

func chargesFolder(acc int, charge Pair[int, ChargeType]) int {
	if charge.Second == In {
		return acc + charge.First
	}
	return acc - charge.First
}

func coinTossFolder(acc int, a, b CoinToss) int {
	if a == Heads && b == Heads {
		return acc + 1
	}
	if a == Tails && b == Tails {
		return acc + 1
	}
	return acc - 1
}

func ExampleFold() {
	r := Fold(chargesFolder, 0, []Pair[int, ChargeType]{
		{First: 1, Second: In},
		{First: 2, Second: Out},
		{First: 3, Second: In},
	})
	fmt.Println(r)
	// Output: 2
}

func ExampleFold2() {
	data1 := []CoinToss{Tails, Heads, Tails}
	data2 := []CoinToss{Tails, Heads, Heads}
	r := Fold2(coinTossFolder, 0, data1, data2)
	fmt.Println(r)
	// Output: 1
}

func ExampleFoldBack() {
	type Count struct {
		Positive int
		Negative int
		Text     string
	}

	countFolder := func(a int, acc Count) Count {
		text := acc.Text + " " + strconv.Itoa(a)
		if a >= 0 {
			return Count{
				Positive: acc.Positive + 1,
				Negative: acc.Negative,
				Text:     text,
			}
		}
		return Count{
			Positive: acc.Positive,
			Negative: acc.Negative + 1,
			Text:     text,
		}
	}

	input := []int{1, 0, -1, -2, 3}
	initialState := Count{Positive: 0, Negative: 0, Text: ""}
	r := FoldBack(countFolder, input, initialState)
	fmt.Println(r.Positive, r.Negative, ":"+r.Text)
	// Output: 3 2 : 3 -2 -1 0 1
}

func ExampleFoldBack2() {
	type Count struct {
		Positive int
		Negative int
		Text     string
	}

	countFolder := func(a, b int, acc Count) Count {
		text := acc.Text + "(" + strconv.Itoa(a) + ", " + strconv.Itoa(b) + ") "
		if a+b >= 0 {
			return Count{
				Positive: acc.Positive + 1,
				Negative: acc.Negative,
				Text:     text,
			}
		}
		return Count{
			Positive: acc.Positive,
			Negative: acc.Negative + 1,
			Text:     text,
		}
	}

	input1 := []int{-1, -2, -3}
	input2 := []int{3, 2, 1}
	initialState := Count{Positive: 0, Negative: 0, Text: ""}
	r := FoldBack2(countFolder, input1, input2, initialState)
	fmt.Println(r.Positive, r.Negative, ":"+r.Text)
	// Output: 2 1 :(-3, 1) (-2, 2) (-1, 3)
}

func ExampleUnfold() {
	r := Unfold(func(s int) option.Option[Pair[int, int]] {
		return IfV(s > 100, option.None[Pair[int, int]]()).
			Else(option.Some(PairOf(s, s*2)))
	}, 1)
	fmt.Println(r)
	// Output: [1 2 4 8 16 32 64]
}
