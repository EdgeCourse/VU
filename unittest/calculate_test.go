package calculate

import "testing"

func TestAdd(t *testing.T) {

	answer := Add(2, 3)
	expected := 5

	if answer != expected {
		t.Errorf("result %q, expected %q", answer, expected)
	}
}

/*
The benchmark function must run the target code b.N times, where N is an integer that can be adjusted. During benchmark execution, b.N is adjusted until the benchmark function lasts long enough to be timed reliably. The --bench flag accepts its arguments in the form of a regular expression.

This benchmarks an individual function:
go test -bench=Subtract
*/

func BenchmarkSubtract(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Subtract(4, 6)
	}
}

//go test -cover
