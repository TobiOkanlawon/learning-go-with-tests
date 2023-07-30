package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size,", func(t *testing.T) {
		got := Sum([]int{1, 2, 3, 4})
		want := 1 + 2 + 3 + 4

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v; want %v", got, want)
		}

	})
}

func TestSumAll(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v; want %v", got, want)
		}

	}

	t.Run("returns a slice of totals for each slice passed in", func(t *testing.T) {

		got := SumAll([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
		want := []int{6, 9, 12}

		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v; want %v", got, want)
		}

	}

	t.Run("returns the correct value for tail addition", func(t *testing.T) {

		got := SumAllTails([]int{1, 2, 3}, []int{2, 3, 4}, []int{3, 4, 5})
		want := []int{5, 7, 9}

		checkSums(t, got, want)
	})

	t.Run("returns the correct value for an empty  slice", func(t *testing.T) {

		var emptySlice []int
		got := SumAllTails(emptySlice)
		want := []int{0}

		checkSums(t, got, want)
	})

	t.Run("returns the correct value for multiple empty slices", func(t *testing.T) {

		var emptySlice []int
		got := SumAllTails(emptySlice, emptySlice, []int{1, 2, 3})
		want := []int{0, 0, 5}

		checkSums(t, got, want)

	})

}
