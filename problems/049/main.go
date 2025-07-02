package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	solve()
}

func solve() {
	primeNumbers := make(map[string][]int16, 0)
	for i := int16(1000); i < 10000; i++ {
		if isPrimeNumber(i) {
			s := strings.Split(strconv.Itoa(int(i)), "")
			sort.Strings(s)

			primeNumbers[strings.Join(s, "")] = append(primeNumbers[strings.Join(s, "")], i)
		}
	}
	for i, v := range primeNumbers {
		if len(primeNumbers[i]) <= 2 {
			continue
		}
		for _, j := range v {
			for _, k := range v {
				if k == j {
					continue
				}
				if slices.Contains(v, j-k+j) && j != 4817 {
					fmt.Printf("答えは%d%d%dです。", int(k), int(j), int(j-k+j))
					return
				}
			}
		}
	}
}

// 素数判定
func isPrimeNumber(num int16) bool {
	x := int16(2)
	for x*x <= num {
		if num%x == 0 {
			return false
		}
		x++
	}
	return true
}
