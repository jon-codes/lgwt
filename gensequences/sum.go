package gensequences

func Sum(numbers []int) int {
	return Reduce(numbers, func(current, acc int) int {
		return acc + current
	}, 0)
}

func SumAll(toSum ...[]int) []int {
	return Reduce(toSum, func(current, acc []int) []int {
		return append(acc, Sum(current))
	}, []int{})
}

func SumAllTails(toSum ...[]int) []int {
	return Reduce(toSum, func(current, acc []int) []int {
		if len(current) == 0 {
			return append(acc, 0)
		} else {
			tail := current[1:]
			return append(acc, Sum(tail))
		}
	}, []int{})
}

func Reduce[T, R any](list []T, fn func(current T, acc R) R, initial R) R {
	result := initial
	for _, el := range list {
		result = fn(el, result)
	}
	return result
}

func Find[T any](list []T, fn func(el T) bool) (value T, found bool) {
	for _, el := range list {
		if fn(el) {
			return el, true
		}
	}
	return value, found
}
