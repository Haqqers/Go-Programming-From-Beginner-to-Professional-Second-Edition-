package main

import "fmt"

func main() {
	var a int8 = 125  // max int8 value is 128, beyond is wraparound
	var b uint8 = 253 // max uint8 value is 256, beyond is wraparound
	for i := 0; i < 5; i++ {
		a++
		b++
		fmt.Println(i, ")", "int8", a, "uint8", b)
	}
}
