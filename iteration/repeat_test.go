package iteration

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", rand.Intn(1000))
	}
}

func ExampleRepeat() {
	repeated := Repeat("ha", 5)
	fmt.Println(repeated)
	// Output: hahahahaha
}
