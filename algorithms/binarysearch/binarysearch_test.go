package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	input := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 19
	want := 9
	got := BinarySearch(input, target)

	if want != got {
		t.Errorf("got %d want %d", got, want)
	}

}
