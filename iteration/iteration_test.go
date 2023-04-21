package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")

	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but received %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}