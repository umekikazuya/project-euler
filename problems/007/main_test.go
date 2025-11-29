package main

import "testing"

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve()
	}
}

func Test_solve(t *testing.T) {
	t.Parallel()
	got := solve()
	if got != 104743 {
		t.Errorf("solve() = %v, want %v", got, 104743)
	}
}

func Test_isPrimeNumber(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		x    int32
		want bool
	}{
		{
			name: "case 1",
			x:    1,
			want: false,
		},
		{
			name: "case 2",
			x:    10,
			want: false,
		},
		{
			name: "case 3",
			x:    9,
			want: false,
		},
		{
			name: "case 4",
			x:    11,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := isPrimeNumber(tt.x)
			if got != tt.want {
				t.Errorf("isPrimeNumber(%d) = %v, want %v", tt.x, got, tt.want)
			}
		})
	}
}
