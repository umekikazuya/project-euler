package main

import (
	"fmt"
	"log"
	"strconv"
)

// 問題
// <p>Some positive integers $n$ have the property that the sum $[n + \operatorname{reverse}(n)]$ consists entirely of odd (decimal) digits. For instance, $36 + 63 = 99$ and $409 + 904 = 1313$. We will call such numbers <dfn>reversible</dfn>; so $36$, $63$, $409$, and $904$ are reversible. Leading zeroes are not allowed in either $n$ or $\operatorname{reverse}(n)$.</p>
// <p>There are $120$ reversible numbers below one-thousand.</p>
// <p>How many reversible numbers are there below one-billion ($10^9$)?</p>

var (
	// LIMIT 探索対象は10億
	LIMIT = (1000000000)
	ONE   = (1)
	ZERO  = (0)
)

func main() {
	fmt.Printf("答えは...。'%d'です。", solve())
}

func solve() int {
	i := (11)
	result := 0
	for {
		copyIndex := i
		if isMutableNumber(copyIndex) {
			result++
		}
		if i >= LIMIT {
			log.Println(i)
			break
		}
		i++
	}
	return result
}

// isMutableNumber は可逆数かどうかを判定
//
// * ベースナンバーが10で割り切れないこと
// * 和が正の整数
// * 和が奇数であること
func isMutableNumber(n int) bool {
	if n%10 == 0 {
		return false
	}
	sum := sumWithReverseNumber(n)
	if sum <= 0 {
		return false
	}
	for _, s := range strconv.Itoa(sum) {
		if s == '0' {
			return false
		}
		if int(s-'0')%2 == 0 {
			return false
		}
	}
	return true
}

func sumWithReverseNumber(n int) int {
	r, err := reverseInt(n)
	if err != nil {
		return (0)
	}
	sum := n + r
	return sum
}

// reverseInt は数字を逆にする関数
func reverseInt(n int) (int, error) {
	r := reverse(strconv.Itoa(n))
	reverseNumber, err := strconv.Atoi(r)
	if err != nil {
		return 0, nil
	}
	return reverseNumber, nil
}

// reverse は文字を逆にする関数
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
