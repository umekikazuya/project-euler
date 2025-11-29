package main

import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve()
	}
}

func Test_solve(t *testing.T) {
	got := solve()
	if got != 23514624000 {
		t.Errorf("solve() = %v, want %v", got, 23514624000)
	}
}
