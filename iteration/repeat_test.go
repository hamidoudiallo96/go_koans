package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Repeated %q expected %q", repeated, expected)
	}
}

func Benchmark(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
