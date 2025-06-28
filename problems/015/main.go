package main

import "fmt"

func main() {
	solve()
}

func solve() {
	mapping := make((map[int8][]int))
	for i := int8(0); i <= 20; i++ {
		for j := 0; j <= 20; j++ {
			if i == 0 || j == 0 {
				mapping[i] = append(mapping[i], 1)
			} else {
				mapping[i] = append(mapping[i], mapping[i][j-1]+mapping[i-1][j])
			}
		}
	}
	fmt.Println(mapping[20][20])
}
