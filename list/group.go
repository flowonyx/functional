package list

import . "github.com/flowonyx/functional"

func GroupBy[T any, Key comparable](projection func(T) Key, input []T) []Pair[Key, []T] {
	output := []Pair[Key, []T]{}
	for i := range input {
		key := projection(input[i])
		o := TryIndexBy(func(p Pair[Key, []T]) bool { return p.First == key }, output)
		index := o.Value()
		a := PairOf(key, []T{input[i]})
		if o.IsSome() {
			a = output[index]
			a.Second = append(a.Second, input[i])
			output[index] = a
		} else {
			output = append(output, a)
		}
	}
	return output
}

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
