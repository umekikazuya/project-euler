package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	solve()
}

func solve() {
	x := 100
	count := 0
	for {
		if isBouncyNumber(x) {
			count++
		}
		if isNinetyNinePer(count, x) {
			break
		}
		x++
	}
	fmt.Println(x)
}

// 99%ぴったりであるかの判定
func isNinetyNinePer(target int, all int) bool {
	return target*100/all == 99 && (target*100%all == 0)
}

func isBouncyNumber(n int) bool {
	s := strings.Split(strconv.Itoa(n), "")
	newS := []string{}
	for i, v := range s {
		if i == 0 {
			newS = append(newS, v)
			continue
		}
		if v == s[i-1] {
			continue
		}
		newS = append(newS, v)
	}
	// 1桁,2桁の場合はBoucyNumberとみなさない
	if len(newS) <= 2 {
		return false
	}
	isIncreasing := newS[0] < newS[1]
	k := 1
	if newS[0] == newS[1] {
		isIncreasing = (newS[1] == newS[2])
		k = 2
	}
	for i := k; i < len(newS); i++ {
		if isIncreasing != (newS[i-1] < newS[i]) {
			return true
		}
	}
	return false
}
