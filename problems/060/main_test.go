package main

import (
	"slices"
	"testing"
)

// goos: darwin
// goarch: arm64
// pkg: github.com/umekikazuya/project-euler/problems/060
// cpu: Apple M2
// BenchmarkSolve-8   	       1	5897967792 ns/op	94265784 B/op	23394573 allocs/op
// PASS
// ok  	github.com/umekikazuya/project-euler/problems/060	6.148s


// goos: darwin
// goarch: arm64
// pkg: github.com/umekikazuya/project-euler/problems/060
// cpu: Apple M2
// BenchmarkSolve-8   	       1	6983458084 ns/op	436046688 B/op	112691233 allocs/// op
// PASS
// ok  	github.com/umekikazuya/project-euler/problems/060	7.344s


func BenchmarkSolve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solve()
	}
}

func Test_isPrimeNumber(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		args int
		want bool
	}{
		{
			name: "case 1",
			args: 10,
			want: false,
		},
		{
			name: "case 2",
			args: 13,
			want: true,
		},
		{
			name: "case 3",
			args: 23,
			want: true,
		},
		{
			name: "case 4",
			args: 37,
			want: true,
		},
		{
			name: "case 5",
			args: 256,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := isPrimeNumber(tt.args); got != tt.want {
				t.Errorf("isPrimeNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPrimeNumbers(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		max  int
		want []int
	}{
		{
			name: "case 1",
			max:  10,
			want: []int{2, 3, 5, 7},
		},
		{
			name: "case 2",
			max:  20,
			want: []int{2, 3, 5, 7, 11, 13, 17, 19},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := getPrimeNumbers(tt.max)
			if !slices.Equal(got, tt.want) {
				t.Errorf("getPrimeNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_concatenateNumbers(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		want2   int
		wantErr bool
	}{
		{
			name:    "case 1",
			a:       2,
			b:       3,
			want:    23,
			want2:   32,
			wantErr: false,
		},
		{
			name:    "case 2",
			a:       1,
			b:       3,
			want:    13,
			want2:   31,
			wantErr: false,
		},
		{
			name:    "case 3",
			a:       10,
			b:       3,
			want:    103,
			want2:   310,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got2 := concatenateNumbers(tt.a, tt.b)
			if tt.wantErr {
				t.Fatal("concatenateNumbers() succeeded unexpectedly")
			}
			if got != tt.want {
				t.Errorf("concatenateNumbers() = %v, want %v", got, tt.want)
			}
			if got2 != tt.want2 {
				t.Errorf("concatenateNumbers() = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func Test_digit(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "case 1",
			x:    1,
			want: 1,
		},
		{
			name: "case 2",
			x:    11,
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := digit(tt.x)
			if got == tt.want {
				t.Errorf("digit() = %v, want %v", got, tt.want)
			}
		})
	}
}
