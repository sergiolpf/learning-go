package arrays

import (
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of 5 items", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d, %v", got, want, numbers)
		}
	})
}

/*
func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})

	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}


func TestSumAllTails(t *testing.T) {
	func TestSumAllTails(t *testing.T) {
	got := SumAllTalis([]int{1, 2}, []int{0, 9})

	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

}
*/
