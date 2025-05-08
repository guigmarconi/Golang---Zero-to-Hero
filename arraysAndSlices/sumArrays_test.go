package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size numbers", func(t *testing.T) {

		numbers := []int{1, 2, 3, 4, 5, 6, 7}

		got := Sum(numbers)
		want := 28

		if got != want {
			t.Errorf("Result: %d is different from expected: %d, given: %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{5, 10, 0}, []int{5, 0, 15})
	want := []int{15, 20}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Result: %v is different from expected: %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Result: %v is different from expected: %v", got, want)
		}

	}

	t.Run("Do the sums of the tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{5, 10}, []int{5, 0})
		want := []int{10, 0}

		checkSums(t, got, want)
	})

	t.Run("Do the sums of a empty slice safely", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 10})
		want := []int{0, 10}

		checkSums(t, got, want)
	})

}
