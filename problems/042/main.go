package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func main() {
	words, err := getWords("./0042_words.txt")
	if err != nil {
		fmt.Errorf("エラー: %s", err)
	}
	solve(words)
}

func solve(words []string) {
	// 三角数の上限値を算出
	maxWordLn := int16(0)
	for _, word := range words {
		if maxWordLn < int16(len(word)) {
			maxWordLn = int16(len(word))
		}
	}
	// 三角数生成
	triNums := getTriNumbers(maxWordLn * int16('Z'-'A'+1))

	count := 0
	for _, word := range words {
		// 単語のスコアを算出し、和が三角数であれば変数countをインクリメント
		if _, ok := triNums[sumWord(word)]; ok {
			count++
		}
	}
	fmt.Println(count)
}

// 上限値まで三角数生成しスライスに格納
func getTriNumbers(limit int16) map[int16]struct{} {
	nums := make(map[int16]struct{})
	nums[1] = struct{}{}
	x, y := int16(2), int16(0)
	for y < limit {
		y = (x*x + x) / int16(2)
		nums[y] = struct{}{}
		x++
	}
	return nums
}

// 受け取った文字列を数字に変換し和を算出
func sumWord(s string) int16 {
	sum := int16(0)
	for i := 0; i < len(s); i++ {
		sum += int16(s[i] - 'A' + 1)
	}
	return sum
}

// ファイルからWordsスライスを取得
func getWords(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("ファイル読み込み失敗")
	}
	defer file.Close()

	words := []string{}
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), ",") {
			words = append(words, strings.Trim(word, `"`))
		}
	}
	return words, nil
}
