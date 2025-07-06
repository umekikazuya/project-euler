package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	solve()
}

// 2: 4*4, 5*5, 6*6, 7*7, 8*8, 9*9
// 3: 5*5*5, 6*6*6, 7*7*7, 8*8*8, 9*9*9
// 4: 6*6*6*6, 7*7*7*7, 8*8*8*8, 9*9*9*9
// 5:7*7*7*7*7,8*8*8*8*8
func solve() {

	x := 2
	var count int
	for x < 20 {
		for i := 2; i < 10; i++ {
			if len(strconv.FormatFloat(math.Pow(float64(i), float64(x)), 'f', -1, 64)) == x {
				fmt.Println(i, x)
				count++
			}
		}
		x++
	}

}
