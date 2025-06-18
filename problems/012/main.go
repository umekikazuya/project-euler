package main

import "fmt"

func main() {
	solve()
}

func solve() {
	x, y := 1, 0
	for {
		if 500 <= count(y) {
			break
		}
		y += x
		x++
	}
	fmt.Println(y)
}

// 約数をカウントする関数
func count(num int) int {
	if num == 1 {
		return 1
	}
	x := 2
	count := 1
	for x*x <= num {
		if num%x == 0 {
			if x*x == num {
				count++
			} else {
				count = count + 2
			}
		}
		x++
	}
	return count
}
