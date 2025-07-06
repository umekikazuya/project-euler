package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 1801, 662664 ns/op
func main() {
	solve()
}

func solve() {
	wordCount := 0
	for i := 1; i <= 1000; i++ {
		wordCount += getWordCount(i)
	}
	fmt.Println(wordCount)
}

func getWordCount(n int) int {
	mapping := map[int]string{
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}
	if v, ok := mapping[n]; ok {
		return len(v)
	}
	switch getDiget(n) {
	case 2:
		return len(getWordTwoDidget(n, mapping))
	case 3:
		return len(getWordThreeDidget(n, mapping))

	default:
		return len("onethousand")

	}
}

func getWordTwoDidget(n int, mapping map[int]string) string {
	mappingForTwoDigeit := map[int]string{
		2:  "twenty",
		3:  "thirty",
		4:  "forty",
		5:  "fifty",
		6:  "sixty",
		7:  "seventy",
		8:  "eighty",
		9:  "ninety",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}
	if v, ok := mapping[n]; ok {
		return v
	}
	x := strings.Split(strconv.Itoa(n), "")
	return mappingForTwoDigeit[int(x[0][0]-'0')] + mapping[int(x[1][0]-'0')]
}

func getWordThreeDidget(n int, mapping map[int]string) string {
	raw := strings.Split(strconv.Itoa(n), "")
	x := raw[0][0]
	y := strings.Join(raw[1:], "")

	word := ""

	// 3桁目を求める
	word += mapping[int(x-'0')]
	word += "hundred"

	if y == "00" {
		return word
	}

	word += ("and")
	numberY, _ := strconv.Atoi(y)
	word += getWordTwoDidget(numberY, mapping)
	return word
}

func getDiget(n int) int {
	return len(strconv.Itoa(n))
}
