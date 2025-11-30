package main

import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve()
	}
}

func BenchmarkSumDivisor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sumDivisor(220)
	}
}

func BenchmarkIsPrime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPrime(997)
	}
}

func Test_sumDivisor(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "case1", n: 28, want: 28},
		{name: "case2", n: 220, want: 284},
		{name: "case3", n: 284, want: 220},
		{name: "case4", n: 6, want: 6},
		{name: "case5", n: 7, want: 1},
		{name: "case6", n: 12496, want: 14288},
		{name: "case7", n: 14288, want: 15472},
		{name: "case8", n: 25, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sumDivisor(tt.n)
			if got != tt.want {
				t.Errorf("sumDivisor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{name: "case1", n: 1, want: false},
		{name: "case2", n: 2, want: true},
		{name: "case3", n: 3, want: true},
		{name: "case4", n: 5, want: true},
		{name: "case5", n: 7, want: true},
		{name: "case6", n: 10, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isPrime(tt.n)
			if got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}
