package sequences

func Sum(numbers []int) int {
	s := 0
	for _, n := range numbers {
		s += n
	}
	return s
}

func SumAll(toSum ...[]int) []int {
	s := []int{}
	for _, a := range toSum {
		s = append(s, Sum(a))
	}
	return s
}

func SumAllTails(toSum ...[]int) []int {
	s := []int{}
	for _, a := range toSum {
		if len(a) == 0 {
			s = append(s, 0)
		} else {
			tail := a[1:]
			s = append(s, Sum(tail))
		}
	}
	return s
}
