package main

import (
	"errors"
	"fmt"
	"log"
	"math/big"
)

// 問題
// <p>Some positive integers $n$ have the property that the sum $[n + \operatorname{reverse}(n)]$ consists entirely of odd (decimal) digits. For instance, $36 + 63 = 99$ and $409 + 904 = 1313$. We will call such numbers <dfn>reversible</dfn>; so $36$, $63$, $409$, and $904$ are reversible. Leading zeroes are not allowed in either $n$ or $\operatorname{reverse}(n)$.</p>
// <p>There are $120$ reversible numbers below one-thousand.</p>
// <p>How many reversible numbers are there below one-billion ($10^9$)?</p>

var (
	// LIMIT 探索対象は10億
	LIMIT = big.NewInt(1000000000)
	ONE   = big.NewInt(1)
	ZERO  = big.NewInt(0)
)

func main() {
	fmt.Printf("答えは...。'%d'です。", solve())
}

func solve() int {
	i := big.NewInt(11)
	result := 0
	for {
		copyIndex := new(big.Int).Set(i)
		if isMutableNumber(*copyIndex) {
			result++
		}
		if i.Cmp(LIMIT) >= 0 {
			log.Println(i)
			break
		}
		i.Add(i, ONE)
	}
	return result
}

// isMutableNumber は可逆数かどうかを判定
//
// * ベースナンバーが10で割り切れないこと
// * 和が正の整数
// * 和が奇数であること
func isMutableNumber(n big.Int) bool {
	mod := new(big.Int).Mod(&n, big.NewInt(10))
	if mod.Cmp(ZERO) == 0 {
		return false
	}
	sum := sumWithReverseNumber(n)
	if sum.Cmp(ZERO) < 0 {
		return false
	}
	for _, s := range sum.String() {
		if s == '0' {
			return false
		}
		if int(s-'0')%2 == 0 {
			return false
		}
	}
	return true
}

func sumWithReverseNumber(n big.Int) big.Int {
	r, err := reverseBigInt(n)
	if err != nil {
		return *big.NewInt(0)
	}
	sum := new(big.Int).Add(&n, &r)
	return *sum
}

// reverseBigInt は数字を逆にする関数
func reverseBigInt(n big.Int) (big.Int, error) {
	r := reverse(n.Text(10))
	i := new(big.Int)
	_, success := i.SetString(r, 10)
	if !success {
		return *big.NewInt(0), errors.New("変換ミス")
	}
	return *i, nil
}

// reverse は文字を逆にする関数
func reverse(s string) string {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return string(rs)
}
