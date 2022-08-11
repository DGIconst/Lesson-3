package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a int
	fmt.Print("Введите число: ")
	fmt.Scanln(&a)
	for i := 1; i <= a; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			fmt.Println(i, "- простое число")
		} else {
			continue
		}
	}
}
