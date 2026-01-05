package main

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve()
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_solve(t *testing.T) {
	tests := []struct {
		name string
		want int32
	}{
		{name: "case", want: 402},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_getFactorialSum(t *testing.T) {
	tests := []struct {
		name string
		args int32
		want int32
	}{
		{
			name: "case1",
			args: 145,
			want: 145,
		},
		{
			name: "case2",
			args: 363601,
			want: 1454,
		},
		{
			name: "case3",
			args: 540,
			want: 145,
		},
	}
	cache := generateFactorialMap()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := getFactorialSum(tt.args, cache)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_generateFactorialMap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		target int32
		want   int32
	}{
		{name: "case1", target: 0, want: 1},
		{name: "case2", target: 1, want: 1},
		{name: "case3", target: 2, want: 2},
		{name: "case4", target: 3, want: 6},
		{name: "case5", target: 4, want: 24},
		{name: "case6", target: 5, want: 120},
		{name: "case7", target: 6, want: 720},
		{name: "case8", target: 7, want: 5040},
		{name: "case9", target: 8, want: 40320},
		{name: "case10", target: 9, want: 362880},
	}
	got := generateFactorialMap()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, got[tt.target])
		})
	}
}
