package strings

import "fmt"

func ExampleFromBool() {
	t := FromBool(true)
	f := FromBool(false)
	fmt.Printf("%s, %s", t, f)
	// Output: true, false
}

func ExampleToBool() {
	t1, err := ToBool("1")
	if err != nil {
		panic(err)
	}
	t2, err := ToBool("t")
	if err != nil {
		panic(err)
	}
	t3, err := ToBool("true")
	if err != nil {
		panic(err)
	}
	t4, err := ToBool("True")
	if err != nil {
		panic(err)
	}

	f1, err := ToBool("0")
	if err != nil {
		panic(err)
	}
	f2, err := ToBool("f")
	if err != nil {
		panic(err)
	}
	f3, err := ToBool("false")
	if err != nil {
		panic(err)
	}
	f4, err := ToBool("False")
	if err != nil {
		panic(err)
	}

	_, err = ToBool("junk")
	if err == nil {
		panic("should have errored")
	}

	fmt.Printf("%t, %t, %t, %t | %t, %t, %t, %t", t1, t2, t3, t4, f1, f2, f3, f4)
	// Output: true, true, true, true | false, false, false, false
}

func ExampleToBoolOpt() {
	t1 := ToBoolOpt("1")
	t2 := ToBoolOpt("t")
	t3 := ToBoolOpt("true")
	t4 := ToBoolOpt("True")
	f1 := ToBoolOpt("0")
	f2 := ToBoolOpt("f")
	f3 := ToBoolOpt("false")
	f4 := ToBoolOpt("False")

	b := ToBoolOpt("junk")

	fmt.Printf("%s, %s, %s, %s | %s, %s, %s, %s | %s", t1.String(), t2.String(), t3.String(), t4.String(), f1.String(), f2.String(), f3.String(), f4.String(), b.String())
	// Output: Some(true), Some(true), Some(true), Some(true) | Some(false), Some(false), Some(false), Some(false) | None
}

func ExampleFromInt() {
	s1 := FromInt(-1)
	s2 := FromInt(1025)

	fmt.Printf("%s | %s", s1, s2)
	// Output: -1 | 1025
}

func ExampleToInt() {
	i1, err := ToInt("-1")
	if err != nil {
		panic(err)
	}
	i2, err := ToInt("1025")
	if err != nil {
		panic(err)
	}
	i3, err := ToInt8("127")
	if err != nil {
		panic(err)
	}
	i4, err := ToInt8("128")
	if err == nil {
		panic("should have error")
	}

	fmt.Printf("%d, %d, %d, %d", i1, i2, i3, i4)
	// Output: -1, 1025, 127, 127
}

func ExampleToIntOpt() {
	i1 := ToIntOpt("-1")
	i2 := ToIntOpt("junk")
	fmt.Printf("%s, %s", i1, i2)
	// Output: Some(-1), None
}

func ExampleFromFloat() {
	s := FromFloat(123.4)
	fmt.Println(s)
	// Output: 123.4
}

func ExampleToFloat() {
	f1, err := ToFloat("123.4")
	if err != nil {
		panic(err)
	}
	f2, err := ToFloat("5")
	if err != nil {
		panic(err)
	}
	_, err = ToFloat("junk")
	if err == nil {
		panic("should have errored")
	}
	fmt.Printf("%.2f, %.2f", f1, f2)
	// Output: 123.40, 5.00
}

func ExampleToFloatOpt() {
	f1 := ToFloatOpt("123.4")
	f2 := ToFloatOpt("5")
	b := ToFloatOpt("junk")

	fmt.Printf("%.2f, %.2f, %s", f1.Value(), f2.Value(), b)
	// Output: 123.40, 5.00, None
}

func ExampleToDate() {
	d, err := ToDate("2022-3-14")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", d.Format(MDYYYYFullName))
	// Output: March 14, 2022
}

func ExampleCamelCaseToUnderscore() {
	s := CamelCaseToUnderscore("exampleString")
	fmt.Print(s)
	// Output: example_string
}

func ExampleNormalizeNewLine() {
	s := NormalizeNewLine("line1\nline2\r\nline3\rline4")
	fmt.Print(Quote(s))
	// Output: "line1\nline2\nline3\nline4"
}
