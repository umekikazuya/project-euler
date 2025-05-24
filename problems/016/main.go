package main

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

func main() {
	const POWER_NUMBER = 1000

	var x = big.NewInt(2)
	var result = big.NewInt(1)
	var loop int16 = 1
	for loop <= POWER_NUMBER {
		result.Mul(result, x)
		loop++
	}

	strings := strings.Split(result.String(), "")

	var sum = 0
	for _, o := range strings {
		num, _ := strconv.Atoi(o)
		sum = sum + num
	}

	fmt.Println(sum)
}
