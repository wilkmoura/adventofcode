package one

func DayOnePartOne(input []int) (count int) {
	previous := input[0]
	count = 0
	for idx := 1; idx < len(input); idx++ {
		current := input[idx]
		if current > previous {
			count++
		}
		previous = current
	}
	return
}

func DayOnePartTwo(input []int) (count int) {
	count = 0
	for idx := 3; idx < len(input); idx++ {
		previous := input[idx-3] + input[idx-2] + input[idx-1]
		current := input[idx-2] + input[idx-1] + input[idx]
		if current > previous {
			count++
		}
	}
	return
}
