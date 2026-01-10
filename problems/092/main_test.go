package main

import (
	"testing"
)

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
		{name: "case", want: 8581146},
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

func Test_isEigtyNineSquareDigit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  int
		want bool
	}{
		{name: "case", arg: 44, want: false},
		{name: "case", arg: 85, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isEigtyNineSquareDigit(tt.arg)
			if got != tt.want {
				t.Errorf("isEigtyNineSquareDigit() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func Test_sumSquareDigit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{name: "case", arg: 44, want: 32},
		{name: "case", arg: 85, want: 89},
		{name: "case", arg: 13, want: 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumSquareDigit(tt.arg)
			if got != tt.want {
				t.Errorf("sumSquareDigit() = %v, want = %v", got, tt.want)
			}
		})
	}
}