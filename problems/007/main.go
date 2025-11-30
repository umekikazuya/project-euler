package main

func main() {
	solve()
}

func solve() int32 {
	var max int32 = 10001
	var x, y int32 = 2, 1
	for {
		if isPrimeNumber(x) {
			y++
		}
		if max < y {
			break
		}
		x++
	}
	return x
}

// isPrimeNumber は素数チェックを実施
func isPrimeNumber(x int32) bool {
	if x < 2 {
		return false
	}
	var a int32 = 2
	for a*a <= x {
		if x%a == 0 {
			return false
		}
		a++
	}
	return true
}
