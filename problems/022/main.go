package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	// スコアマップの生成
	scoreMap := generateScoreMap()
	// ファイル内のテキスト情報を文字列スライスに変換する
	data, _ := loadFileData()

	result := 0
	for index, name := range data {
		x := (index + 1) * calculateScoreFromName(name, scoreMap)
		result = result + x
	}
	fmt.Println(result)
}

// ファイルの読み込み、生データをスライスに格納、アルファベット順にソート。
func loadFileData() (data []string, err error) {
	file, err := os.Open("./0022_names.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data = []string{}
	for scanner.Scan() {
		data = strings.Split(scanner.Text(), ",")
	}
	sort.Strings(data)
	return data, nil
}

// 文字列を受け取ってスコアを計算する
func calculateScoreFromName(name string, scoreMap map[string]int) int {
	result := 0
	for _, v := range strings.Split(name, "") {
		if score, ok := scoreMap[v]; ok {
			result += score
		}
	}
	return result
}

func generateScoreMap() map[string]int {
	return map[string]int{
		"A": 1,
		"B": 2,
		"C": 3,
		"D": 4,
		"E": 5,
		"F": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"J": 10,
		"K": 11,
		"L": 12,
		"M": 13,
		"N": 14,
		"O": 15,
		"P": 16,
		"Q": 17,
		"R": 18,
		"S": 19,
		"T": 20,
		"U": 21,
		"V": 22,
		"W": 23,
		"X": 24,
		"Y": 25,
		"Z": 26,
	}
}
