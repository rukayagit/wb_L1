package main

import (
	"fmt"
	"math/big"
)

func main() {
	// инициализация переменных a и b со  значениями больше 2^20
	a := new(big.Int)
	b := new(big.Int)
	a.SetString("8934588", 10)
	b.SetString("5368274", 10)

	// операции с целыми числами
	sum := new(big.Int).Add(a, b)
	difference := new(big.Int).Sub(a, b)
	product := new(big.Int).Mul(a, b)
	quotient := new(big.Int).Quo(a, b)

	fmt.Printf("a = %s\n", a.String())
	fmt.Printf("b = %s\n", b.String())
	fmt.Printf("a + b = %s\n", sum.String())
	fmt.Printf("a - b = %s\n", difference.String())
	fmt.Printf("a * b  = %s\n", product.String())
	fmt.Printf("a / b = %s\n", quotient.String())

	// операция деления с плавающей запятой
	af := new(big.Float).SetInt(a)
	bf := new(big.Float).SetInt(b)
	quotientFloat := new(big.Float).Quo(af, bf)

	fmt.Printf("a / b (float) = %s\n", quotientFloat.Text('f', 10))
}
