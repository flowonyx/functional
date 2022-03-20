package list

// Create makes a slice of the given length with each item set to value.
func Create[T any](length int, value T) []T {
	output := make([]T, length)
	for i := range output {
		output[i] = value
	}
	return output
}

// Create2D makes a two dimensional slice of the given lengths with each item set to value.
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

// Create3D makes a three dimensional slice of the given lengths with each item set to value.
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

// Create4D makes a four dimensional slice of the given lengths with each item set to value.
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

// ZeroCreate2D makes a two dimensional slice of the given lengths.
func ZeroCreate2D[T any](length1 int, length2 int) [][]T {
	output := make([][]T, length1)
	for i := range output {
		output[i] = make([]T, length2)
	}
	return output
}

// ZeroCreate3D makes a three dimensional slice of the given lengths.
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

// ZeroCreate4D makes a four dimensional slice of the given lengths.
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

// CreateFromStructure2D makes a two dimensional slice sized according the lengths in structure.
func CreateFromStructure2D[T, R any](structure [][]T) [][]R {
	output := make([][]R, len(structure))
	Iteri(func(i int, t []T) { output[i] = make([]R, len(t)) }, structure)
	return output
}

// CreateFromStructure3D makes a three dimensional slice sized according the lengths in structure.
func CreateFromStructure3D[T, R any](structure [][][]T) [][][]R {
	output := make([][][]R, len(structure))
	Iteri(func(i int, t [][]T) {
		output[i] = make([][]R, len(t))
		Iteri(func(j int, t []T) {
			output[i][j] = make([]R, len(t))
		}, structure[i])
	}, structure)
	return output
}
