package p001

import (
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
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
			want: 233168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solve(); got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
