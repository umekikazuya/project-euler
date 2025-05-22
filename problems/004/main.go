package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 左辺
	var x int = 999
	arr := []int{}
	for 99 < x {
		// 右辺
		var y int = x - 1
		for 99 < y {
			// 回分数チェック
			if isfrequency(x * y) {
				arr = append(arr, (x * y))
			}
			y--
		}
		x--
	}
	sort.Sort(sort.IntSlice(arr))
	fmt.Println(arr[len(arr)-1])
}

func isfrequency(num int) bool {
	var x = strings.Split(strconv.Itoa(int(num)), "")
	var y = reverseArray(strings.Split(strconv.Itoa(int(num)), ""))
	for i, o := range x {
		if o != y[i] {
			return false
		}
	}
	return true
}

func reverseArray(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
