package main

import (
	"fmt"
	"testing"
)

func BenchmarkSolve(b *testing.B) {
	words, err := getWords("./0042_words.txt")
	if err != nil {
		fmt.Errorf("エラー: %s", err)
	}
	for i := 0; i < b.N; i++ {
		solve(words)
	}
}
