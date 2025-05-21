package main

import "fmt"

// a+b+c=1000
// a*a + b*b = c*c
// a < b < c
func main() {
	const MAX int32 = 1000

	var a int32 = 1
	var b int32 = 1
	for a+b < 1000 {
		var b int32 = 1
		for a+b < 1000 {
			var sum int32 = (a*a + b*b)
			c := func1(sum)
			if a+b+c == MAX {
				fmt.Println(a * b * c)
				return
			}
			b = b + 1
		}
		a = a + 1
	}
}

func func1(num int32) int32 {
	var isDivisible int32 = 1
	for isDivisible*isDivisible <= num {
		if (isDivisible * isDivisible) == num {
			return isDivisible
		}
		isDivisible = isDivisible + 1
	}
	return 0
}
