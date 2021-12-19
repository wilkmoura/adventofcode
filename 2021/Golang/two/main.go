package two

import (
	"strconv"
	"strings"
)

const FORWARD = "forward"
const DOWN = "down"
const UP = "up"

func DayTwoPartOne(inputs []string) int {
	horizontal_position := 0
	vertical_position := 0

	for _, value := range inputs {
		pair := strings.Split(value, " ")
		moviment := pair[0]
		qty, _ := strconv.Atoi(pair[1])
		if moviment == FORWARD {
			horizontal_position += qty
		} else if moviment == DOWN {
			vertical_position += qty
		} else if moviment == UP {
			vertical_position -= qty
		}
	}
	return vertical_position * horizontal_position
}

func DayTwoPartTwo(inputs []string) (count int) {
	horizontal_position := 0
	vertical_position := 0
	aim := 0

	for _, value := range inputs {
		pair := strings.Split(value, " ")
		moviment := pair[0]
		qty, _ := strconv.Atoi(pair[1])
		if moviment == FORWARD {
			horizontal_position += qty
			vertical_position += aim * qty
		} else if moviment == DOWN {
			aim += qty
		} else if moviment == UP {
			aim -= qty
		}
	}
	return vertical_position * horizontal_position
}
