package main

import "fmt"

func main() {
	var x int64 = 1
	for {
		if isDivisible(x) {
			// 割り切れたらループを終了
			break
		}
		x = x + 1
	}
	fmt.Println(x)
}

// 1~10で割り切れるかチェック
func isDivisible(num int64) bool {
	for i := 1; i <= 20; i++ {
		if int(num)%i != 0 {
			return false
		}
	}
	return true
}
