package main

import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve()
	}
}

func Test_solve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		want int
	}{
		{
			name: "case1",
			want: 5777,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := solve()
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPrime(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name  string
		n     int
		cache map[int]bool
		want  bool
	}{
		{
			name:  "case1",
			n:     2,
			cache: nil,
			want:  true,
		},
		{
			name:  "case2",
			n:     9,
			cache: nil,
			want:  false,
		},
		{
			name:  "case3",
			n:     11,
			cache: nil,
			want:  true,
		},
		{
			name:  "case4",
			n:     23,
			cache: nil,
			want:  true,
		},
		{
			name: "case5(cache)",
			n:    23,
			cache: map[int]bool{
				1:  false,
				2:  true,
				3:  false,
				23: true,
			},
			want: true,
		},
		{
			name: "case6",
			n:    9739,
			cache: map[int]bool{
				1:  false,
				2:  true,
				3:  false,
				23: true,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := isPrime(tt.n, tt.cache)
			if got != tt.want {
				t.Errorf("isPrime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSqrt(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		i    int
		want int
	}{
		{
			name: "case1",
			i:    10,
			want: 3,
		},
		{
			name: "case2",
			i:    100,
			want: 10,
		},
		{
			name: "case3",
			i:    25,
			want: 5,
		},
		{
			name: "case4",
			i:    26,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := getSqrt(tt.i)
			if got != tt.want {
				t.Errorf("getSqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
