package one

import (
	"testing"
	"wilkmoura/adventofcode/common"
)

func TestDayOneExample(t *testing.T) {
	t.Run("Part One", func(t *testing.T) {
		want := 7
		got := DayOnePartOne(common.GetIntArrayFromFile("./input_example.txt"))

		if got != want {
			t.Errorf("not the right solution got %d want %d", got, want)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		want := 5
		got := DayOnePartTwo(common.GetIntArrayFromFile("./input_example.txt"))

		if got != want {
			t.Errorf("not the right solution got %d want %d", got, want)
		}
	})
}

func TestDayOne(t *testing.T) {
	t.Run("Part One", func(t *testing.T) {
		want := 1195
		got := DayOnePartOne(common.GetIntArrayFromFile("./input.txt"))

		if got != want {
			t.Errorf("not the right solution got %d want %d", got, want)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		want := 1235
		got := DayOnePartTwo(common.GetIntArrayFromFile("./input.txt"))

		if got != want {
			t.Errorf("not the right solution got %d want %d", got, want)
		}
	})
}
