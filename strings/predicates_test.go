package strings

import "fmt"

func ExampleIsDate() {
	t := IsDate("2000-1-1")
	f := IsDate("1-1-1")
	fmt.Printf("%t, %t", t, f)
	// Output: true, false
}

func ExampleIsEmail() {
	t1 := IsEmail("test@test.com")
	t2 := IsEmail("\"Tester\" <test@test.com>")
	f := IsEmail("@handle")
	fmt.Printf("%t, %t, %t", t1, t2, f)
	// Output: true, true, false
}

func ExampleRuneIsAnyOf() {
	t := RuneIsAnyOf([]rune("123"), '2')
	f := RuneIsAnyOf([]rune("123"), '4')
	fmt.Printf("%t, %t", t, f)
	// Output: true, false
}

func ExampleRuneIsASCIILower() {
	t := RuneIsASCIILower('a')
	f := RuneIsASCIILower('A')
	fmt.Printf("%t, %t", t, f)
	// Output: true, false
}

func ExampleRuneIsASCIIUpper() {
	t := RuneIsASCIIUpper('A')
	f := RuneIsASCIIUpper('a')
	fmt.Printf("%t, %t", t, f)
	// Output: true, false
}

func ExampleRuneIsASCIILetter() {
	t1 := RuneIsASCIILetter('A')
	t2 := RuneIsASCIILetter('a')
	f := RuneIsASCIILetter('!')
	fmt.Printf("%t, %t, %t", t1, t2, f)
	// Output: true, true, false
}

func ExampleRuneIsHex() {
	t1 := RuneIsHex('0')
	t2 := RuneIsHex('9')
	t3 := RuneIsHex('F')
	t4 := RuneIsHex('f')
	f := RuneIsHex('G')
	fmt.Printf("%t, %t, %t, %t, %t", t1, t2, t3, t4, f)
	// Output: true, true, true, true, false
}

func ExampleRuneIsOctal() {
	t1 := RuneIsOctal('0')
	t2 := RuneIsOctal('7')
	f := RuneIsOctal('8')
	fmt.Printf("%t, %t, %t", t1, t2, f)
	// Output: true, true, false
}

func ExampleRuneIsNewLine() {
	t1 := RuneIsNewLine('\n')
	t2 := RuneIsNewLine('\r')
	f := RuneIsNewLine('\t')
	fmt.Printf("%t, %t, %t", t1, t2, f)
	// Output: true, true, false
}

func ExampleRuneIsNoneOf() {
	t := RuneIsNoneOf([]rune("123"), '4')
	f := RuneIsNoneOf([]rune("123"), '2')
	fmt.Printf("%t, %t", t, f)
	// Output: true, false
}
