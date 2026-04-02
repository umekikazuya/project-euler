package main

import "testing"

func BenchmarkSolve(b *testing.B) {
	for b.Loop() {
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
			want: 17427258,
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

func Test_gerneratePrimes(t *testing.T) {
	got := generatePrimes()
	tests := []struct {
		name string
		arg  int
		want bool // 素数はfalse
	}{
		{
			name: "0",
			arg:  0,
			want: true,
		},
		{
			name: "1",
			arg:  1,
			want: true,
		},
		{
			name: "2",
			arg:  2,
			want: false,
		},
		{
			name: "3",
			arg:  3,
			want: false,
		},
		{
			name: "10",
			arg:  10,
			want: true,
		},
		{
			name: "29",
			arg:  29,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got[tt.arg] != tt.want {
				t.Errorf("arg = %v, want = %v", tt.arg, tt.want)
			}
		})
	}
}
