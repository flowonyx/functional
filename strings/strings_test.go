package strings

import "fmt"

func ExampleCollect() {
	input := "12345"
	r := Collect(func(r rune) string {
		i, err := ToInt(r)
		if err != nil {
			return err.Error()
		}
		return Repeat(r, i)
	}, input)
	fmt.Println(r)
	// Output: 122333444455555
}

func ExampleCollecti() {
	input := "12345"
	r := Collecti(func(i int, r rune) string {
		return Repeat(r, i)
	}, input)
	fmt.Println(r)
	// Output: 2334445555
}

func ExampleConcat() {
	input := []string{
		"hello",
		" ",
		"world",
	}
	r := Concat(input)
	fmt.Println(r)
	// Output: hello world
}

func ExampleExists() {
	input := "12345"
	ft := func(r rune) bool { return r < '5' }
	ff := func(r rune) bool { return r > '5' }
	r1 := Exists(ft, input)
	r2 := Exists(ff, input)
	fmt.Printf("%t, %t", r1, r2)
	// Output: true, false
}

func ExampleFilter() {
	input := "12345"
	r := Filter(func(r rune) bool { return r < '4' }, input)
	fmt.Println(r)
	// Output: 123
}

func ExampleInitString() {
	r := InitString(5, func(i int) string { return FromInt(i) })
	fmt.Println(r)
	// Output: 01234
}

func ExampleIter() {
	input := "12345"
	s := ""
	Iter(func(r rune) {
		s += Repeat(r, 2)
	}, input)
	fmt.Println(s)
	// Output: 1122334455
}

func ExampleIteri() {
	input := "12345"
	s := ""
	Iteri(func(i int, r rune) {
		s += Repeat(r, i)
	}, input)
	fmt.Println(s)
	// Output: 2334445555
}

func ExampleJoin() {
	input := []int{1, 2, 3}
	s := Join(input, ", ")
	fmt.Println(s)
	// Output: 1, 2, 3
}

func ExampleJoinRunes() {
	input := []rune{'a', 'b', 'c'}
	s := JoinRunes(input, ", ")
	fmt.Println(s)
	// Output: a, b, c
}

func ExampleJoinStrings() {
	input := []string{"a", "b", "c"}
	s := JoinStrings(input, ", ")
	fmt.Println(s)
	// Output: a, b, c
}

func ExampleLastIndexRune() {
	input := "abcdeab"
	r := LastIndexRune(input, 'b')
	fmt.Println(r)
	// Output: 6
}

func ExampleLines() {
	input := "line1\nline2\rline3\r\nline4"
	r := Lines(input)
	fmt.Println(len(r), r)
	// Output: 4 [line1 line2 line3 line4]
}

func ExampleGetLine() {
	input := "line1\nline2\rline3\r\nline4"
	r, err := GetLine(input, 2)
	if err != nil {
		panic(err)
	}
	fmt.Println(Quote(r))
	// Output: "line3"
}

func ExampleGetLineOpt() {
	input := "line1\nline2\rline3\r\nline4"
	r := GetLineOpt(input, 4)
	fmt.Println(r)
	// Output: None
}

func ExampleMapi() {
	input := "12345"
	output := Mapi(func(i int, r rune) rune {
		if i > 2 {
			r = r + 13
		}
		return r
	}, input)
	fmt.Println(output)
	// Output: 123AB
}

func ExampleRange() {
	r := string(Range('a', 'd'))
	fmt.Println(r)
	// Output: abcd
}
