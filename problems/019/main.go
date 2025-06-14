package main

import "fmt"

// 20 世紀 (1901 年 1 月 1 日から 2000 年 12 月 31 日) に、月の 1 日が日曜日だった日は何日ありましたか?
func main() {
	// それぞれの月の日数を格納
	var days = (map[int]int{
		1:  31,
		2:  28,
		3:  31,
		4:  30,
		5:  31,
		6:  30,
		7:  31,
		8:  31,
		9:  30,
		10: 31,
		11: 30,
		12: 31,
	})

	// 1901/01/01は月曜日スタート、曜日を変数で管理
	current := 2
	// 日曜日だった数
	count := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			lastDay := days[month]
			if month == 2 && isLeapYear(year) {
				lastDay++
			}
			current = getNextStartDay(current, lastDay)
			current++
			if current%7 == 0 {
				count++
			}
		}
	}
	if current%7 == 0 {
		count--
	}
	fmt.Println(count)
}

// 閏年
func isLeapYear(year int) bool {
	if year%400 == 0 {
		return false
	}
	return year%4 == 0
}

// 　来月のスタートが何曜日かを返す関数
func getNextStartDay(current int, lastDay int) int {
	return (((lastDay - 1) + current) % 7)
}
