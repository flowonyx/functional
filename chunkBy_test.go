package functional_test

import (
	"fmt"

	"github.com/flowonyx/functional"
)

func ExampleChunkBySize() {
	chunks := functional.ChunkBySize(2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(chunks)
	// Output: [[1 2] [3 4] [5 6] [7 8] [9]]
}
