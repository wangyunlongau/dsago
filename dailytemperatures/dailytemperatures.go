package dailytemperatures

func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return []int{}
	}

	if len(T) == 1 {
		return []int{0}
	}

	var result = make([]int, len(T))
	var stack = []int{len(T) - 1}

	for i := len(T) - 2; i >= 0; i-- {
		for _, element := range stack {
			if T[i] >= T[element] {
				stack = stack[1:]
			} else {
				break
			}
		}

		//stack push
		stack = append([]int{i}, stack...)
		if len(stack) > 1 {
			result[i] = stack[1] - stack[0]
		}
	}
	return result
}
