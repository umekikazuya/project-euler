package main

import (
	"math/big"
	"testing"
)

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
			name: "case",
			want: 608720, // TODO: 答えが出た後に入力し直す
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// got := solve()
			// if got != tt.want {
			// 	t.Errorf("solve() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func Test_isMutableNumber(t *testing.T) {
	tests := []struct {
		name string
		arg  big.Int
		want bool
	}{
		{name: "case1", arg: *big.NewInt(36), want: true},
		{name: "case2", arg: *big.NewInt(37), want: false},
		{name: "case3", arg: *big.NewInt(12), want: true},
		{name: "case4", arg: *big.NewInt(13), want: false},
		{name: "case5", arg: *big.NewInt(80), want: false},
		{name: "case6", arg: *big.NewInt(15), want: false},
		{name: "case7", arg: *big.NewInt(16), want: true},
		{name: "case8", arg: *big.NewInt(5), want: false},
		{name: "case9", arg: *big.NewInt(6), want: false},
		{name: "case10", arg: *big.NewInt(19), want: false},
		{name: "case11", arg: *big.NewInt(99), want: false},
		{name: "case12", arg: *big.NewInt(10), want: false},
		{name: "case13", arg: *big.NewInt(409), want: true},
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

func Test_reverseBigInt(t *testing.T) {
	tests := []struct {
		name string
		arg  big.Int
		want big.Int
	}{
		{
			name: "case1",
			arg:  *big.NewInt(123),
			want: *big.NewInt(321),
		},
		{name: "case2", arg: *big.NewInt(1234), want: *big.NewInt(4321)},
		{name: "case3", arg: *big.NewInt(2000), want: *big.NewInt(2)},
		{name: "case4", arg: *big.NewInt(12), want: *big.NewInt(21)},
		{name: "case5", arg: *big.NewInt(32), want: *big.NewInt(23)},
		{name: "case6", arg: *big.NewInt(89), want: *big.NewInt(98)},
		{name: "case7", arg: *big.NewInt(76), want: *big.NewInt(67)},
		{name: "case7", arg: *big.NewInt(19), want: *big.NewInt(91)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := reverseBigInt(tt.arg)
			if err != nil {
				t.Errorf("reverseBigInt() = %v, err %v", err, err)
			}
			if got.Text(10) != tt.want.Text(10) {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sumWithReverseNumber(t *testing.T) {
	tests := []struct {
		name string
		arg  big.Int
		want big.Int
	}{
		{name: "case1", arg: *big.NewInt(36), want: *big.NewInt(99)},
		{name: "case2", arg: *big.NewInt(44), want: *big.NewInt(88)},
		{name: "case3", arg: *big.NewInt(20), want: *big.NewInt(22)},
		{name: "case4", arg: *big.NewInt(12), want: *big.NewInt(33)},
		{name: "case5", arg: *big.NewInt(15), want: *big.NewInt(66)},
		{name: "case6", arg: *big.NewInt(16), want: *big.NewInt(77)},
		{name: "case7", arg: *big.NewInt(19), want: *big.NewInt(110)},
		{name: "case8", arg: *big.NewInt(95), want: *big.NewInt(154)},
		{name: "case9", arg: *big.NewInt(28), want: *big.NewInt(110)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := sumWithReverseNumber(tt.arg)
			if got.String() != tt.want.String() {
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
