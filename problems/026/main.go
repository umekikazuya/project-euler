package main

import (
	"fmt"
)

func main() {
	fmt.Println(solve(1000))
}

// solve は問題の解答を求める関数
func solve(n int16) int16 {
	c := int16(0)
	for i := int16(1); i < n; i++ {
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
func getDiv(n int16, d int16) int16 {
	return (10 * n) % d
}

// getReciprocalCycles は任意の数字を受け取って、その数の循環節の数を返却する関数
//
// 割り切れる場合は0を返却
// 循環節がない場合は0を返却
func getReciprocalCycles(n int16) int16 {
	d := int16(1)
	// "同じ余りが出てきた場合循環節数をを返却したい"ため、インデックス数と剰余算の結果の管理を行う配列
	cache := make(map[int16]int16, n-1)
	for i := int16(1); i <= n; i++ {
		d = getDiv(d, n)
		if d == 0 {
			return 0 // 割り切れる場合は0を返却する
		}
		if v, ok := cache[d]; ok {
			return i - v
		} else {
			cache[d] = i
		}
	}
	return 0 // 循環節がない場合は0を返却
}
