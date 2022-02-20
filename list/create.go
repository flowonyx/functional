package list

func Create[T any](length int, value T) []T {
	output := make([]T, length)
	for i := range output {
		output[i] = value
	}
	return output
}

func Create2D[T any](length1, length2 int, value T) [][]T {
	output := make([][]T, length1)
	for i := range output {
		output[i] = make([]T, length2)
		for j := range output[i] {
			output[i][j] = value
		}
	}
	return output
}

func Create3D[T any](length1, length2, length3 int, value T) [][][]T {
	output := make([][][]T, length1)
	for i := range output {
		output[i] = make([][]T, length2)
		for j := range output[i] {
			output[i][j] = make([]T, length3)
			for h := range output[i][j] {
				output[i][j][h] = value
			}
		}
	}
	return output
}

func Create4D[T any](length1, length2, length3, length4 int, value T) [][][][]T {
	output := make([][][][]T, length1)
	for i := range output {
		output[i] = make([][][]T, length2)
		for j := range output[i] {
			output[i][j] = make([][]T, length3)
			for h := range output[i][j] {
				output[i][j][h] = make([]T, length4)
				for k := range output[i][j][h] {
					output[i][j][h][k] = value
				}
			}
		}
	}
	return output
}

func ZeroCreate[T any](length int) []T {
	return make([]T, length)
}

func ZeroCreate2D[T any](length1 int, length2 int) [][]T {
	output := make([][]T, length1)
	for i := range output {
		output[i] = make([]T, length2)
	}
	return output
}

func ZeroCreate3D[T any](length1, length2, length3 int) [][][]T {
	output := make([][][]T, length1)
	for i := range output {
		output[i] = make([][]T, length2)
		for j := range output[i] {
			output[i][j] = make([]T, length3)
		}
	}
	return output
}

func ZeroCreate4D[T any](length1, length2, length3, length4 int) [][][][]T {
	output := make([][][][]T, length1)
	for i := range output {
		output[i] = make([][][]T, length2)
		for j := range output[i] {
			output[i][j] = make([][]T, length3)
			for h := range output[i][j] {
				output[i][j][h] = make([]T, length4)
			}
		}
	}
	return output
}

func CreateFromStructure2D[T, R any](input [][]T) [][]R {
	output := make([][]R, len(input))
	Iteri(func(i int, t []T) { output[i] = make([]R, len(t)) }, input)
	return output
}

func CreateFromStructure3D[T, R any](input [][][]T) [][][]R {
	output := make([][][]R, len(input))
	Iteri(func(i int, t [][]T) {
		output[i] = make([][]R, len(t))
		Iteri(func(j int, t []T) {
			output[i][j] = make([]R, len(t))
		}, input[i])
	}, input)
	return output
}
