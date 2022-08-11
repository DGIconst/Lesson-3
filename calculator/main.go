package main

import (
	"fmt"
	"math"
	"math/big"
	"os"
)

func faktorial(n int64) int64 {
	if n == 0 {
		return 1
	}
	return n * faktorial(n-1)
}

func main() {
	var a, b, res float32
	var op string
	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)
	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
	fmt.Print("Введите арифметическую операцию (+, -, *, /, ^, fakt): ")
	fmt.Scanln(&op)
	switch op {
	case "+":
		res = a + b
	case "-":
		res = a - b
	case "*":
		res = a * b
	case "/":
		if b == 0 {
			fmt.Print("деление на 0 не допустимо")
			return
		}
		res = a / b
	case "^":
		a := float64(a)
		b := float64(b)
		res := math.Pow(a, b)
		fmt.Printf("Результат выполнения операции: %.2f\n", res)
		return
	case "fakt":
		a := int64(a)
		b := int64(b)
		if a < 0 || b < 0 {
			fmt.Println("отрицательное число факториала не допустимо")
			os.Exit(2)
		}
		bigFakt1 := big.NewInt(faktorial(a))
		bigFakt2 := big.NewInt(faktorial(b))
		fmt.Println("Результат выполнения операции:", bigFakt1)
		fmt.Println("Результат выполнения операции:", bigFakt2)
		return
		// всю голову сломал так и не понял как вывести факториал числа больше 20 ((
	default:
		fmt.Println("Операция выбрана неверно")
		os.Exit(1)
	}
	fmt.Printf("Результат выполнения операции: %.2f\n", res)
}
