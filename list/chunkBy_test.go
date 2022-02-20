package list_test

import (
	"fmt"

	"github.com/flowonyx/functional/list"
)

func ExampleChunkBySize() {
	chunks := list.ChunkBySize(2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(chunks)
	// Output: [[1 2] [3 4] [5 6] [7 8] [9]]
}
