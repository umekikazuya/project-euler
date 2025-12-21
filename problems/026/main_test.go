package main

import "testing"

// BenchmarkSolve は解答結果を求める処理を算出する処理のベンチマーク
func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve(1000)
	}
}

// Test_Solve は解答結果を求める処理を算出する処理のテスト
func Test_Solve(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "case1", n: 10, want: 6},
		{name: "case2", n: 1000, want: 983},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.n)
			if got != tt.want {
				t.Errorf("solve() = %d, want %d", got, tt.want)
			}
		})
	}
}

// Test_GetDiv
func Test_GetDiv(t *testing.T) {
	tests := []struct {
		name string
		n    int
		d    int
		want int
	}{
		{name: "case1", n: 3, d: 7, want: 2},
		{name: "case2", n: 4, d: 7, want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getDiv(tt.n, tt.d)
			if got != tt.want {
				t.Errorf("getDiv() = %d, want %d", got, tt.want)
			}
		})
	}
}

// Test_GetDiv
func Test_GetReciprocalCycles(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{name: "case1", n: 2, want: 0},
		{name: "case2", n: 3, want: 1},
		{name: "case3", n: 4, want: 0},
		{name: "case4", n: 5, want: 0},
		{name: "case5", n: 6, want: 1},
		{name: "case6", n: 7, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getReciprocalCycles(tt.n)
			if got != tt.want {
				t.Errorf("getReciprocalCycles() = %d, want %d", got, tt.want)
			}
		})
	}
}
