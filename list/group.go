package list

import . "github.com/flowonyx/functional"

// GroupBy applies projection to each value in values and returns a slice of Pairs where
// each Pair has a key and a slice of values for which projection returned that key.
func GroupBy[T any, Key comparable](projection func(T) Key, values []T) []Pair[Key, []T] {
	output := []Pair[Key, []T]{}
	for i := range values {
		key := projection(values[i])
		o := TryIndexBy(func(p Pair[Key, []T]) bool { return p.First == key }, output)
		index := o.Value()
		a := PairOf(key, []T{values[i]})
		if o.IsSome() {
			a = output[index]
			a.Second = append(a.Second, values[i])
			output[index] = a
		} else {
			output = append(output, a)
		}
	}
	return output
}

// GroupAsMap applies projection to each value in values and returns a map where
// each item has a key and a slice of values for which projection returned that key.
func GroupByAsMap[T any, Key comparable](projection func(T) Key, input []T) map[Key][]T {
	output := map[Key][]T{}
	for i := range input {
		key := projection(input[i])
		a, ok := output[key]
		if !ok {
			a = []T{}
		}
		a = append(a, input[i])
		output[key] = a
	}
	return output
}
