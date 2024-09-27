package main

import (
	"fmt"
	"math"
	"math/big"
)

func main() {
	intA := math.MaxInt64
	intA = intA + 1
	bigA := big.NewInt(math.MaxInt64)
	bigA.Add(bigA, big.NewInt(1))

	bigB := big.NewFloat(math.MaxFloat64)
	bigB.Add(bigB, big.NewFloat(1.01))

	fmt.Println("MaxInt64: ", math.MaxInt64)
	fmt.Println("Int     :", intA)
	fmt.Println("Big Int : ", bigA.String())

	fmt.Println("Big Float :", bigB)
}
