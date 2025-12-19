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
	tests := []struct {
		name string
		want int
	}{
		{
			name: "case",
			want: 608720,
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

func Test_isMutableNumber(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want bool
	}{
		{
			name: "case1", arg: 36, want: true,
		},
		{
			name: "case2", arg: 37, want: false,
		},
		{
			name: "case3", arg: 12, want: true,
		},
		{
			name: "case4", arg: 13, want: false,
		},
		{
			name: "case5", arg: 80, want: false,
		},
		{
			name: "case6", arg: 15, want: false,
		},
		{
			name: "case7", arg: 16, want: true,
		},
		{
			name: "case8", arg: 5, want: false,
		},
		{
			name: "case9", arg: 6, want: false,
		},
		{
			name: "case10", arg: 19, want: false,
		},
		{
			name: "case11", arg: 99, want: false,
		},
		{
			name: "case12", arg: 10, want: false,
		},
		{
			name: "case13", arg: 409, want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := isMutableNumber(tt.arg)
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseInt(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{
			name: "case1",
			arg:  123,
			want: 321,
		},
		{name: "case2", arg: 1234, want: 4321},
		{name: "case3", arg: 2000, want: 2},
		{name: "case4", arg: 12, want: 21},
		{name: "case5", arg: 32, want: 23},
		{name: "case6", arg: 89, want: 98},
		{name: "case7", arg: 76, want: 67},
		{name: "case7", arg: 19, want: 91},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := reverseInt(tt.arg)
			if err != nil {
				t.Errorf("reverseInt() = %v, err %v", err, err)
			}
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumWithReverseNumber(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want int
	}{
		{name: "case1", arg: 36, want: 99},
		{name: "case2", arg: 44, want: 88},
		{name: "case3", arg: 20, want: 22},
		{name: "case4", arg: 12, want: 33},
		{name: "case5", arg: 15, want: 66},
		{name: "case6", arg: 16, want: 77},
		{name: "case7", arg: 19, want: 110},
		{name: "case8", arg: 95, want: 154},
		{name: "case9", arg: 28, want: 110},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := sumWithReverseNumber(tt.arg)
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		arg  string
		want string
	}{
		{
			name: "case1",
			arg:  "abc",
			want: "cba",
		},
		{name: "case2", arg: "012", want: "210"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := reverse(tt.arg)
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
