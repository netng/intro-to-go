package arrayslice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}
		got := Sum(numbers)
		want := 28

		if got != want {
			t.Errorf("got: %d, want: %d, %v", got, want, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)

	}
}

func TestSumAllTail(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}
	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{3, 5})
		want := []int{2, 5}
		checkSums(t, got, want)

	})

	t.Run("safely sums empty slices", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{3, 5})
		want := []int{0, 5}
		checkSums(t, got, want)
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5, 3}
	got := Sum(numbers)

	fmt.Println(got)
	// Output: 18
}

func ExampleSumAll() {
	got := SumAll([]int{1, 2, 3}, []int{1, 3})
	fmt.Println(got)
	// Output: [6 4]
}

func ExampleSumAllTail() {
	got := SumAllTail([]int{2, 5}, []int{1, 4})
	fmt.Println(got)
	// Output: [5 4]
}
