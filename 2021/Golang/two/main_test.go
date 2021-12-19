package two

import (
	"testing"
	"wilkmoura/adventofcode/common"
)

func TestDayTwoExample(t *testing.T) {
	t.Run("Part One", func(t *testing.T) {
		want := 150
		got := DayTwoPartOne(common.GetStringArrayFromFile("./input_example.txt"))

		if got != want {
			t.Errorf("not the right solution got %d want %d", got, want)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		want := 900
		got := DayTwoPartTwo(common.GetStringArrayFromFile("./input_example.txt"))

		if got != want {
			t.Errorf("not the right solution got %d want %d", got, want)
		}
	})
}

func TestDayTwo(t *testing.T) {
	t.Run("Part One", func(t *testing.T) {
		want := 2272262
		got := DayTwoPartOne(common.GetStringArrayFromFile("./input.txt"))

		if got != want {
			t.Errorf("not the right solution got %d want %d", got, want)
		}
	})

	t.Run("Part Two", func(t *testing.T) {
		want := 2134882034
		got := DayTwoPartTwo(common.GetStringArrayFromFile("./input.txt"))

		if got != want {
			t.Errorf("not the right solution got %d want %d", got, want)
		}
	})
}
