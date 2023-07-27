package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("t", 5)
	expected := "ttttt"

	if repeated != expected {
		t.Errorf("Return value: %q is not equal to expcted value %q", repeated, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 50000)
	}
}

func BenchmarkOldRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 50000000)
	}
}

func ExampleRepeat() {
	repeated_string := Repeat("r", 3)
	fmt.Println(repeated_string)
	// Output: rrr
}
