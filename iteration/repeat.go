package iteration

func Repeat(str string) string {
	var res string
	for i := 0; i < 5; i++ {
		res += str
	}
	return res
}
