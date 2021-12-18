package common

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetIntArrayFromFile(file_name string) (input []int) {
	example, err := os.ReadFile(file_name)
	check(err)
	scanner := bufio.NewScanner(strings.NewReader(string(example)))
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		input = append(input, value)
	}
	return input
}
