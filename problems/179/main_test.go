package main

import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve()
	}
}

func Test_solve(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		want int
	}{
		{
			name: "case",
			want: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve()
			if got != tt.want {
				t.Errorf("solve() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_numuberOfDivisors(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{name: "case", arg: 1, want: 1},
		{name: "case", arg: 2, want: 2},
		{name: "case", arg: 14, want: 4},
		{name: "case", arg: 15, want: 4},
		{name: "case", arg: 99, want: 6},
		{name: "case", arg: 100, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numuberOfDivisors(tt.arg)
			if got != tt.want {
				t.Errorf("numuberOfDivisors(%v) = %v, want = %v", tt.arg, got, tt.want)
			}
		})
	}
}
