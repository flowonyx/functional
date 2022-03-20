package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func mod2(i int) bool { return i%2 == 0 }
func mod6(i int) bool { return i%6 == 0 }

func ExampleFind() {
	r, err := list.Find(mod2, 1, 2, 3, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: 2
}

func Example2Find() {
	_, err := list.Find(mod6, 1, 2, 3)
	fmt.Println(err)
	// Output: key not found
}

func ExampleFindBack() {
	r, err := list.FindBack(mod2, 1, 2, 3, 4)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
	// Output: 4
}

func ExampleTryFind() {
	r := list.TryFind(mod2, 1, 2, 3, 4)
	if r.IsNone() {
		panic("TryFind should have found 2")
	}
	fmt.Println(r.Value())
	// Output: 2
}

func ExampleTryFindBack() {
	r := list.TryFindBack(mod2, 1, 2, 3, 4)
	if r.IsNone() {
		panic("TryFind should have found 4")
	}
	fmt.Println(r.Value())
	// Output: 4
}
