package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	solve()
}

// 1/3, 2/6, 3/9
// 2/3, 4/6, 6/9
func solve() {
	numbers := make(map[int]int)
	for i := 2; i <= 9; i++ {
		for j := 1; j < i; j++ {
			a, b := j, i
			for b < 100 {
				// fmt.Println(a, "/", b)
				if condition(a, b) {
					numbers[a] = b
				}
				a += j
				b += i
			}
		}
	}
	x, y := 1, 1
	for key, value := range numbers {
		x *= key
		y *= value
	}
	divide(x, y)
}

// 13/39
func condition(a int, b int) bool {
	stringA, stringB := strconv.Itoa(a), strconv.Itoa(b)
	if len(stringA) != 2 || len(stringB) != 2 {
		return false
	}
	if stringA[len(stringA)-1] == '0' || stringB[len(stringB)-1] == '0' {
		return false
	}
	if !containSameNum(stringA, stringB) {
		return false
	}
	rawX, rawY := func1(stringA, stringB)
	x, _ := strconv.Atoi(rawX)
	y, _ := strconv.Atoi(rawY)
	return float64(y)/float64(x) == float64(b)/float64(a)
}

func containSameNum(a string, b string) bool {
	arrA := strings.Split(a, "")
	arrB := strings.Split(b, "")
	same := 0
	for _, num := range arrA {
		if slices.Contains(arrB, num) {
			same++
		}
	}
	return same == 1
}

func func1(a string, b string) (string, string) {
	arrA := strings.Split(a, "")
	copyArrA := arrA
	arrB := strings.Split(b, "")
	for i, num := range arrA {
		if slices.Contains(arrB, num) {
			arrA = append(arrA[:i], arrA[i+1:]...)
		}
	}
	for i, num := range arrB {
		if slices.Contains(copyArrA, num) {
			arrB = append(arrB[:i], arrB[i+1:]...)
		}
	}
	return strings.Join(arrA, ""), strings.Join(arrB, "")
}

func divide(x int, y int) {
	a := 2
	for a <= x {
		if x%a == 0 && y%a == 0 {
			x /= a
			y /= a
		} else {
			a++
		}
	}
	fmt.Println(x, y)
}
