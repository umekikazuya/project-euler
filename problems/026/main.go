package main

import (
	"fmt"
)

func main() {
	fmt.Println(solve(1000))
}

func solve(n int) int {
	c := 0
	for i := 2; i < n; i++ {
		if c <= getReciprocalCycles(i) {
			c = i
		}
	}
	return c
}

// 1. 任意の数を10倍してnで割った余りを返す関数(小数点何桁なのかの場所も記憶させる)
// 2. 1の関数を実行しながら同じ余りがでてきたかを検証する関数
// 3. ループ(d-1)を回しながら、1を実行する、2が出てきたら処理終了

// getDiv は任意の数を10倍してnで割った余りを返す関数
func getDiv(n int, d int) int {
	return (10 * n) % d
}

func getReciprocalCycles(n int) int {
	d := 1
	// "同じ余りが出てきた場合、循環節を返す"ため、インデックスの管理を行うメモリ上の配列
	cache := make(map[int]int, n-1)
	for i := 1; i <= n; i++ {
		d = getDiv(d, n)
		if d == 0 {
			return 0 // 割り切れる場合はとりあえずリターンする、ここも循環節を計算する必要があるか後で確認
		}
		if v, ok := cache[d]; ok {
			return i - v
		} else {
			cache[d] = i
		}
	}
	return 0
}
