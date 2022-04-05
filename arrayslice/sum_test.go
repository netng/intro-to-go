package arrayslice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4}
	got := Sum(numbers)
	want := 10

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestSumAllTail(t *testing.T) {
	checkSum := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	}

	t.Run("sum all tail collection", func(t *testing.T) {
		got := SumAllTail([]int{1, 2}, []int{4, 5})
		want := []int{2, 5}
		checkSum(t, got, want)
	})

	t.Run("sum all tail when empty array provided", func(t *testing.T) {
		got := SumAllTail([]int{}, []int{4, 5})
		want := []int{0, 5}
		checkSum(t, got, want)
	})
}

func ExampleSum() {
	numbers := []int{1, 2, 3, 4}
	got := Sum(numbers)
	fmt.Println(got)
	// Output: 10
}

func ExampleSumAll() {
	got := SumAll([]int{}, []int{1, 2, 3})
	fmt.Println(got)
	// Output: [0 6]
}

func ExampleSumAllTail() {
	got := SumAllTail([]int{2, 5}, []int{1, 2})
	fmt.Println(got)
	// Output: [5 2]
}
