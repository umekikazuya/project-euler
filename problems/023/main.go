package main

import (
	"fmt"
	"slices"
)

// <p>完全数とは、その約数（自分自身を除く）の和がその数自身と正確に等しい数のことです。例えば、28 の約数の和は 1 + 2 + 4 + 7 + 14 = 28 であるため、28 は完全数とされます。</p>
// <p>ある数 n が、約数の和が n より小さい場合は「不足数」と呼ばれ、逆に約数の和が n を超える場合は「過剰数」と呼ばれます。</p>
// <p>12 は最も小さい過剰数であり、1 + 2 + 3 + 4 + 6 = 16 となります。そして、2つの過剰数の和として表すことができる最小の数は 24 です。数学的解析により、28123 を超えるすべての整数は2つの過剰数の和として表すことができることが示されています。しかし、2つの過剰数の和として表せない最大の数がこの上限より小さいことは知られているものの、解析によってこの上限をさらに小さくすることはできません。</p>
// <p>2つの過剰数の和として表すことができないすべての正の整数の総和を求めなさい。</p>

// 過剰数のリストを作成
// 全ての和のパターンを算出
// それ以外の数字の和を出す
func main() {
	// 過剰数のリストを作成
	abundantNumbers := []int{}
	for i := 1; i < 28123; i++ {
		if isAbundantNumber(i) {
			abundantNumbers = append(abundantNumbers, i)
		}
	}

	// 全部の和をスライスに格納
	allSum := []int{}
	for _, abundantNumberA := range abundantNumbers {
		a := abundantNumberA + abundantNumberA
		if a <= 28123 {
			allSum = append(allSum, a)
		}
		for _, abundantNumberB := range abundantNumbers {
			b := abundantNumberA + abundantNumberB
			if b <= 28123 {
				allSum = append(allSum, b)
			}
		}
	}
	uniqArr := Uniq(allSum)
	sum := 0
	for i := 1; i < 28123; i++ {
		if !slices.Contains(uniqArr, i) {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}

// 過剰数であればTRUEを返却
func isAbundantNumber(num int) bool {
	divisor := 2
	factor := []int{1}
	for divisor*divisor <= num {
		if num%divisor == 0 {
			factor = append(factor, divisor)
			if divisor*divisor != num {
				factor = append(factor, num/divisor)
			}
		}
		divisor++
	}
	sum := 0
	for _, item := range factor {
		sum = sum + item
	}
	return num < sum
}

// スライス内の重複を削除
func Uniq(s []int) []int {
	slices.Sort(s)
	return slices.Clip(slices.Compact(s))
}
