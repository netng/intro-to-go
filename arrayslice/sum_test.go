package arrayslice

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func ExampleSum() {
	numbers := [5]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	fmt.Println(got)
	// Output: 15

}
