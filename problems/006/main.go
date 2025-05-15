package main

import "fmt"

func main() {
	const TARGET = 100
	func1(TARGET)
}

func func1(target int) {
	x := 0
	y := 0

	for i := 1; i < target+1; i++ {
		x = i*i + x
		y = y + i
	}
	fmt.Println(y*y - x)
}
