package homework

func reverse(input []int64) (result []int64) {
	if len(input) == 0 {
		return input
	}
	return append(reverse(input[1:]), input[0])
}
